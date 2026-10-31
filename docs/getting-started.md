# Bắt đầu nhanh

HiTechCloud WAF là tường lửa ứng dụng web (WAF). SDK này bọc toàn bộ API quản trị
tại `https://vcwaf.hitechcloud.vn/api` theo hướng đối tượng, nhóm theo tài nguyên.

## Cài đặt

### Python
```bash
pip install hitechcloud-waf
```

### Node.js / TypeScript
```bash
npm install @hitechcloud-vietnam/waf-sdk
```

### PHP
```bash
composer require hitechcloud-vietnam/waf-sdk
```

### Go
```bash
go get github.com/hitechcloud-vietnam/hitechcloud-waf-sdk/go
```

### Java (Maven)
```xml
<dependency>
    <groupId>vn.hitechcloud</groupId>
    <artifactId>waf-sdk</artifactId>
    <version>1.0.0</version>
</dependency>
```

## Gọi API đầu tiên

```python
from hitechcloud_waf import WafClient

client = WafClient(username="admin", password="********")  # CSRF -> login -> JWT
print(client.sites.list())
```

Dùng môi trường thử nghiệm bằng `WafClient(sandbox=True, ...)` để trỏ tới `https://vcwaf-sandbox.hitechcloud.vn`.

## Môi trường

| Môi trường | Base URL |
|------------|----------|
| Production | `https://vcwaf.hitechcloud.vn/api` |
| Sandbox | `https://vcwaf-sandbox.hitechcloud.vn/api` |
