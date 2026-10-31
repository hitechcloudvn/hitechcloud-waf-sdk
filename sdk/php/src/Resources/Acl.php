<?php

namespace HiTechCloud\WAF\Resources;

/**
 * ACL & Rate-Limit: Access-control / rate-limit rules and block relief.
 */
class Acl extends BaseResource
{
    /** PUT /acl/relieve - Relieve (unblock) an ACL-blocked client */
    public function relieve(?array $data = null): ?array
    {
        return $this->http->put('/acl/relieve', $data);
    }

    /** GET /open/records/acl - ACL hit records */
    public function records(array $params = []): ?array
    {
        return $this->http->get('/open/records/acl', $params);
    }

    /** GET /open/site/{id}/acl - List ACL / rate-limit rules for a site */
    public function list(string $id, array $params = []): ?array
    {
        return $this->http->get('/open/site/' . $id . '/acl', $params);
    }

    /** POST /open/site/{id}/acl - Create an ACL rule */
    public function create(string $id, ?array $data = null): ?array
    {
        return $this->http->post('/open/site/' . $id . '/acl', $data);
    }

    /** PUT /open/site/{id}/acl - Update ACL rules for a site */
    public function update(string $id, ?array $data = null): ?array
    {
        return $this->http->put('/open/site/' . $id . '/acl', $data);
    }

    /** PUT /open/site/{id}/acl/{rule_id} - Update a single ACL rule */
    public function updateRule(string $id, string $ruleId, ?array $data = null): ?array
    {
        return $this->http->put('/open/site/' . $id . '/acl/' . $ruleId, $data);
    }

    /** DELETE /open/site/{id}/acl/{rule_id} - Delete a single ACL rule */
    public function deleteRule(string $id, string $ruleId, ?array $data = null): ?array
    {
        return $this->http->delete('/open/site/' . $id . '/acl/' . $ruleId, $data);
    }

}