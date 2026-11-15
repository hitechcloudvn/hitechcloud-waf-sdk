# HiTechCloud WAF MCP Server

> [English](README.md) | **Tiếng Việt**

---

Một máy chủ [Model Context Protocol](https://modelcontextprotocol.io) cung cấp
các thao tác quản trị cốt lõi của HiTechCloud WAF cho AI agent.

## Cài đặt

```bash
cd mcp && npm install
export WAF_BASE_URL=https://vcwaf.hitechcloud.vn
export WAF_TOKEN=YOUR_JWT
npm start
```

## Công cụ

| Công cụ | Endpoint |
|------|----------|
| `waf_list_sites` | `GET /api/open/site` |
| `waf_get_site` | `GET /api/open/site/{id}` |
| `waf_list_policies` | `GET /api/open/policy` |
| `waf_list_events` | `GET /api/open/events` |
| `waf_qps` | `GET /api/stat/qps` |
| `waf_attack_stats` | `GET /api/stat/advance/attack` |
| `waf_security_posture` | `GET /api/open/security_posture/realtime` |
| `waf_list_certs` | `GET /api/open/cert` |
| `waf_license_status` | `GET /api/open/system/license/status` |
| `waf_acl_relieve` | `PUT /api/acl/relieve` |

Toàn bộ API (227 endpoint) sẵn sàng qua các SDK ngôn ngữ tại `sdk/`.
