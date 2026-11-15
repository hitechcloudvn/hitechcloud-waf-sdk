<?php

// HiTechCloud WAF SDK — PHP example / Ví dụ PHP
//
// Run: / Cách chạy:
//   composer install -d ../../sdk/php
//   php basic_usage.php

require __DIR__ . '/../../sdk/php/vendor/autoload.php';

use HiTechCloud\WAF\WafClient;
use HiTechCloud\WAF\WafException;

$client = new WafClient([
    'username' => 'admin',
    'password' => '********',
]);

try {
    // List sites / Liệt kê site
    $sites = $client->sites->list(['page' => 1, 'page_size' => 20]);
    print_r($sites);

    // Create an ACL rule / Tạo quy tắc ACL
    $client->acl->create('123', ['type' => 'ip', 'value' => '203.0.113.10', 'action' => 'deny']);

    // Read statistics and license / Đọc thống kê và giấy phép
    print_r($client->statistics->qps());
    print_r($client->license->status());
} catch (WafException $e) {
    // Error handling / Xử lý lỗi
    echo "API error {$e->getCode()}: {$e->getMessage()}\n";
}
