# Xác thực

> [English](authentication.md) | **Tiếng Việt**

---

WAF dùng xác thực dựa trên JWT. Luồng chuẩn gồm ba bước (SDK tự động hoá trong
`client.login(...)`):

| Bước | Endpoint | Mô tả |
|------|----------|-------|
| 1. CSRF | `GET /api/open/auth/csrf` | Lấy `csrf_token` |
| 2. Đăng nhập | `POST /api/open/auth/login` | Gửi `username`, `password`, `csrf_token`; nhận `jwt` |
| 3. TFA (nếu bật) | `POST /api/open/auth/tfa` | Xác thực mã OTP 2 lớp |
| Token API | `GET/PUT/DELETE /api/open/auth/token` | Lấy / xoay / thu hồi token |
| Đăng xuất | `POST /api/open/auth/logout` | Kết thúc phiên |

Sau khi có token, mọi lời gọi gửi header `Authorization: Bearer <jwt>`.

```python
client = WafClient(username="admin", password="********", tfa_code="123456")
# hoặc dùng token có sẵn:
client = WafClient(token="EXISTING_JWT")
```

```bash
curl "https://vcwaf.hitechcloud.vn/api/open/auth/csrf"
curl -X POST "https://vcwaf.hitechcloud.vn/api/open/auth/login" \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"********","csrf_token":"..."}'
curl "https://vcwaf.hitechcloud.vn/api/open/site" -H "Authorization: Bearer YOUR_TOKEN"
```

> Token API có toàn quyền quản trị WAF. Lưu an toàn, luân chuyển định kỳ, thu hồi
> ngay khi nghi ngờ lộ. Không nhúng token vào mã frontend.
