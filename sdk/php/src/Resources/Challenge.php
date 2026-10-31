<?php

namespace HiTechCloud\WAF\Resources;

/**
 * Challenge & Semantics: Human-verification challenge, semantic analysis and chaos protection.
 */
class Challenge extends BaseResource
{
    /** GET /open/challenge/config - Get challenge config */
    public function getConfig(array $params = []): ?array
    {
        return $this->http->get('/open/challenge/config', $params);
    }

    /** POST /open/challenge/config - Set challenge config */
    public function setConfig(?array $data = null): ?array
    {
        return $this->http->post('/open/challenge/config', $data);
    }

    /** GET /open/records/challenge - Challenge event records */
    public function records(array $params = []): ?array
    {
        return $this->http->get('/open/records/challenge', $params);
    }

    /** GET /open/site/{id}/semantics - Get semantic-analysis config for a site */
    public function getSemantics(string $id, array $params = []): ?array
    {
        return $this->http->get('/open/site/' . $id . '/semantics', $params);
    }

    /** PUT /open/site/{id}/semantics - Set semantic-analysis config for a site */
    public function setSemantics(string $id, ?array $data = null): ?array
    {
        return $this->http->put('/open/site/' . $id . '/semantics', $data);
    }

    /** GET /open/site/{id}/chaos - Get chaos/dynamic-protection config for a site */
    public function getChaos(string $id, array $params = []): ?array
    {
        return $this->http->get('/open/site/' . $id . '/chaos', $params);
    }

    /** POST /open/site/{id}/chaos - Set chaos/dynamic-protection config for a site */
    public function setChaos(string $id, ?array $data = null): ?array
    {
        return $this->http->post('/open/site/' . $id . '/chaos', $data);
    }

}