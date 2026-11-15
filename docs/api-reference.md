# API reference

> **English** | [Tiếng Việt](api-reference.vi.md)

---

The SDK organizes 227 endpoints into 25 resource groups. Each group is a property on `client`
(for example `client.sites`, `client.policies`).

| Group | Endpoints | Access |
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

## Method conventions

- `GET` -> accepts optional query parameters (`params` / `Map` / `array`).
- `POST` / `PUT` / `DELETE` -> accepts a body (`data`).
- Path parameters (`{id}`, `{rule_id}`, ...) are the first positional parameters.

## Response format

Every endpoint returns a uniform JSON structure:

```json
{ "data": {}, "err": "", "msg": "ok" }
```

Always check `err` before reading `data`. The SDK throws an exception automatically when HTTP >= 400.

See the full parameters of every endpoint in [openapi.json](../openapi.json).
