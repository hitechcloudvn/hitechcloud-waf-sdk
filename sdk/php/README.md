# HiTechCloud WAF SDK (PHP)

> **English** | [Tiếng Việt](README.vi.md)

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

Namespace: `HiTechCloud\WAF`. Pass `['sandbox' => true]` for `https://vcwaf-sandbox.hitechcloud.vn`.
