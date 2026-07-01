package main

import (
	"net"
	"strings"
)

func summarizeRisk(data diagnostics) riskSummary {
	signals := make([]string, 0)
	level := "low"

	if data.PublicIP.IP == "" {
		level = maxRisk(level, "unknown")
		signals = append(signals, "所有公共 IP 查询接口均失败，无法确认出口 IP")
	}

	// Compare egress IPs by address family to avoid dual-stack false positives.
	v4, v6 := distinctIPsByFamily(data.PublicIP.Checks)
	if len(v4) > 1 {
		level = maxRisk(level, "warning")
		signals = append(signals, "多个服务返回了不同的 IPv4 出口地址（"+strings.Join(v4, ", ")+"），可能存在代理链、NAT 或接口异常")
	}
	if len(v6) > 1 {
		level = maxRisk(level, "warning")
		signals = append(signals, "多个服务返回了不同的 IPv6 出口地址，可能存在多出口或代理异常")
	}

	// IP reputation signals are high-value indicators for AI service blocks.
	if data.IPRisk.Determined {
		if data.IPRisk.IsTor {
			level = maxRisk(level, "high")
			signals = append(signals, "出口 IP 被识别为 Tor 出口节点，几乎必然被 AI 服务拦截")
		}
		if data.IPRisk.IsProxy {
			level = maxRisk(level, "high")
			signals = append(signals, "出口 IP 被标记为公共代理，属于高风险类型")
		}
		if data.IPRisk.IsAbuser {
			level = maxRisk(level, "high")
			signals = append(signals, "出口 IP 有滥用/攻击历史记录，风控评分高")
		}
		if data.IPRisk.IsVPN {
			level = maxRisk(level, "warning")
			signals = append(signals, "出口 IP 被识别为 VPN，部分 AI 服务会限制此类访问")
		}
		if data.IPRisk.IsDatacen && !data.IPRisk.IsVPN && !data.IPRisk.IsProxy {
			level = maxRisk(level, "warning")
			signals = append(signals, "出口 IP 属于机房/IDC 地址（"+valueOr(data.IPRisk.Org, data.IPRisk.ASN)+"），比住宅 IP 更容易触发 AI 风控")
		}
	}

	// OpenAI availability is different from raw connectivity.
	if data.OpenAI.Determined {
		if data.OpenAI.UnsupportedCountry {
			level = maxRisk(level, "high")
			signals = append(signals, "OpenAI 判定当前出口 IP 所在国家/地区不受支持（unsupported_country），无法正常使用")
		}
	}

	// Geography consistency.
	if !data.Geo.Consistent {
		level = maxRisk(level, "warning")
		signals = append(signals, data.Geo.Signals...)
	}

	for _, item := range data.DNS {
		if !item.OK {
			level = maxRisk(level, "warning")
			signals = append(signals, "DNS 解析失败: "+item.Host)
		}
	}
	for _, item := range data.Connectivity {
		if item.Blocked {
			level = maxRisk(level, "high")
			signals = append(signals, "目标站点返回 IP 拦截页: "+item.Name)
		} else if !item.Reachable {
			level = maxRisk(level, "warning")
			signals = append(signals, "OpenAI 相关连通性失败: "+item.Name)
		}
	}

	if len(signals) == 0 {
		signals = append(signals, "出口 IP、IP 画像、OpenAI 可用性、DNS、连通性均正常")
	}
	return riskSummary{
		Level:   level,
		Label:   riskLabel(level),
		Signals: signals,
		Note:    "风险画像来自第三方 IP 风控接口与 OpenAI 侧信号，结果依赖接口可用性，供参考不构成绝对结论。",
	}
}

// distinctIPsByFamily deduplicates detected egress IPs by address family.
func distinctIPsByFamily(checks []publicIPEndpoint) (v4 []string, v6 []string) {
	seen4 := make(map[string]struct{})
	seen6 := make(map[string]struct{})
	for _, check := range checks {
		if check.IP == "" {
			continue
		}
		parsed := net.ParseIP(check.IP)
		if parsed == nil {
			continue
		}
		if parsed.To4() != nil {
			if _, ok := seen4[check.IP]; !ok {
				seen4[check.IP] = struct{}{}
				v4 = append(v4, check.IP)
			}
			continue
		}
		if _, ok := seen6[check.IP]; !ok {
			seen6[check.IP] = struct{}{}
			v6 = append(v6, check.IP)
		}
	}
	return v4, v6
}

func riskLabel(level string) string {
	switch level {
	case "high":
		return "存在高风险信号"
	case "warning":
		return "存在需关注的信号"
	case "unknown":
		return "部分检测无法确认"
	default:
		return "未发现明显风险"
	}
}

func maxRisk(current, next string) string {
	order := map[string]int{"low": 1, "unknown": 2, "warning": 3, "high": 4}
	if order[next] > order[current] {
		return next
	}
	return current
}
