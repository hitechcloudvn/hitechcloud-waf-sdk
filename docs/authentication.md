# Authentication

> **English** | [Tiếng Việt](authentication.vi.md)

---

WAF uses JWT-based authentication. The standard flow has three steps (the SDK automates them
in `client.login(...)`):

| Step | Endpoint | Description |
|------|----------|-------|
| 1. CSRF | `GET /api/open/auth/csrf` | Obtain `csrf_token` |
| 2. Login | `POST /api/open/auth/login` | Send `username`, `password`, `csrf_token`; receive `jwt` |
| 3. TFA (if enabled) | `POST /api/open/auth/tfa` | Verify the second-factor OTP code |
| API token | `GET/PUT/DELETE /api/open/auth/token` | Obtain / rotate / revoke tokens |
| Logout | `POST /api/open/auth/logout` | End the session |

Once you have a token, every call sends the `Authorization: Bearer <jwt>` header.

```python
client = WafClient(username="admin", password="********", tfa_code="123456")
# or use an existing token:
client = WafClient(token="EXISTING_JWT")
```

```bash
curl "https://vcwaf.hitechcloud.vn/api/open/auth/csrf"
curl -X POST "https://vcwaf.hitechcloud.vn/api/open/auth/login" \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"********","csrf_token":"..."}'
curl "https://vcwaf.hitechcloud.vn/api/open/site" -H "Authorization: Bearer YOUR_TOKEN"
```

> API tokens carry full WAF administration privileges. Store them safely, rotate them
> periodically and revoke them immediately if you suspect a leak. Never embed tokens
> in frontend code.
