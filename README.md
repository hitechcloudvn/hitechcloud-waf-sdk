# HiTechCloud WAF SDK

[![CI](https://github.com/hitechcloudvn/hitechcloud-waf-sdk/actions/workflows/ci.yml/badge.svg)](https://github.com/hitechcloudvn/hitechcloud-waf-sdk/actions/workflows/ci.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

> Bộ SDK đa ngôn ngữ chính thức cho API quản trị **HiTechCloud WAF** (tường lửa ứng dụng web) — 227 endpoint, 25 nhóm tài nguyên.

HiTechCloud WAF bảo vệ website khỏi SQLi, XSS, dò quét, DDoS lớp 7 và bot độc hại
qua mô hình reverse proxy. Toàn bộ năng lực quản trị (site, policy, Skynet AI,
chống giả mạo, ACL/rate-limit, challenge, waiting-room, JA4, chống bot, SSL, sự kiện
tấn công, thống kê, báo cáo, cảnh báo, license...) đều mở ra qua REST API.

- Host: `https://vcwaf.hitechcloud.vn` · Base: `/api` · Sandbox: `https://vcwaf-sandbox.hitechcloud.vn`
- Tài liệu: https://developers.hitechcloud.vn

## Ngôn ngữ hỗ trợ

| Ngôn ngữ | Package | Trạng thái |
|----------|---------|------------|
| Node.js / TypeScript | `@hitechcloudvn/waf-sdk` (npm) | ✅ |
| Python | `hitechcloud-waf` (PyPI) | ✅ |
| PHP | `hitechcloudvn/waf-sdk` (Composer) | ✅ |
| Go | `github.com/hitechcloudvn/hitechcloud-waf-sdk/go` | ✅ |
| Java | `vn.hitechcloud:waf-sdk` (Maven) | ✅ |

## Luồng xác thực

`GET /api/open/auth/csrf` → `POST /api/open/auth/login` (nhận `jwt`) → (tuỳ chọn) `POST /api/open/auth/tfa` → gửi `Authorization: Bearer <jwt>`.

Hàm `client.login(username, password)` của mỗi SDK tự động hoá toàn bộ luồng này.

## Bắt đầu nhanh

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

## Nhóm tài nguyên

| Nhóm | Số endpoint | Mô tả |
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

## Cấu trúc kho

```
sdk/        # 5 SDK: nodejs (TS), python, php, go, java
docs/       # Tài liệu tiếng Việt
examples/   # Ví dụ cho từng ngôn ngữ
mcp/        # MCP server (Node) cho AI agent
openapi.json
```

## Giấy phép

MIT — xem [LICENSE](LICENSE).

## Hỗ trợ

- Email: support@hitechcloud.vn
- Issues: https://github.com/hitechcloudvn/hitechcloud-waf-sdk/issues
