package main

func renderDashboardHTML() string {
	return `<!doctype html>
<html lang="zh-CN">
<head>
<meta charset="utf-8">
<meta name="viewport" content="width=device-width, initial-scale=1">
<title>CPA 网络诊断</title>
<style>
:root{color-scheme:light;--bg:#f5f7fb;--panel:#ffffff;--text:#1f2937;--muted:#6b7280;--line:#e5e7eb;--line-strong:#d1d5db;--primary:#1677ff;--primary-hover:#0958d9;--good:#15803d;--warn:#b45309;--bad:#dc2626;--soft-blue:#eff6ff;--soft-green:#ecfdf3;--soft-red:#fef2f2;--soft-amber:#fffbeb;--shadow:0 1px 2px rgba(16,24,40,.04)}
@media (prefers-color-scheme:dark){:root{color-scheme:dark;--bg:#111827;--panel:#1f2937;--text:#f9fafb;--muted:#9ca3af;--line:#374151;--line-strong:#4b5563;--primary:#4096ff;--primary-hover:#69b1ff;--soft-blue:#102a43;--soft-green:#0f2f24;--soft-red:#3b1115;--soft-amber:#33260a;--shadow:none}}
*{box-sizing:border-box}body{margin:0;background:var(--bg);color:var(--text);font:14px/1.55 system-ui,-apple-system,BlinkMacSystemFont,"Segoe UI",sans-serif}main{width:min(1180px,calc(100% - 40px));margin:0 auto;padding:24px 0 36px}.pageHeader{display:flex;align-items:flex-start;justify-content:space-between;gap:16px;margin-bottom:16px}.pageTitle{font-size:24px;line-height:1.25;margin:0 0 6px;font-weight:650}.pageDesc{margin:0;color:var(--muted)}.toolbar{display:flex;gap:8px;flex-wrap:wrap;justify-content:flex-end}.btn{border:1px solid var(--line-strong);background:var(--panel);color:var(--text);height:34px;padding:0 12px;border-radius:6px;text-decoration:none;display:inline-flex;align-items:center;gap:6px;font-size:13px;cursor:pointer}.btn:hover{border-color:var(--primary);color:var(--primary)}.btnPrimary{border-color:var(--primary);background:var(--primary);color:#fff}.btnPrimary:hover{background:var(--primary-hover);border-color:var(--primary-hover);color:#fff}.btn[disabled]{opacity:.58;cursor:not-allowed}.summary{background:var(--panel);border:1px solid var(--line);border-radius:8px;padding:18px;box-shadow:var(--shadow);margin-bottom:14px}.summaryGrid{display:grid;grid-template-columns:repeat(auto-fit,minmax(160px,1fr));gap:16px}.label{color:var(--muted);font-size:12px;margin-bottom:6px}.ip{font-size:30px;line-height:1.15;font-weight:700;overflow-wrap:anywhere}.value{font-size:15px;overflow-wrap:anywhere}.grid{display:grid;grid-template-columns:repeat(2,minmax(0,1fr));gap:14px}.panel{background:var(--panel);border:1px solid var(--line);border-radius:8px;padding:14px;box-shadow:var(--shadow)}.panelHead{display:flex;align-items:center;justify-content:space-between;gap:8px;margin-bottom:10px}.panel h2{font-size:15px;margin:0;font-weight:650}.rows{display:grid;gap:8px}.row{display:grid;grid-template-columns:150px minmax(0,1fr) auto;gap:10px;align-items:start;border-top:1px solid var(--line);padding-top:8px}.row:first-child{border-top:0;padding-top:0}.name{font-weight:600;overflow-wrap:anywhere}.meta{color:var(--muted);overflow-wrap:anywhere}.status{font-weight:650;white-space:nowrap}.ok{color:var(--good)}.warn{color:var(--warn)}.bad{color:var(--bad)}.badge{display:inline-flex;align-items:center;border-radius:999px;padding:2px 8px;font-size:12px;font-weight:650}.badgeOk{background:var(--soft-green);color:var(--good)}.badgeWarn{background:var(--soft-amber);color:var(--warn)}.badgeBad{background:var(--soft-red);color:var(--bad)}.badgeInfo{background:var(--soft-blue);color:var(--primary)}.chips{display:flex;gap:6px;flex-wrap:wrap}.chip{background:var(--soft-blue);color:var(--primary);border-radius:999px;padding:3px 8px;font-size:12px}.ipCards{display:grid;gap:10px}.ipCard{border:1px solid var(--line);border-radius:8px;padding:12px;background:rgba(148,163,184,.04)}.ipCardHead{display:flex;align-items:center;justify-content:space-between;gap:8px;margin-bottom:10px}.ipCardTitle{font-weight:650}.kvGrid{display:grid;grid-template-columns:repeat(2,minmax(0,1fr));gap:8px}.kv{min-width:0}.kvLabel{color:var(--muted);font-size:12px;margin-bottom:2px}.kvValue{overflow-wrap:anywhere}.mono{font-family:ui-monospace,SFMono-Regular,Menlo,Consolas,monospace}.small{font-size:12px;color:var(--muted);margin-top:12px}.browser{display:grid;grid-template-columns:repeat(2,minmax(0,1fr));gap:10px}.loadingBox{background:var(--panel);border:1px solid var(--line);border-radius:8px;box-shadow:var(--shadow);padding:34px 18px;margin-bottom:14px;display:flex;align-items:center;justify-content:center;gap:12px;color:var(--muted)}.spinner{width:22px;height:22px;border:3px solid var(--line);border-top-color:var(--primary);border-radius:999px;animation:spin .8s linear infinite}.skeleton{position:relative;overflow:hidden;background:linear-gradient(90deg,rgba(148,163,184,.14),rgba(148,163,184,.28),rgba(148,163,184,.14));background-size:240% 100%;animation:shimmer 1.2s ease-in-out infinite;border-radius:6px;min-height:16px}.hidden{display:none!important}.errorBox{border:1px solid #fecaca;background:var(--soft-red);color:var(--bad);border-radius:8px;padding:12px 14px;margin-bottom:14px}@keyframes spin{to{transform:rotate(360deg)}}@keyframes shimmer{to{background-position:-240% 0}}@media (max-width:900px){main{width:min(100% - 20px,1180px);padding-top:18px}.pageHeader{display:block}.toolbar{justify-content:flex-start;margin-top:12px}.summaryGrid,.grid,.browser,.kvGrid{grid-template-columns:1fr}.row{grid-template-columns:1fr}.status{white-space:normal}.ip{font-size:25px}}
</style>
</head>
<body>
<main>
  <section class="pageHeader">
    <div>
      <h1 class="pageTitle">CPA 网络诊断</h1>
      <p class="pageDesc">检测位置：CPA 插件进程所在环境。无论主机直装、Docker 还是云容器部署，这里显示的都是实际运行环境看到的出口状态。</p>
    </div>
    <div class="toolbar">
      <button class="btn btnPrimary" id="refreshBtn" type="button">重新检测</button>
      <a class="btn" href="/v0/management/diagnostics/status">JSON API</a>
    </div>
  </section>

  <div id="loading" class="loadingBox">
    <span class="spinner" aria-hidden="true"></span>
    <span>正在检测部署环境的出口 IP、DNS 和 OpenAI 连通性...</span>
  </div>
  <div id="error" class="errorBox hidden"></div>
  <div id="content" class="hidden"></div>
</main>
<script>
const statusUrl = '/v0/resource/plugins/diagnostics/status';
const loading = document.getElementById('loading');
const errorBox = document.getElementById('error');
const content = document.getElementById('content');
const refreshBtn = document.getElementById('refreshBtn');
var currentRun = 0;
var state = {};
var fullStatusPromise = null;

refreshBtn.addEventListener('click', function(){ runDiagnostics(); });
runDiagnostics();

async function runDiagnostics(){
  loading.classList.remove('hidden');
  errorBox.classList.add('hidden');
  content.classList.add('hidden');
  refreshBtn.disabled = true;
  try {
    const resp = await fetch(statusUrl + '?t=' + Date.now(), {cache:'no-store'});
    if (!resp.ok) throw new Error('HTTP ' + resp.status);
    const data = await resp.json();
    content.innerHTML = render(data);
    content.classList.remove('hidden');
    renderBrowserInfo();
  } catch (err) {
    errorBox.textContent = '诊断加载失败：' + (err && err.message ? err.message : String(err));
    errorBox.classList.remove('hidden');
  } finally {
    loading.classList.add('hidden');
    refreshBtn.disabled = false;
  }
}

function render(data){
  const pub = data.public_ip || {};
  const risk = data.risk || {};
  const ipRisk = data.ip_risk || {};
  const openai = data.openai || {};
  return '<section class="summary">' +
    '<div class="summaryGrid">' +
      metric('公共出口 IP', pub.ip || '未获取', 'ip mono') +
      metric('国家/地区', pub.country || '未知', 'value') +
      metric('IP 类型', ipTypeBadge(ipRisk), 'value raw') +
      metric('OpenAI 可用性', openaiBadge(openai), 'value raw') +
      metric('运营商/组织', pub.org || ipRisk.org || '未知', 'value') +
      metric('风险概览', badge(risk.level, risk.label || '未知'), 'value raw') +
    '</div>' +
    '<div class="small">检测时间：' + esc(data.checked_at || '-') + '，耗时 ' + esc(data.duration_ms || 0) + ' ms，来源：' + esc(pub.source || '无') + '。</div>' +
  '</section>' +
  '<section class="grid">' +
    panel('风险信号', renderRisk(risk)) +
    panel('IP 风险画像', renderIPRisk(ipRisk)) +
    panel('OpenAI 可用性', renderOpenAI(openai)) +
    panel('地区一致性', renderGeo(data.geo_consistency || {})) +
    panel('运行环境', renderRuntime(data.runtime || {})) +
    panel('本机 IP', renderLocalIPs(data.local_ips || [])) +
    panel('出口源地址', renderOutbound(data.outbound_sources || [])) +
    panel('公共 IP 查询', renderPublicChecks((pub.checks || []))) +
    panel('DNS 解析', renderDNS(data.dns || [])) +
    panel('OpenAI 连通性', renderConnectivity(data.connectivity || [])) +
    panel('浏览器信息', '<div id="browser" class="browser"></div><div class="small">浏览器信息来自当前页面，用来对比访问者环境和 CPA 进程环境。</div>') +
  '</section>';
}

function metric(label, value, cls){
  const raw = cls && cls.indexOf('raw') >= 0;
  return '<div><div class="label">' + esc(label) + '</div><div class="' + esc(cls || 'value') + '">' + (raw ? value : esc(value)) + '</div></div>';
}
function panel(title, body){
  return '<div class="panel"><div class="panelHead"><h2>' + esc(title) + '</h2></div>' + body + '</div>';
}
function rows(items){
  return items.length ? '<div class="rows">' + items.join('') + '</div>' : '<div class="meta">暂无数据</div>';
}
function row(name, meta, right){
  return '<div class="row"><div class="name">' + esc(name) + '</div><div class="meta mono">' + esc(meta || '未知') + '</div><div>' + (right || '') + '</div></div>';
}
function renderRisk(risk){
  const signals = risk.signals || [];
  return rows(signals.map(function(signal){ return row('信号', signal, ''); })) +
    '<div class="small">' + esc(risk.note || '这是基础网络可达性与出口一致性检测，不等同于专业 IP 风控评分。') + '</div>';
}
function renderRuntime(info){
  return rows([
    row('Hostname', info.hostname || '-', ''),
    row('OS / Arch', compact([info.goos, info.goarch], ' / ') || '-', ''),
    row('时区', compact([info.timezone_name, info.timezone_utc], ' ') || '-', ''),
    row('PID', String(info.pid || '-'), '')
  ]);
}
function renderLocalIPs(items){
  return rows(items.map(function(item){
    const tags = [item.version || 'IP'];
    if (item.private) tags.push('private');
    if (item.loopback) tags.push('loopback');
    return row(item.interface || '-', item.address || '-', chips(tags));
  }));
}
function renderOutbound(items){
  return rows(items.map(function(item){
    return row(item.target || '-', item.local_ip || item.error || '-', status(item.ok, item.latency_ms));
  }));
}
function renderPublicChecks(items){
  if (!items.length) return '<div class="meta">暂无数据</div>';
  return '<div class="ipCards">' + items.map(function(item){
    return '<div class="ipCard">' +
      '<div class="ipCardHead"><div class="ipCardTitle">' + esc(item.name || '查询源') + '</div>' + status(item.ok, item.latency_ms) + '</div>' +
      '<div class="kvGrid">' +
        kv('IP 地址', item.ip || '未获取', true) +
        kv('国家/地区', item.country || '未知') +
        kv('地区', item.region || '未知') +
        kv('城市', item.city || '未知') +
        kv('运营商/组织', item.org || '未知') +
        kv('接口地址', item.url || '-', true) +
      '</div>' +
      (item.error ? '<div class="small bad">说明：' + esc(item.error) + '</div>' : '') +
    '</div>';
  }).join('') + '</div>';
}
function renderDNS(items){
  return rows(items.map(function(item){
    return row(item.host || '-', (item.addresses || []).join(', ') || item.error || '-', status(item.ok, item.latency_ms));
  }));
}
function renderConnectivity(items){
  return rows(items.map(function(item){
    const meta = item.status_code ? ('HTTP ' + item.status_code + ' | ' + (item.expected_note || '')) : (item.error || '-');
    const right = item.blocked ? '<span class="status bad">被拦截' + (item.latency_ms || item.latency_ms === 0 ? ' · ' + esc(item.latency_ms) + ' ms' : '') + '</span>' : status(item.reachable, item.latency_ms);
    return row(item.name || '-', meta, right);
  }));
}
function renderIPRisk(ip){
  if (!ip.determined) return '<div class="meta">未能确定 IP 画像（风控接口不可达或未获取到出口 IP）。</div>';
  const flags = [];
  if (ip.is_datacenter) flags.push('机房/IDC');
  if (ip.is_proxy) flags.push('代理');
  if (ip.is_vpn) flags.push('VPN');
  if (ip.is_tor) flags.push('Tor');
  if (ip.is_abuser) flags.push('滥用历史');
  if (ip.is_mobile) flags.push('移动网络');
  return rows([
    row('IP 类型', ipTypeLabel(ip.type), ipTypeBadge(ip)),
    row('风险标记', flags.length ? flags.join('、') : '无代理/VPN/机房标记', flags.length ? '<span class="status bad">命中</span>' : '<span class="status ok">干净</span>'),
    row('ASN', ip.asn || '-', ''),
    row('组织', ip.org || '-', ''),
    row('数据来源', ip.source || '-', '')
  ]);
}
function renderOpenAI(o){
  if (!o.determined) return '<div class="meta">未能确定 OpenAI 可用性（compliance 接口与 Cloudflare trace 均不可达）。</div><div class="small">' + esc(o.note || '') + '</div>';
  const supported = o.supported && !o.unsupported_country;
  return rows([
    row('可用性结论', supported ? '当前出口 IP 可用' : '当前出口 IP 不可用/受限', supported ? '<span class="status ok">可用</span>' : '<span class="status bad">不可用</span>'),
    row('unsupported_country', o.unsupported_country ? '命中（该地区不被支持）' : '未命中', o.unsupported_country ? '<span class="status bad">命中</span>' : '<span class="status ok">正常</span>'),
    row('CF 识别国家', o.cf_country || '-', ''),
    row('compliance 接口', o.compliance_ok ? '成功返回' : (o.error || '未响应'), status(o.compliance_ok, o.latency_ms))
  ]) + '<div class="small">' + esc(o.note || '') + '</div>';
}
function renderGeo(g){
  const signals = g.signals || [];
  return rows([
    row('IP 国家', g.ip_country || '-', ''),
    row('CF 识别国家', g.cf_country || '-', ''),
    row('进程时区', compact([g.timezone_name, g.timezone_utc], ' ') || '-', g.consistent ? '<span class="status ok">一致</span>' : '<span class="status bad">不一致</span>')
  ]) + '<div class="rows" style="margin-top:8px">' + signals.map(function(s){ return row('信号', s, ''); }).join('') + '</div>' +
    '<div class="small">进程时区用于辅助判断部署环境，当前版本不把时区直接计入一致性结论。</div>';
}
function ipTypeLabel(t){
  const map = {hosting:'机房 / 数据中心', residential:'住宅宽带', mobile:'移动网络', business:'商业宽带', unknown:'未知'};
  return map[t] || '未知';
}
function ipTypeBadge(ip){
  if (!ip || !ip.determined) return '<span class="badge badgeInfo">未知</span>';
  const t = ip.type;
  let cls = 'badgeInfo';
  if (t === 'residential' || t === 'mobile' || t === 'business') cls = 'badgeOk';
  if (t === 'hosting') cls = 'badgeBad';
  return '<span class="badge ' + cls + '">' + esc(ipTypeLabel(t)) + '</span>';
}
function openaiBadge(o){
  if (!o || !o.determined) return '<span class="badge badgeInfo">未知</span>';
  const supported = o.supported && !o.unsupported_country;
  return '<span class="badge ' + (supported ? 'badgeOk' : 'badgeBad') + '">' + (supported ? '可用' : '不可用') + '</span>';
}
function renderBrowserInfo(){
  const browser = document.getElementById('browser');
  if (!browser) return;
  const items = [
    ['语言', navigator.language || '未知'],
    ['平台', navigator.platform || '未知'],
    ['时区', Intl.DateTimeFormat().resolvedOptions().timeZone || '未知'],
    ['User Agent', navigator.userAgent || '未知'],
    ['页面地址', location.href]
  ];
  browser.innerHTML = items.map(function(item){
    return '<div><div class="label">' + esc(item[0]) + '</div><div class="value mono">' + esc(item[1]) + '</div></div>';
  }).join('');
}
function status(ok, ms){
  return '<span class="status ' + (ok ? 'ok' : 'bad') + '">' + (ok ? '正常' : '失败') + (ms || ms === 0 ? ' · ' + esc(ms) + ' ms' : '') + '</span>';
}
function kv(label, value, mono){
  return '<div class="kv"><div class="kvLabel">' + esc(label) + '</div><div class="kvValue ' + (mono ? 'mono' : '') + '">' + esc(value) + '</div></div>';
}
function badge(level, label){
  let cls = 'badgeInfo';
  if (level === 'low') cls = 'badgeOk';
  if (level === 'warning') cls = 'badgeWarn';
  if (level === 'high' || level === 'unknown') cls = 'badgeBad';
  return '<span class="badge ' + cls + '">' + esc(label || level || '未知') + '</span>';
}
function chips(values){
  return '<span class="chips">' + values.map(function(value){ return '<span class="chip">' + esc(value) + '</span>'; }).join('') + '</span>';
}
function compact(values, sep){
  return values.filter(function(v){ return v !== undefined && v !== null && String(v).trim() !== ''; }).join(sep);
}
function esc(value){
  return String(value == null ? '' : value).replace(/[&<>"']/g, function(ch){
    return {'&':'&amp;','<':'&lt;','>':'&gt;','"':'&quot;',"'":'&#39;'}[ch];
  });
}

// Progressive dashboard override. Later function declarations intentionally replace the initial simple loader.
async function runDiagnostics(){
  const runId = ++currentRun;
  const started = Date.now();
  state = {checked_at: new Date().toISOString()};
  fullStatusPromise = null;
  errorBox.classList.add('hidden');
  loading.classList.add('hidden');
  content.innerHTML = renderShell();
  content.classList.remove('hidden');
  refreshBtn.disabled = true;
  renderBrowserInfo();
  updateSummary();

  const tasks = [
    loadSection(runId, 'runtimePanel', statusUrl + '/runtime', function(data){
      state.runtime = data.runtime || {};
      state.local_ips = data.local_ips || [];
      state.proxy = data.proxy || {};
      setPanel('runtimePanel', renderRuntime(state.runtime));
      setPanel('localPanel', renderLocalIPs(state.local_ips));
      setPanel('proxyPanel', renderProxy(state.proxy));
    }),
    loadSection(runId, 'publicPanel', statusUrl + '/public-ip', function(data){
      state.public_ip = data || {};
      setPanel('publicPanel', renderPublicChecks((state.public_ip.checks || [])));
      updateSummary();
      updateRiskPanel();
    }),
    loadSection(runId, 'ipRiskPanel', statusUrl + '/ip-risk', function(data){
      if (!state.public_ip || !state.public_ip.ip) state.public_ip = data.public_ip || state.public_ip || {};
      state.ip_risk = data.ip_risk || {};
      setPanel('ipRiskPanel', renderIPRisk(state.ip_risk));
      updateSummary();
      updateRiskPanel();
    }),
    loadSection(runId, 'openaiPanel', statusUrl + '/openai', function(data){
      state.openai = data || {};
      setPanel('openaiPanel', renderOpenAI(state.openai));
      updateSummary();
      updateRiskPanel();
    }),
    loadSection(runId, 'geoPanel', statusUrl + '/geo', function(data){
      state.geo_consistency = data || {};
      setPanel('geoPanel', renderGeo(state.geo_consistency));
      updateRiskPanel();
    }),
    loadSection(runId, 'dnsPanel', statusUrl + '/dns', function(data){
      state.dns = data || [];
      setPanel('dnsPanel', renderDNS(state.dns));
      updateRiskPanel();
    }),
    loadSection(runId, 'connectivityPanel', statusUrl + '/connectivity', function(data){
      state.connectivity = data || [];
      setPanel('connectivityPanel', renderConnectivity(state.connectivity));
      updateRiskPanel();
    }),
    loadSection(runId, 'outboundPanel', statusUrl + '/outbound', function(data){
      state.outbound_sources = data || [];
      setPanel('outboundPanel', renderOutbound(state.outbound_sources));
    })
  ];

  await Promise.allSettled(tasks);
  if (runId !== currentRun) return;
  state.duration_ms = Date.now() - started;
  const meta = document.getElementById('summaryMeta');
  if (meta) meta.textContent = '检测时间：' + (state.checked_at || '-') + '，耗时 ' + state.duration_ms + ' ms。结果按面板渐进加载，服务端完整 JSON 仍保留 30 秒缓存。';
  refreshBtn.disabled = false;
}

async function loadSection(runId, panelId, url, onData){
  try {
    const resp = await fetch(url + '?t=' + Date.now(), {cache:'no-store'});
    if (!resp.ok) throw new Error('HTTP ' + resp.status);
    const data = await resp.json();
    if (runId !== currentRun) return;
    onData(data);
  } catch (err) {
    if (runId !== currentRun) return;
    setPanel(panelId, '<div class="meta bad">加载失败：' + esc(err && err.message ? err.message : String(err)) + '</div>');
  }
}

function loadFullStatusFallback(){
  if (!fullStatusPromise) {
    fullStatusPromise = fetch(statusUrl + '?t=' + Date.now(), {cache:'no-store'}).then(function(resp){
      if (!resp.ok) throw new Error('HTTP ' + resp.status);
      return resp.json();
    });
  }
  return fullStatusPromise;
}
function fallbackSection(panelId, full){
  if (panelId === 'runtimePanel') return {runtime: full.runtime || {}, local_ips: full.local_ips || [], proxy: full.proxy || {}};
  if (panelId === 'publicPanel') return full.public_ip || {};
  if (panelId === 'ipRiskPanel') return {public_ip: full.public_ip || {}, ip_risk: full.ip_risk || {}};
  if (panelId === 'openaiPanel') return full.openai || {};
  if (panelId === 'geoPanel') return full.geo_consistency || {};
  if (panelId === 'dnsPanel') return full.dns || [];
  if (panelId === 'connectivityPanel') return full.connectivity || [];
  if (panelId === 'outboundPanel') return full.outbound_sources || [];
  return full;
}
function renderShell(){
  return '<section class="summary">' +
    '<div class="summaryGrid">' +
      metric('公共出口 IP', '<span id="summaryIp">检测中...</span>', 'ip mono raw') +
      metric('国家/地区', '<span id="summaryCountry">检测中...</span>', 'value raw') +
      metric('IP 类型', '<span id="summaryIPType" class="badge badgeInfo">检测中</span>', 'value raw') +
      metric('OpenAI 可用性', '<span id="summaryOpenAI" class="badge badgeInfo">检测中</span>', 'value raw') +
      metric('运营商/组织', '<span id="summaryOrg">检测中...</span>', 'value raw') +
      metric('风险概览', '<span id="summaryRisk" class="badge badgeInfo">检测中</span>', 'value raw') +
    '</div>' +
    '<div class="small" id="summaryMeta">各检测项正在并行加载，先完成的面板会先显示。</div>' +
  '</section>' +
  '<section class="grid">' +
    panelSlot('riskPanel', '风险信号') +
    panelSlot('ipRiskPanel', 'IP 风险画像') +
    panelSlot('openaiPanel', 'OpenAI 可用性') +
    panelSlot('geoPanel', '地区一致性') +
    panelSlot('runtimePanel', '运行环境') +
    panelSlot('proxyPanel', '代理环境') +
    panelSlot('localPanel', '本机 IP') +
    panelSlot('outboundPanel', '出口源地址') +
    panelSlot('publicPanel', '公共 IP 查询') +
    panelSlot('dnsPanel', 'DNS 解析') +
    panelSlot('connectivityPanel', 'OpenAI 连通性') +
    panel('浏览器信息', '<div id="browser" class="browser"></div><div class="small">浏览器信息来自当前页面，用来对比访问者环境和 CPA 进程环境。</div>') +
  '</section>';
}

function panelSlot(id, title){
  return '<div class="panel"><div class="panelHead"><h2>' + esc(title) + '</h2></div><div id="' + esc(id) + '">' + loadingRows() + '</div></div>';
}
function loadingRows(){
  return '<div class="rows"><div class="row"><div class="name">状态</div><div class="meta">正在检测...</div><div><span class="spinner" aria-hidden="true"></span></div></div></div>';
}
function setPanel(id, html){
  const el = document.getElementById(id);
  if (el) el.innerHTML = html;
}
function updateSummary(){
  const pub = state.public_ip || {};
  const ipRisk = state.ip_risk || {};
  const openai = state.openai || {};
  setText('summaryIp', pub.ip || '检测中...');
  setText('summaryCountry', pub.country || '检测中...');
  setHTML('summaryIPType', ipTypeBadge(ipRisk));
  setHTML('summaryOpenAI', openaiBadge(openai));
  setText('summaryOrg', pub.org || ipRisk.org || '检测中...');
  const risk = riskFromState();
  setHTML('summaryRisk', badge(risk.level, risk.label));
}
function updateRiskPanel(){
  setPanel('riskPanel', renderRisk(riskFromState()));
  updateSummary();
}
function riskFromState(){
  const signals = [];
  let level = 'low';
  const ipRisk = state.ip_risk || {};
  const openai = state.openai || {};
  if (state.public_ip && !state.public_ip.ip) { level = maxRiskJS(level, 'unknown'); signals.push('所有公共 IP 查询接口均失败，无法确认出口 IP'); }
  if (ipRisk.determined) {
    if (ipRisk.is_tor || ipRisk.is_proxy || ipRisk.is_abuser) { level = maxRiskJS(level, 'high'); signals.push('出口 IP 命中高风险代理/Tor/滥用信号'); }
    if (ipRisk.is_vpn || ipRisk.is_datacenter) { level = maxRiskJS(level, 'warning'); signals.push('出口 IP 被识别为 VPN 或机房地址，可能触发 AI 服务风控'); }
  }
  if (openai.unsupported_country) { level = maxRiskJS(level, 'high'); signals.push('OpenAI 返回 unsupported_country，当前地区不受支持'); }
  (state.geo_consistency && state.geo_consistency.consistent === false ? (state.geo_consistency.signals || []) : []).forEach(function(s){ level = maxRiskJS(level, 'warning'); signals.push(s); });
  (state.dns || []).forEach(function(item){ if (!item.ok) { level = maxRiskJS(level, 'warning'); signals.push('DNS 解析失败: ' + item.host); } });
  (state.connectivity || []).forEach(function(item){ if (item.blocked) { level = maxRiskJS(level, 'high'); signals.push('目标站点返回 IP 拦截页: ' + item.name); } else if (item.reachable === false) { level = maxRiskJS(level, 'warning'); signals.push('OpenAI 相关连通性失败: ' + item.name); } });
  if (!signals.length) signals.push('已完成的检测项暂未发现明显风险');
  return {level: level, label: riskLabelJS(level), signals: signals, note: '风险概览会随着各面板结果返回逐步更新。'};
}
function riskLabelJS(level){
  if (level === 'high') return '存在高风险信号';
  if (level === 'warning') return '存在需关注的信号';
  if (level === 'unknown') return '部分检测无法确认';
  return '未发现明显风险';
}
function maxRiskJS(current, next){
  const order = {low:1, unknown:2, warning:3, high:4};
  return order[next] > order[current] ? next : current;
}
function setText(id, value){ const el = document.getElementById(id); if (el) el.textContent = value; }
function setHTML(id, value){ const el = document.getElementById(id); if (el) el.innerHTML = value; }
function renderProxy(proxy){
  const vars = (proxy && proxy.variables) || [];
  if (!vars.length) return '<div class="meta">未检测到代理环境变量。</div><div class="small">CPA 仍可能通过运行环境、容器网络或上游系统代理出站。</div>';
  return rows(vars.map(function(item){ return row(item.name || '-', item.value || '(空值)', item.value ? '<span class="status warn">已设置</span>' : '<span class="status warn">空值</span>'); })) +
    '<div class="small">' + esc(proxy.note || '代理变量值已脱敏显示。') + '</div>';
}
</script>
</body>
</html>`
}
