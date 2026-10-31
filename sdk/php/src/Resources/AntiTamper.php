<?php

namespace HiTechCloud\WAF\Resources;

/**
 * Anti-Tamper: Content anti-tamper protection per site and resource.
 */
class AntiTamper extends BaseResource
{
    /** GET /business/anti_tamper/{id} - Get an anti-tamper config */
    public function get(string $id, array $params = []): ?array
    {
        return $this->http->get('/business/anti_tamper/' . $id, $params);
    }

    /** PUT /business/anti_tamper/{id} - Update an anti-tamper config */
    public function update(string $id, ?array $data = null): ?array
    {
        return $this->http->put('/business/anti_tamper/' . $id, $data);
    }

    /** DELETE /business/anti_tamper/{id} - Delete an anti-tamper config */
    public function delete(string $id, ?array $data = null): ?array
    {
        return $this->http->delete('/business/anti_tamper/' . $id, $data);
    }

    /** GET /business/anti_tamper/{id}/page - Get protected pages for a config */
    public function getPage(string $id, array $params = []): ?array
    {
        return $this->http->get('/business/anti_tamper/' . $id . '/page', $params);
    }

    /** GET /business/site/{site_id}/anti_tamper - List anti-tamper configs for a site */
    public function listForSite(string $siteId, array $params = []): ?array
    {
        return $this->http->get('/business/site/' . $siteId . '/anti_tamper', $params);
    }

    /** POST /business/site/{site_id}/anti_tamper - Create an anti-tamper config for a site */
    public function createForSite(string $siteId, ?array $data = null): ?array
    {
        return $this->http->post('/business/site/' . $siteId . '/anti_tamper', $data);
    }

    /** PUT /business/site/{site_id}/anti_tamper - Update anti-tamper for a site */
    public function updateForSite(string $siteId, ?array $data = null): ?array
    {
        return $this->http->put('/business/site/' . $siteId . '/anti_tamper', $data);
    }

    /** GET /business/site/{site_id}/anti_tamper/resource/{resource_id} - Get a protected resource */
    public function getResource(string $siteId, string $resourceId, array $params = []): ?array
    {
        return $this->http->get('/business/site/' . $siteId . '/anti_tamper/resource/' . $resourceId, $params);
    }

}