<?php

require __DIR__ . '/../../sdk/php/vendor/autoload.php';

use HiTechCloud\WAF\WafClient;
use HiTechCloud\WAF\WafException;

$client = new WafClient([
    'username' => 'admin',
    'password' => '********',
]);

try {
    $sites = $client->sites->list(['page' => 1, 'page_size' => 20]);
    print_r($sites);

    $client->acl->create('123', ['type' => 'ip', 'value' => '203.0.113.10', 'action' => 'deny']);

    print_r($client->statistics->qps());
    print_r($client->license->status());
} catch (WafException $e) {
    echo "API error {$e->getCode()}: {$e->getMessage()}\n";
}
