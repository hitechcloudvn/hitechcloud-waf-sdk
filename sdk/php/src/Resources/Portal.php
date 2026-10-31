<?php

namespace HiTechCloud\WAF\Resources;

/**
 * Portal & Branding: Self-service portal, block pages, branding and accounts.
 */
class Portal extends BaseResource
{
    /** PUT /commercial/portal/style - Update portal style (commercial) */
    public function updateStyle(?array $data = null): ?array
    {
        return $this->http->put('/commercial/portal/style', $data);
    }

    /** GET /open/portal - Get portal config */
    public function get(array $params = []): ?array
    {
        return $this->http->get('/open/portal', $params);
    }

    /** PUT /open/portal - Update portal config */
    public function update(?array $data = null): ?array
    {
        return $this->http->put('/open/portal', $data);
    }

    /** GET /open/portal/proxy_config - Get portal proxy config */
    public function getProxyConfig(array $params = []): ?array
    {
        return $this->http->get('/open/portal/proxy_config', $params);
    }

    /** PUT /open/portal/proxy_config - Set portal proxy config */
    public function setProxyConfig(?array $data = null): ?array
    {
        return $this->http->put('/open/portal/proxy_config', $data);
    }

    /** GET /open/portal/style - Get portal style */
    public function getStyle(array $params = []): ?array
    {
        return $this->http->get('/open/portal/style', $params);
    }

    /** PUT /commercial/block_page - Update the block page (commercial) */
    public function updateBlockPage(?array $data = null): ?array
    {
        return $this->http->put('/commercial/block_page', $data);
    }

    /** GET /open/block_page - Get the active block page */
    public function getBlockPage(array $params = []): ?array
    {
        return $this->http->get('/open/block_page', $params);
    }

    /** GET /open/block_page_list - List available block pages */
    public function listBlockPages(array $params = []): ?array
    {
        return $this->http->get('/open/block_page_list', $params);
    }

    /** GET /business/frontend_style - Get frontend style */
    public function getFrontendStyle(array $params = []): ?array
    {
        return $this->http->get('/business/frontend_style', $params);
    }

    /** PUT /business/frontend_style - Set frontend style */
    public function setFrontendStyle(?array $data = null): ?array
    {
        return $this->http->put('/business/frontend_style', $data);
    }

    /** GET /business/account - Get portal account(s) */
    public function getAccount(array $params = []): ?array
    {
        return $this->http->get('/business/account', $params);
    }

    /** POST /business/account - Create a portal account */
    public function createAccount(?array $data = null): ?array
    {
        return $this->http->post('/business/account', $data);
    }

}