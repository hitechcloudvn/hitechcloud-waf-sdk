<?php

namespace HiTechCloud\WAF\Resources;

/**
 * Sites (Web Services): Protected sites, groups, proxy/origin config, logs and forwarding rules.
 */
class Sites extends BaseResource
{
    /** GET /open/site - List protected sites */
    public function list(array $params = []): ?array
    {
        return $this->http->get('/open/site', $params);
    }

    /** POST /open/site - Create a protected site */
    public function create(?array $data = null): ?array
    {
        return $this->http->post('/open/site', $data);
    }

    /** PUT /open/site - Update a site */
    public function update(?array $data = null): ?array
    {
        return $this->http->put('/open/site', $data);
    }

    /** DELETE /open/site - Delete a site */
    public function delete(?array $data = null): ?array
    {
        return $this->http->delete('/open/site', $data);
    }

    /** PUT /open/site/challenge - Toggle challenge mode for sites */
    public function setChallenge(?array $data = null): ?array
    {
        return $this->http->put('/open/site/challenge', $data);
    }

    /** PUT /open/site/defense - Toggle defense mode for sites */
    public function setDefense(?array $data = null): ?array
    {
        return $this->http->put('/open/site/defense', $data);
    }

    /** GET /open/site/group - List site groups */
    public function listGroups(array $params = []): ?array
    {
        return $this->http->get('/open/site/group', $params);
    }

    /** POST /open/site/group - Create a site group */
    public function createGroup(?array $data = null): ?array
    {
        return $this->http->post('/open/site/group', $data);
    }

    /** GET /open/site/group/switch - Get site-group switch state */
    public function getGroupSwitch(array $params = []): ?array
    {
        return $this->http->get('/open/site/group/switch', $params);
    }

    /** PUT /open/site/group/{id} - Update a site group */
    public function updateGroup(string $id, ?array $data = null): ?array
    {
        return $this->http->put('/open/site/group/' . $id, $data);
    }

    /** DELETE /open/site/group/{id} - Delete a site group */
    public function deleteGroup(string $id, ?array $data = null): ?array
    {
        return $this->http->delete('/open/site/group/' . $id, $data);
    }

    /** PUT /open/site/group/{id}/sort - Sort a site group */
    public function sortGroup(string $id, ?array $data = null): ?array
    {
        return $this->http->put('/open/site/group/' . $id . '/sort', $data);
    }

    /** PUT /commercial/site/group/switch - Switch the active site group (commercial) */
    public function switchGroup(?array $data = null): ?array
    {
        return $this->http->put('/commercial/site/group/switch', $data);
    }

    /** POST /open/site/healthcheck - Run an origin health check */
    public function healthcheck(?array $data = null): ?array
    {
        return $this->http->post('/open/site/healthcheck', $data);
    }

    /** PUT /open/site/mode - Set the protection mode for a site */
    public function setMode(?array $data = null): ?array
    {
        return $this->http->put('/open/site/mode', $data);
    }

    /** GET /open/site/{id} - Get a site */
    public function get(string $id, array $params = []): ?array
    {
        return $this->http->get('/open/site/' . $id, $params);
    }

    /** PUT /open/site/{id}/basic_info - Update site basic info */
    public function updateBasicInfo(string $id, ?array $data = null): ?array
    {
        return $this->http->put('/open/site/' . $id . '/basic_info', $data);
    }

    /** GET /open/site/{id}/excludes - List exclusion rules */
    public function getExcludes(string $id, array $params = []): ?array
    {
        return $this->http->get('/open/site/' . $id . '/excludes', $params);
    }

    /** POST /open/site/{id}/excludes - Add an exclusion rule */
    public function addExclude(string $id, ?array $data = null): ?array
    {
        return $this->http->post('/open/site/' . $id . '/excludes', $data);
    }

    /** PUT /open/site/{id}/group - Move a site into a group */
    public function setGroup(string $id, ?array $data = null): ?array
    {
        return $this->http->put('/open/site/' . $id . '/group', $data);
    }

    /** GET /open/site/{id}/log - Get site logs */
    public function getLog(string $id, array $params = []): ?array
    {
        return $this->http->get('/open/site/' . $id . '/log', $params);
    }

    /** GET /open/site/{id}/log/detail - Get a log detail */
    public function getLogDetail(string $id, array $params = []): ?array
    {
        return $this->http->get('/open/site/' . $id . '/log/detail', $params);
    }

    /** GET /open/site/{id}/log/download - Download site logs */
    public function downloadLog(string $id, array $params = []): ?array
    {
        return $this->http->get('/open/site/' . $id . '/log/download', $params);
    }

    /** GET /open/site/{id}/log/limit - Get log limit config */
    public function getLogLimit(string $id, array $params = []): ?array
    {
        return $this->http->get('/open/site/' . $id . '/log/limit', $params);
    }

    /** POST /open/site/{id}/log/limit - Set log limit config */
    public function setLogLimit(string $id, ?array $data = null): ?array
    {
        return $this->http->post('/open/site/' . $id . '/log/limit', $data);
    }

    /** GET /open/site/{id}/nginx_config - Get generated nginx config */
    public function getNginxConfig(string $id, array $params = []): ?array
    {
        return $this->http->get('/open/site/' . $id . '/nginx_config', $params);
    }

    /** PUT /open/site/{id}/nginx_config - Update custom nginx config */
    public function updateNginxConfig(string $id, ?array $data = null): ?array
    {
        return $this->http->put('/open/site/' . $id . '/nginx_config', $data);
    }

    /** GET /open/site/{id}/proxy - Get reverse-proxy / origin config */
    public function getProxy(string $id, array $params = []): ?array
    {
        return $this->http->get('/open/site/' . $id . '/proxy', $params);
    }

    /** PUT /open/site/{id}/proxy - Update reverse-proxy / origin config */
    public function updateProxy(string $id, ?array $data = null): ?array
    {
        return $this->http->put('/open/site/' . $id . '/proxy', $data);
    }

    /** GET /open/site/{id}/resources - List static resources */
    public function getResources(string $id, array $params = []): ?array
    {
        return $this->http->get('/open/site/' . $id . '/resources', $params);
    }

    /** DELETE /open/site/{id}/resources - Delete static resources */
    public function deleteResources(string $id, ?array $data = null): ?array
    {
        return $this->http->delete('/open/site/' . $id . '/resources', $data);
    }

    /** PUT /open/site/{id}/sort - Sort a site */
    public function sort(string $id, ?array $data = null): ?array
    {
        return $this->http->put('/open/site/' . $id . '/sort', $data);
    }

    /** GET /open/site/{id}/static - List static-site rules */
    public function getStatic(string $id, array $params = []): ?array
    {
        return $this->http->get('/open/site/' . $id . '/static', $params);
    }

    /** POST /open/site/{id}/static - Add a static-site rule */
    public function addStatic(string $id, ?array $data = null): ?array
    {
        return $this->http->post('/open/site/' . $id . '/static', $data);
    }

    /** PUT /open/site/{id}/static - Update a static-site rule */
    public function updateStatic(string $id, ?array $data = null): ?array
    {
        return $this->http->put('/open/site/' . $id . '/static', $data);
    }

    /** DELETE /open/site/{id}/static - Delete a static-site rule */
    public function deleteStatic(string $id, ?array $data = null): ?array
    {
        return $this->http->delete('/open/site/' . $id . '/static', $data);
    }

    /** GET /open/site/{id}/forwarding_rules - List forwarding rules */
    public function listForwardingRules(string $id, array $params = []): ?array
    {
        return $this->http->get('/open/site/' . $id . '/forwarding_rules', $params);
    }

    /** POST /open/site/{id}/forwarding_rules - Create a forwarding rule */
    public function createForwardingRule(string $id, ?array $data = null): ?array
    {
        return $this->http->post('/open/site/' . $id . '/forwarding_rules', $data);
    }

    /** PUT /open/site/{id}/forwarding_rules/{rule_id} - Update a forwarding rule */
    public function updateForwardingRule(string $id, string $ruleId, ?array $data = null): ?array
    {
        return $this->http->put('/open/site/' . $id . '/forwarding_rules/' . $ruleId, $data);
    }

    /** DELETE /open/site/{id}/forwarding_rules/{rule_id} - Delete a forwarding rule */
    public function deleteForwardingRule(string $id, string $ruleId, ?array $data = null): ?array
    {
        return $this->http->delete('/open/site/' . $id . '/forwarding_rules/' . $ruleId, $data);
    }

    /** PUT /open/site/{id}/forwarding_rules/{rule_id}/switch - Enable / disable a forwarding rule */
    public function switchForwardingRule(string $id, string $ruleId, ?array $data = null): ?array
    {
        return $this->http->put('/open/site/' . $id . '/forwarding_rules/' . $ruleId . '/switch', $data);
    }

}