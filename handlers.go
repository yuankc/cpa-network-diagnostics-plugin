package main

import (
	"encoding/json"
	"net/http"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/router-for-me/CLIProxyAPI/v7/sdk/pluginabi"
	"github.com/router-for-me/CLIProxyAPI/v7/sdk/pluginapi"
)

var (
	diagnosticsCacheMu      sync.Mutex
	diagnosticsCacheData    diagnostics
	diagnosticsCacheExpires time.Time
	diagnosticsCacheLoading bool
	diagnosticsCacheCond    = sync.NewCond(&diagnosticsCacheMu)
)

func handleMethod(method string, payload []byte) ([]byte, error) {
	switch method {
	case pluginabi.MethodPluginRegister, pluginabi.MethodPluginReconfigure:
		return okEnvelope(registrationPayload())
	case pluginabi.MethodManagementRegister:
		return okEnvelope(pluginapi.ManagementRegistrationResponse{
			Routes: []pluginapi.ManagementRoute{
				{Method: http.MethodGet, Path: "/diagnostics/status", Description: "Returns CPA process network diagnostics as JSON."},
			},
			Resources: []pluginapi.ResourceRoute{
				{Path: "/dashboard", Menu: "网络诊断", Description: "显示公网 IP、本地 IP、DNS 和 OpenAI 连接情况。"},
				{Path: "/status", Description: "Returns CPA process network diagnostics as JSON for the diagnostics dashboard."},
				{Path: "/status/runtime", Description: "Returns runtime, local IP, and proxy diagnostics."},
				{Path: "/status/public-ip", Description: "Returns public IP diagnostics."},
				{Path: "/status/ip-risk", Description: "Returns IP reputation diagnostics."},
				{Path: "/status/openai", Description: "Returns OpenAI availability diagnostics."},
				{Path: "/status/geo", Description: "Returns geographic consistency diagnostics."},
				{Path: "/status/dns", Description: "Returns DNS diagnostics."},
				{Path: "/status/connectivity", Description: "Returns HTTP connectivity diagnostics."},
				{Path: "/status/outbound", Description: "Returns outbound source address diagnostics."},
			},
		})
	case pluginabi.MethodManagementHandle:
		var req pluginapi.ManagementRequest
		if len(payload) > 0 {
			if errDecode := json.Unmarshal(payload, &req); errDecode != nil {
				return nil, errDecode
			}
		}
		return okEnvelope(handleManagement(req))
	default:
		return errorEnvelope("unknown_method", "unknown method: "+method), nil
	}
}

func registrationPayload() registration {
	return registration{
		SchemaVersion: pluginabi.SchemaVersion,
		Metadata: pluginapi.Metadata{
			Name:             pluginName,
			Version:          pluginVersion,
			Author:           pluginAuthor,
			GitHubRepository: pluginRepo,
			Logo:             "",
		},
		Capabilities: capabilities{ManagementAPI: true},
	}
}

func handleManagement(req pluginapi.ManagementRequest) pluginapi.ManagementResponse {
	if kind := statusPathKind(req.Path); kind != "" {
		return diagnosticsJSONResponse(kind)
	}
	return pluginapi.ManagementResponse{
		StatusCode: http.StatusOK,
		Headers: map[string][]string{
			"content-type":  {"text/html; charset=utf-8"},
			"cache-control": {"no-store"},
		},
		Body: []byte(renderDashboardHTML()),
	}
}

func diagnosticsJSONResponse(kind string) pluginapi.ManagementResponse {
	body, errMarshal := json.MarshalIndent(diagnosticsPayload(kind), "", "  ")
	if errMarshal != nil {
		return textResponse(http.StatusInternalServerError, errMarshal.Error())
	}
	return pluginapi.ManagementResponse{
		StatusCode: http.StatusOK,
		Headers: map[string][]string{
			"content-type":  {"application/json; charset=utf-8"},
			"cache-control": {"no-store"},
		},
		Body: body,
	}
}

func isStatusPath(path string) bool {
	return statusPathKind(path) != ""
}

func statusPathKind(path string) string {
	cleaned := "/" + strings.Trim(strings.TrimSuffix(path, "/"), "/")
	bases := []string{"/status", "/diagnostics/status", "/v0/management/diagnostics/status", "/v0/resource/plugins/diagnostics/status"}
	for _, base := range bases {
		if cleaned == base {
			return "full"
		}
		if strings.HasPrefix(cleaned, base+"/") {
			kind := strings.TrimPrefix(cleaned, base+"/")
			switch kind {
			case "runtime", "proxy", "public-ip", "ip-risk", "openai", "geo", "dns", "connectivity", "outbound":
				return kind
			}
		}
	}
	return ""
}

func diagnosticsPayload(kind string) any {
	switch kind {
	case "runtime":
		hostname, _ := os.Hostname()
		tzName, tzUTC := localTimezone()
		proxy := collectProxyInfo()
		return map[string]any{
			"runtime":   runtimeInfo{Hostname: hostname, GOOS: runtime.GOOS, GOARCH: runtime.GOARCH, PID: os.Getpid(), TimezoneName: tzName, TimezoneUTC: tzUTC},
			"local_ips": collectLocalIPs(),
			"proxy":     proxy,
		}
	case "proxy":
		return collectProxyInfo()
	case "public-ip":
		return detectPublicIP()
	case "ip-risk":
		publicIP := detectPublicIP()
		payload := map[string]any{"public_ip": publicIP, "ip_risk": ipRiskProfile{}}
		if publicIP.IP != "" {
			payload["ip_risk"] = detectIPRisk(publicIP.IP)
		}
		return payload
	case "openai":
		return detectOpenAIAvailability()
	case "geo":
		tzName, tzUTC := localTimezone()
		publicIP := detectPublicIP()
		openAI := detectOpenAIAvailability()
		return evaluateGeoConsistency(publicIP, openAI, tzName, tzUTC)
	case "dns":
		return checkDNS([]string{"chatgpt.com", "api.openai.com", "auth.openai.com", "cdn.openai.com"})
	case "connectivity":
		return checkConnectivity()
	case "outbound":
		return detectOutboundSources([]string{"api.openai.com:443", "chatgpt.com:443", "1.1.1.1:443"})
	default:
		return cachedDiagnostics()
	}
}

func cachedDiagnostics() diagnostics {
	now := time.Now()
	diagnosticsCacheMu.Lock()
	for {
		if !diagnosticsCacheExpires.IsZero() && now.Before(diagnosticsCacheExpires) {
			cached := diagnosticsCacheData
			diagnosticsCacheMu.Unlock()
			return cached
		}
		if !diagnosticsCacheLoading {
			diagnosticsCacheLoading = true
			diagnosticsCacheMu.Unlock()

			data := collectDiagnostics()

			diagnosticsCacheMu.Lock()
			diagnosticsCacheData = data
			diagnosticsCacheExpires = time.Now().Add(30 * time.Second)
			diagnosticsCacheLoading = false
			diagnosticsCacheCond.Broadcast()
			diagnosticsCacheMu.Unlock()
			return data
		}
		diagnosticsCacheCond.Wait()
		now = time.Now()
	}
}
