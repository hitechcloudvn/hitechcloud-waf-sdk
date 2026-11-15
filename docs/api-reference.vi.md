# Tham chiếu API

> [English](api-reference.md) | **Tiếng Việt**

---

SDK tổ chức 227 endpoint thành 25 nhóm tài nguyên. Mỗi nhóm là một thuộc
tính trên `client` (ví dụ `client.sites`, `client.policies`).

| Nhóm | Số endpoint | Truy cập |
|------|-------------|----------|
| Authentication & Token | 7 | `client.auth` |
| Users & Account | 7 | `client.users` |
| Dynamic Authentication Defense | 25 | `client.authDefense` |
| Sites (Web Services) | 41 | `client.sites` |
| Policies & Rules | 9 | `client.policies` |
| Skynet AI Rules | 4 | `client.skynet` |
| Anti-Tamper | 8 | `client.antiTamper` |
| ACL & Rate-Limit | 7 | `client.acl` |
| Challenge & Semantics | 7 | `client.challenge` |
| Waiting Room | 3 | `client.waitingRoom` |
| JA4 Fingerprint | 2 | `client.ja4` |
| Global Mode & Proxy | 6 | `client.globalSettings` |
| Cloud Policies | 2 | `client.cloud` |
| SSL Certificates | 4 | `client.ssl` |
| Attacks & Events | 7 | `client.events` |
| Security Posture | 4 | `client.securityPosture` |
| Statistics | 11 | `client.statistics` |
| Reports | 4 | `client.report` |
| Alarms | 3 | `client.alarm` |
| Portal & Branding | 13 | `client.portal` |
| Anti-Bot & Intelligence | 19 | `client.antiBot` |
| Settings | 10 | `client.settings` |
| System | 15 | `client.system` |
| License & Authorization | 7 | `client.license` |
| MCP | 2 | `client.mcp` |

## Quy ước phương thức

- `GET` -> nhận tham số truy vấn tuỳ chọn (`params` / `Map` / `array`).
- `POST` / `PUT` / `DELETE` -> nhận body (`data`).
- Tham số đường dẫn (`{id}`, `{rule_id}`, ...) là tham số vị trí đầu tiên.

## Định dạng phản hồi

Mọi endpoint trả về cấu trúc JSON thống nhất:

```json
{ "data": {}, "err": "", "msg": "ok" }
```

Luôn kiểm tra `err` trước khi đọc `data`. SDK tự động ném ngoại lệ khi HTTP >= 400.

Xem đầy đủ tham số từng endpoint trong [openapi.json](../openapi.json).
