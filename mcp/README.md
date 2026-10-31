# HiTechCloud WAF MCP Server

A [Model Context Protocol](https://modelcontextprotocol.io) server that exposes
core HiTechCloud WAF management actions to AI agents.

## Setup

```bash
cd mcp && npm install
export WAF_BASE_URL=https://vcwaf.hitechcloud.vn
export WAF_TOKEN=YOUR_JWT
npm start
```

## Tools

| Tool | Endpoint |
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

The full API (227 endpoints) is available through the language SDKs under `sdk/`.
