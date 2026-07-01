package main

import (
	"encoding/json"
	"errors"
	"io"
	"net"
	"strings"

	"github.com/router-for-me/CLIProxyAPI/v7/sdk/pluginabi"
	"github.com/router-for-me/CLIProxyAPI/v7/sdk/pluginapi"
)

func valueOr(value, fallback string) string {
	if strings.TrimSpace(value) != "" {
		return value
	}
	if strings.TrimSpace(fallback) != "" {
		return fallback
	}
	return "未知"
}

func textResponse(status int, text string) pluginapi.ManagementResponse {
	return pluginapi.ManagementResponse{
		StatusCode: status,
		Headers:    map[string][]string{"content-type": {"text/plain; charset=utf-8"}},
		Body:       []byte(text),
	}
}

func okEnvelope(result any) ([]byte, error) {
	raw, errMarshal := json.Marshal(result)
	if errMarshal != nil {
		return nil, errMarshal
	}
	return json.Marshal(pluginabi.Envelope{OK: true, Result: json.RawMessage(raw)})
}

func errorEnvelope(code, message string) []byte {
	raw, _ := json.Marshal(pluginabi.Envelope{OK: false, Error: &pluginabi.Error{Code: code, Message: message}})
	return raw
}

func closeBody(body io.Closer) {
	if body == nil {
		return
	}
	_ = body.Close()
}

func compactError(err error) string {
	if err == nil {
		return ""
	}
	text := err.Error()
	var dnsErr *net.DNSError
	if errors.As(err, &dnsErr) && dnsErr.Name != "" {
		text = dnsErr.Err + ": " + dnsErr.Name
	}
	return truncateRunes(text, 500)
}

func truncateRunes(text string, max int) string {
	if max <= 0 {
		return ""
	}
	runes := []rune(text)
	if len(runes) <= max {
		return text
	}
	return string(runes[:max]) + "..."
}

func retryablePublicIPError(err error) bool {
	if err == nil {
		return false
	}
	text := strings.ToLower(err.Error())
	return strings.Contains(text, "forcibly closed") ||
		strings.Contains(text, "connection reset") ||
		strings.Contains(text, "wsarecv") ||
		strings.Contains(text, "unexpected eof") ||
		strings.Contains(text, "timeout")
}

func publicIPErrorMessage(err error) string {
	if err == nil {
		return ""
	}
	text := strings.ToLower(err.Error())
	if strings.Contains(text, "forcibly closed") || strings.Contains(text, "connection reset") || strings.Contains(text, "wsarecv") {
		return "连接被远程服务关闭，已跳过该查询源"
	}
	if strings.Contains(text, "timeout") {
		return "查询超时，已跳过该查询源"
	}
	if strings.Contains(text, "no such host") {
		return "DNS 解析失败，已跳过该查询源"
	}
	return compactError(err)
}
