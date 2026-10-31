<?php

namespace HiTechCloud\WAF\Resources;

/**
 * SSL Certificates: TLS/SSL certificate management.
 */
class Ssl extends BaseResource
{
    /** GET /open/cert - List certificates */
    public function list(array $params = []): ?array
    {
        return $this->http->get('/open/cert', $params);
    }

    /** POST /open/cert - Upload / create a certificate */
    public function create(?array $data = null): ?array
    {
        return $this->http->post('/open/cert', $data);
    }

    /** GET /open/cert/{id} - Get a certificate */
    public function get(string $id, array $params = []): ?array
    {
        return $this->http->get('/open/cert/' . $id, $params);
    }

    /** DELETE /open/cert/{id} - Delete a certificate */
    public function delete(string $id, ?array $data = null): ?array
    {
        return $this->http->delete('/open/cert/' . $id, $data);
    }

}