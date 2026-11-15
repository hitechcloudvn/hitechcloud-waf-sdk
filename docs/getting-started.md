# Quick start

> **English** | [Tiếng Việt](getting-started.vi.md)

---

HiTechCloud WAF is a web application firewall (WAF). This SDK wraps the entire administration
API at `https://vcwaf.hitechcloud.vn/api` in an object-oriented way, grouped by resource.

## Installation

### Python
```bash
pip install hitechcloud-waf
```

### Node.js / TypeScript
```bash
npm install @hitechcloudvn/waf-sdk
```

### PHP
```bash
composer require hitechcloudvn/waf-sdk
```

### Go
```bash
go get github.com/hitechcloudvn/hitechcloud-waf-sdk/go
```

### Java (Maven)
```xml
<dependency>
    <groupId>vn.hitechcloud</groupId>
    <artifactId>waf-sdk</artifactId>
    <version>1.0.0</version>
</dependency>
```

## Making your first API call

```python
from hitechcloud_waf import WafClient

client = WafClient(username="admin", password="********")  # CSRF -> login -> JWT
print(client.sites.list())
```

Use the sandbox environment with `WafClient(sandbox=True, ...)` to target `https://vcwaf-sandbox.hitechcloud.vn`.

## Environments

| Environment | Base URL |
|------------|----------|
| Production | `https://vcwaf.hitechcloud.vn/api` |
| Sandbox | `https://vcwaf-sandbox.hitechcloud.vn/api` |
