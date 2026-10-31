<?php

namespace HiTechCloud\WAF\Resources;

/**
 * JA4 Fingerprint: JA4/JA4+ TLS client-fingerprint rules.
 */
class Ja4 extends BaseResource
{
    /** GET /open/ja4 - Get JA4 fingerprint rules */
    public function get(array $params = []): ?array
    {
        return $this->http->get('/open/ja4', $params);
    }

    /** PUT /open/ja4 - Update JA4 fingerprint rules */
    public function update(?array $data = null): ?array
    {
        return $this->http->put('/open/ja4', $data);
    }

}