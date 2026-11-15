# HiTechCloud WAF SDK (PHP)

> [English](README.md) | **Tiếng Việt**

---

```bash
composer require hitechcloudvn/waf-sdk
```

```php
use HiTechCloud\WAF\WafClient;

$client = new WafClient([
    'username' => 'admin',
    'password' => '********',
]);

$sites = $client->sites->list();
$site = $client->sites->get('123');
$client->policies->create(['name' => 'block-sqli']);
$client->statistics->qps();
```

Namespace: `HiTechCloud\WAF`. Truyền `['sandbox' => true]` để dùng `https://vcwaf-sandbox.hitechcloud.vn`.
