# Rate limiting & pagination

> **English** | [Tiếng Việt](rate-limiting.vi.md)

---

## Pagination

List endpoints support the following query parameters:

| Parameter | Type | Description |
|---------|------|-------|
| `page` | integer | Page number, starting from 1 |
| `page_size` | integer | Records per page |
| `keyword` | string | Search keyword (where supported) |

## Rate limiting

When a threshold is exceeded, the API returns HTTP `429` with a `Retry-After` header. The SDK
retries automatically (exponential backoff for 5xx, honoring `Retry-After` for 429) up to the
configured number of attempts (`max_retries`, default 3). If it still fails, the SDK throws
`RateLimitError` with a `retry_after` attribute.

```python
client = WafClient(username="admin", password="********", max_retries=5)
```
