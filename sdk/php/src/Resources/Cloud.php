<?php

namespace HiTechCloud\WAF\Resources;

/**
 * Cloud Policies: Shared cloud threat-intelligence policies.
 */
class Cloud extends BaseResource
{
    /** GET /open/cloud/policies - List available cloud policies */
    public function listPolicies(array $params = []): ?array
    {
        return $this->http->get('/open/cloud/policies', $params);
    }

    /** POST /open/cloud/policies/subscribe - Subscribe to a cloud policy */
    public function subscribe(?array $data = null): ?array
    {
        return $this->http->post('/open/cloud/policies/subscribe', $data);
    }

}