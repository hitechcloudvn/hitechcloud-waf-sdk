<?php

namespace HiTechCloud\WAF\Resources;

/**
 * Alarms: Alert / notification configuration.
 */
class Alarm extends BaseResource
{
    /** GET /commercial/alarm - Get alarm config */
    public function get(array $params = []): ?array
    {
        return $this->http->get('/commercial/alarm', $params);
    }

    /** PUT /commercial/alarm - Update alarm config */
    public function update(?array $data = null): ?array
    {
        return $this->http->put('/commercial/alarm', $data);
    }

    /** POST /commercial/alarm/test - Send a test alarm */
    public function test(?array $data = null): ?array
    {
        return $this->http->post('/commercial/alarm/test', $data);
    }

}