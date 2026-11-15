<?php

// EN: Smoke tests - keep the suite green so CI can gate releases.
// VI: Smoke test - giu bo test xanh de CI kiem soat viec phat hanh.

namespace HiTechCloud\WAF\Tests;

use PHPUnit\Framework\TestCase;

final class SmokeTest extends TestCase
{
    public function testMainClientIsLoadable(): void
    {
        $this->assertTrue(class_exists(\HiTechCloud\WAF\WafClient::class));
    }

    public function testExceptionHierarchyIsLoadable(): void
    {
        $this->assertTrue(class_exists(\HiTechCloud\WAF\WafException::class));
        $this->assertTrue(class_exists(\HiTechCloud\WAF\AuthenticationException::class));
        $this->assertTrue(is_subclass_of(\HiTechCloud\WAF\AuthenticationException::class, \HiTechCloud\WAF\WafException::class));
    }
}
