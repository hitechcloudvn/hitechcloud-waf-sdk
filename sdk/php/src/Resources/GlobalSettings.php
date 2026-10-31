<?php

namespace HiTechCloud\WAF\Resources;

/**
 * Global Mode & Proxy: Global protection mode, outbound proxy and log cleanup.
 */
class GlobalSettings extends BaseResource
{
    /** GET /open/global/mode - Get the global protection mode */
    public function getMode(array $params = []): ?array
    {
        return $this->http->get('/open/global/mode', $params);
    }

    /** PUT /open/global/mode - Set the global protection mode */
    public function setMode(?array $data = null): ?array
    {
        return $this->http->put('/open/global/mode', $data);
    }

    /** GET /open/global/proxy - Get the global proxy config */
    public function getProxy(array $params = []): ?array
    {
        return $this->http->get('/open/global/proxy', $params);
    }

    /** PUT /open/global/proxy - Set the global proxy config */
    public function setProxy(?array $data = null): ?array
    {
        return $this->http->put('/open/global/proxy', $data);
    }

    /** GET /open/global/log_clean - Get the log-cleanup policy */
    public function getLogClean(array $params = []): ?array
    {
        return $this->http->get('/open/global/log_clean', $params);
    }

    /** POST /open/global/log_clean - Set the log-cleanup policy */
    public function setLogClean(?array $data = null): ?array
    {
        return $this->http->post('/open/global/log_clean', $data);
    }

}