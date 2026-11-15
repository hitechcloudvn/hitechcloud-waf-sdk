# HiTechCloud WAF SDK

> **English** | [Tiếng Việt](README.vi.md)

---

[![CI](https://github.com/hitechcloudvn/hitechcloud-waf-sdk/actions/workflows/ci.yml/badge.svg)](https://github.com/hitechcloudvn/hitechcloud-waf-sdk/actions/workflows/ci.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

> The official multi-language SDK suite for the **HiTechCloud WAF** administration API (web application firewall) — 227 endpoints, 25 resource groups.

HiTechCloud WAF protects websites from SQLi, XSS, scanning, layer-7 DDoS and malicious bots
through a reverse-proxy model. The entire administration surface (sites, policies, Skynet AI,
anti-tamper, ACL/rate-limit, challenges, waiting room, JA4, anti-bot, SSL, attack events,
statistics, reports, alarms, licensing and more) is exposed through a REST API.

- Host: `https://vcwaf.hitechcloud.vn` · Base: `/api` · Sandbox: `https://vcwaf-sandbox.hitechcloud.vn`
- Documentation: https://developers.hitechcloud.vn

## Supported languages

| Language | Package | Status |
|----------|---------|------------|
| Node.js / TypeScript | `@hitechcloudvn/waf-sdk` (npm) | ✅ |
| Python | `hitechcloud-waf` (PyPI) | ✅ |
| PHP | `hitechcloudvn/waf-sdk` (Composer) | ✅ |
| Go | `github.com/hitechcloudvn/hitechcloud-waf-sdk/go` | ✅ |
| Java | `vn.hitechcloud:waf-sdk` (Maven) | ✅ |

## Authentication flow

`GET /api/open/auth/csrf` → `POST /api/open/auth/login` (returns `jwt`) → (optional) `POST /api/open/auth/tfa` → send `Authorization: Bearer <jwt>`.

Each SDK's `client.login(username, password)` helper automates this whole flow.

## Quick start

### Node.js / TypeScript
```ts
import { WafClient } from "@hitechcloudvn/waf-sdk";
const client = new WafClient({ username: "admin", password: "********" });
await client.login("admin", "********");
console.log(await client.sites.list());
```

### Python
```python
from hitechcloud_waf import WafClient
client = WafClient(username="admin", password="********")
print(client.sites.list())
```

### PHP
```php
use HiTechCloud\WAF\WafClient;
$client = new WafClient(['username' => 'admin', 'password' => '********']);
print_r($client->sites->list());
```

### Go
```go
client, _ := waf.NewClient(waf.WithCredentials("admin", "********"))
sites, _ := client.Sites.List(nil)
```

### Java
```java
WafClient client = new WafClient(null);
client.login("admin", "********", null);
client.sites().list(null);
```

## Resource groups

| Group | Endpoints | Description |
|------|-------------|-------|
| Authentication & Token | 7 | CSRF, login, TFA, logout and API token lifecycle |
| Users & Account | 7 | Manage administrative users, TOTP and the current manager profile |
| Dynamic Authentication Defense | 25 | Authentication-defense sources, protected users, reviews and records |
| Sites (Web Services) | 41 | Protected sites, groups, proxy/origin config, logs and forwarding rules |
| Policies & Rules | 9 | Custom protection policies and rule ordering |
| Skynet AI Rules | 4 | AI-driven Skynet detection rules and switches |
| Anti-Tamper | 8 | Content anti-tamper protection per site and resource |
| ACL & Rate-Limit | 7 | Access-control / rate-limit rules and block relief |
| Challenge & Semantics | 7 | Human-verification challenge, semantic analysis and chaos protection |
| Waiting Room | 3 | Virtual waiting-room queue for traffic peaks |
| JA4 Fingerprint | 2 | JA4/JA4+ TLS client-fingerprint rules |
| Global Mode & Proxy | 6 | Global protection mode, outbound proxy and log cleanup |
| Cloud Policies | 2 | Shared cloud threat-intelligence policies |
| SSL Certificates | 4 | TLS/SSL certificate management |
| Attacks & Events | 7 | Attack events, matched rules, raw records and export |
| Security Posture | 4 | Realtime posture, per-site scoring, statistics and trends |
| Statistics | 11 | Advanced access/attack analytics and QPS |
| Reports | 4 | Scheduled / on-demand security reports |
| Alarms | 3 | Alert / notification configuration |
| Portal & Branding | 13 | Self-service portal, block pages, branding and accounts |
| Anti-Bot & Intelligence | 19 | Threat intelligence, IP groups, crawlers and behaviour/fingerprint sharing |
| Settings | 10 | Login methods, syslog, detector, downgrade and audit log |
| System | 15 | System info, edition, protocols, network proxy and helpers |
| License & Authorization | 7 | Authorization records and license status / lifecycle |
| MCP | 2 | Model Context Protocol bridge endpoints |

## Repository layout

```
sdk/        # 5 SDKs: nodejs (TS), python, php, go, java
docs/       # Vietnamese documentation
examples/   # Examples for each language
mcp/        # MCP server (Node) for AI agents
openapi.json
```

## License

MIT — see [LICENSE](LICENSE).

## Support

- Email: support@hitechcloud.vn
- Issues: https://github.com/hitechcloudvn/hitechcloud-waf-sdk/issues
