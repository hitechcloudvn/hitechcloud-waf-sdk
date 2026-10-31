<?php

namespace HiTechCloud\WAF\Resources;

/**
 * License & Authorization: Authorization records and license status / lifecycle.
 */
class License extends BaseResource
{
    /** GET /open/system/authorize - Get authorization records */
    public function getAuthorize(array $params = []): ?array
    {
        return $this->http->get('/open/system/authorize', $params);
    }

    /** POST /open/system/authorize - Create an authorization */
    public function createAuthorize(?array $data = null): ?array
    {
        return $this->http->post('/open/system/authorize', $data);
    }

    /** PUT /open/system/authorize - Update an authorization */
    public function updateAuthorize(?array $data = null): ?array
    {
        return $this->http->put('/open/system/authorize', $data);
    }

    /** DELETE /open/system/authorize - Delete an authorization */
    public function deleteAuthorize(?array $data = null): ?array
    {
        return $this->http->delete('/open/system/authorize', $data);
    }

    /** POST /open/system/license/expired/confirm-downgrade - Confirm downgrade after license expiry */
    public function confirmDowngrade(?array $data = null): ?array
    {
        return $this->http->post('/open/system/license/expired/confirm-downgrade', $data);
    }

    /** POST /open/system/license/expiry/acknowledge - Acknowledge an upcoming license expiry */
    public function acknowledgeExpiry(?array $data = null): ?array
    {
        return $this->http->post('/open/system/license/expiry/acknowledge', $data);
    }

    /** GET /open/system/license/status - Get the license status */
    public function status(array $params = []): ?array
    {
        return $this->http->get('/open/system/license/status', $params);
    }

}