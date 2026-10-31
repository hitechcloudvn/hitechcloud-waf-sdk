<?php

namespace HiTechCloud\WAF\Resources;

/**
 * Settings: Login methods, syslog, detector, downgrade and audit log.
 */
class Settings extends BaseResource
{
    /** PUT /business/system/login_method - Set the admin login method */
    public function setLoginMethod(?array $data = null): ?array
    {
        return $this->http->put('/business/system/login_method', $data);
    }

    /** POST /business/system/login_method/test - Test the admin login method */
    public function testLoginMethod(?array $data = null): ?array
    {
        return $this->http->post('/business/system/login_method/test', $data);
    }

    /** GET /commercial/syslog - Get syslog config */
    public function getSyslog(array $params = []): ?array
    {
        return $this->http->get('/commercial/syslog', $params);
    }

    /** POST /commercial/syslog - Set syslog config */
    public function setSyslog(?array $data = null): ?array
    {
        return $this->http->post('/commercial/syslog', $data);
    }

    /** POST /commercial/syslog/test - Test syslog delivery */
    public function testSyslog(?array $data = null): ?array
    {
        return $this->http->post('/commercial/syslog/test', $data);
    }

    /** GET /open/Commercial/downgrade - Get downgrade config */
    public function getDowngrade(array $params = []): ?array
    {
        return $this->http->get('/open/Commercial/downgrade', $params);
    }

    /** PUT /open/Commercial/downgrade - Set downgrade config */
    public function setDowngrade(?array $data = null): ?array
    {
        return $this->http->put('/open/Commercial/downgrade', $data);
    }

    /** GET /open/audit_log - Get the admin audit log */
    public function auditLog(array $params = []): ?array
    {
        return $this->http->get('/open/audit_log', $params);
    }

    /** GET /open/detector - Get detector config */
    public function getDetector(array $params = []): ?array
    {
        return $this->http->get('/open/detector', $params);
    }

    /** POST /open/detector - Set detector config */
    public function setDetector(?array $data = null): ?array
    {
        return $this->http->post('/open/detector', $data);
    }

}