<?php

namespace HiTechCloud\WAF\Resources;

/**
 * Skynet AI Rules: AI-driven Skynet detection rules and switches.
 */
class Skynet extends BaseResource
{
    /** GET /open/skynet/rule - Get Skynet rules */
    public function getRules(array $params = []): ?array
    {
        return $this->http->get('/open/skynet/rule', $params);
    }

    /** PUT /open/skynet/rule - Update Skynet rules */
    public function updateRules(?array $data = null): ?array
    {
        return $this->http->put('/open/skynet/rule', $data);
    }

    /** GET /open/skynet/rule/switch - Get Skynet rule switch state */
    public function getRuleSwitch(array $params = []): ?array
    {
        return $this->http->get('/open/skynet/rule/switch', $params);
    }

    /** PUT /open/skynet/rule/switch - Toggle Skynet rules */
    public function setRuleSwitch(?array $data = null): ?array
    {
        return $this->http->put('/open/skynet/rule/switch', $data);
    }

}