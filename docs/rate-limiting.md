# Giới hạn tần suất & Phân trang

## Phân trang

Các endpoint danh sách hỗ trợ tham số truy vấn:

| Tham số | Kiểu | Mô tả |
|---------|------|-------|
| `page` | integer | Số trang, bắt đầu từ 1 |
| `page_size` | integer | Số bản ghi mỗi trang |
| `keyword` | string | Từ khoá tìm kiếm (nếu hỗ trợ) |

## Giới hạn tần suất

Khi vượt ngưỡng, API trả về HTTP `429` kèm header `Retry-After`. SDK tự động thử lại
(theo cấp số nhân cho 5xx, theo `Retry-After` cho 429) tới số lần cấu hình
(`max_retries`, mặc định 3). Khi vẫn thất bại, SDK ném `RateLimitError` với thuộc
tính `retry_after`.

```python
client = WafClient(username="admin", password="********", max_retries=5)
```
