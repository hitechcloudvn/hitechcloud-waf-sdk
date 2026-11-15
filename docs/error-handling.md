# Error handling

> **English** | [Tiếng Việt](error-handling.vi.md)

---

The SDK maps HTTP status codes to specialized exception classes (derived from `WafError`).

| Code | Exception | Meaning |
|----|----------|---------|
| 401 | `AuthenticationError` | Missing or invalid token |
| 403 | `AuthorizationError` | Insufficient permissions |
| 404 | `NotFoundError` | Resource not found |
| 422 | `ValidationError` | Invalid request payload |
| 429 | `RateLimitError` | Rate limit exceeded (has `retry_after`) |
| 5xx | `ServerError` | Server-side error |

```python
from hitechcloud_waf import WafError, RateLimitError

try:
    client.sites.list()
except RateLimitError as e:
    print("Retry after", e.retry_after, "seconds")
except WafError as e:
    print(e.status_code, e.message, e.data)
```

Error messages are taken from the `msg` (or `err`) field of the API response.
