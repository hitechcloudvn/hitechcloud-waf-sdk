<?php

namespace HiTechCloud\WAF\Resources;

/**
 * Waiting Room: Virtual waiting-room queue for traffic peaks.
 */
class WaitingRoom extends BaseResource
{
    /** GET /open/records/waiting - Waiting-room event records */
    public function records(array $params = []): ?array
    {
        return $this->http->get('/open/records/waiting', $params);
    }

    /** GET /open/site/{id}/waiting - Get waiting-room config for a site */
    public function get(string $id, array $params = []): ?array
    {
        return $this->http->get('/open/site/' . $id . '/waiting', $params);
    }

    /** POST /open/site/{id}/waiting - Set waiting-room config for a site */
    public function set(string $id, ?array $data = null): ?array
    {
        return $this->http->post('/open/site/' . $id . '/waiting', $data);
    }

}