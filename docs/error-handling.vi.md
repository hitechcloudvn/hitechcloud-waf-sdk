# Xử lý lỗi

> [English](error-handling.md) | **Tiếng Việt**

---

SDK ánh xạ mã HTTP sang các lớp ngoại lệ chuyên biệt (kế thừa `WafError`).

| Mã | Ngoại lệ | Ý nghĩa |
|----|----------|---------|
| 401 | `AuthenticationError` | Thiếu hoặc sai token |
| 403 | `AuthorizationError` | Không đủ quyền |
| 404 | `NotFoundError` | Không tìm thấy tài nguyên |
| 422 | `ValidationError` | Dữ liệu gửi lên không hợp lệ |
| 429 | `RateLimitError` | Vượt giới hạn tần suất (có `retry_after`) |
| 5xx | `ServerError` | Lỗi phía máy chủ |

```python
from hitechcloud_waf import WafError, RateLimitError

try:
    client.sites.list()
except RateLimitError as e:
    print("Thử lại sau", e.retry_after, "giây")
except WafError as e:
    print(e.status_code, e.message, e.data)
```

Thông điệp lỗi được lấy từ trường `msg` (hoặc `err`) trong phản hồi API.
