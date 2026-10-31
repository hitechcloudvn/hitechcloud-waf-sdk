package waf

import "strings"

// SitesResource handles Sites (Web Services) endpoints.
type SitesResource struct {
	client *HTTPClient
}

// List - List protected sites (GET /open/site)
func (r *SitesResource) List(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/open/site", params)
}

// Create - Create a protected site (POST /open/site)
func (r *SitesResource) Create(data interface{}) (map[string]interface{}, error) {
	return r.client.Post("/open/site", data)
}

// Update - Update a site (PUT /open/site)
func (r *SitesResource) Update(data interface{}) (map[string]interface{}, error) {
	return r.client.Put("/open/site", data)
}

// Delete - Delete a site (DELETE /open/site)
func (r *SitesResource) Delete(data interface{}) (map[string]interface{}, error) {
	return r.client.Delete("/open/site", data)
}

// SetChallenge - Toggle challenge mode for sites (PUT /open/site/challenge)
func (r *SitesResource) SetChallenge(data interface{}) (map[string]interface{}, error) {
	return r.client.Put("/open/site/challenge", data)
}

// SetDefense - Toggle defense mode for sites (PUT /open/site/defense)
func (r *SitesResource) SetDefense(data interface{}) (map[string]interface{}, error) {
	return r.client.Put("/open/site/defense", data)
}

// ListGroups - List site groups (GET /open/site/group)
func (r *SitesResource) ListGroups(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/open/site/group", params)
}

// CreateGroup - Create a site group (POST /open/site/group)
func (r *SitesResource) CreateGroup(data interface{}) (map[string]interface{}, error) {
	return r.client.Post("/open/site/group", data)
}

// GetGroupSwitch - Get site-group switch state (GET /open/site/group/switch)
func (r *SitesResource) GetGroupSwitch(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/open/site/group/switch", params)
}

// UpdateGroup - Update a site group (PUT /open/site/group/{id})
func (r *SitesResource) UpdateGroup(id string, data interface{}) (map[string]interface{}, error) {
	return r.client.Put(strings.NewReplacer("{id}", id).Replace("/open/site/group/{id}"), data)
}

// DeleteGroup - Delete a site group (DELETE /open/site/group/{id})
func (r *SitesResource) DeleteGroup(id string, data interface{}) (map[string]interface{}, error) {
	return r.client.Delete(strings.NewReplacer("{id}", id).Replace("/open/site/group/{id}"), data)
}

// SortGroup - Sort a site group (PUT /open/site/group/{id}/sort)
func (r *SitesResource) SortGroup(id string, data interface{}) (map[string]interface{}, error) {
	return r.client.Put(strings.NewReplacer("{id}", id).Replace("/open/site/group/{id}/sort"), data)
}

// SwitchGroup - Switch the active site group (commercial) (PUT /commercial/site/group/switch)
func (r *SitesResource) SwitchGroup(data interface{}) (map[string]interface{}, error) {
	return r.client.Put("/commercial/site/group/switch", data)
}

// Healthcheck - Run an origin health check (POST /open/site/healthcheck)
func (r *SitesResource) Healthcheck(data interface{}) (map[string]interface{}, error) {
	return r.client.Post("/open/site/healthcheck", data)
}

// SetMode - Set the protection mode for a site (PUT /open/site/mode)
func (r *SitesResource) SetMode(data interface{}) (map[string]interface{}, error) {
	return r.client.Put("/open/site/mode", data)
}

// Get - Get a site (GET /open/site/{id})
func (r *SitesResource) Get(id string, params map[string]string) (map[string]interface{}, error) {
	return r.client.Get(strings.NewReplacer("{id}", id).Replace("/open/site/{id}"), params)
}

// UpdateBasicInfo - Update site basic info (PUT /open/site/{id}/basic_info)
func (r *SitesResource) UpdateBasicInfo(id string, data interface{}) (map[string]interface{}, error) {
	return r.client.Put(strings.NewReplacer("{id}", id).Replace("/open/site/{id}/basic_info"), data)
}

// GetExcludes - List exclusion rules (GET /open/site/{id}/excludes)
func (r *SitesResource) GetExcludes(id string, params map[string]string) (map[string]interface{}, error) {
	return r.client.Get(strings.NewReplacer("{id}", id).Replace("/open/site/{id}/excludes"), params)
}

// AddExclude - Add an exclusion rule (POST /open/site/{id}/excludes)
func (r *SitesResource) AddExclude(id string, data interface{}) (map[string]interface{}, error) {
	return r.client.Post(strings.NewReplacer("{id}", id).Replace("/open/site/{id}/excludes"), data)
}

// SetGroup - Move a site into a group (PUT /open/site/{id}/group)
func (r *SitesResource) SetGroup(id string, data interface{}) (map[string]interface{}, error) {
	return r.client.Put(strings.NewReplacer("{id}", id).Replace("/open/site/{id}/group"), data)
}

// GetLog - Get site logs (GET /open/site/{id}/log)
func (r *SitesResource) GetLog(id string, params map[string]string) (map[string]interface{}, error) {
	return r.client.Get(strings.NewReplacer("{id}", id).Replace("/open/site/{id}/log"), params)
}

// GetLogDetail - Get a log detail (GET /open/site/{id}/log/detail)
func (r *SitesResource) GetLogDetail(id string, params map[string]string) (map[string]interface{}, error) {
	return r.client.Get(strings.NewReplacer("{id}", id).Replace("/open/site/{id}/log/detail"), params)
}

// DownloadLog - Download site logs (GET /open/site/{id}/log/download)
func (r *SitesResource) DownloadLog(id string, params map[string]string) (map[string]interface{}, error) {
	return r.client.Get(strings.NewReplacer("{id}", id).Replace("/open/site/{id}/log/download"), params)
}

// GetLogLimit - Get log limit config (GET /open/site/{id}/log/limit)
func (r *SitesResource) GetLogLimit(id string, params map[string]string) (map[string]interface{}, error) {
	return r.client.Get(strings.NewReplacer("{id}", id).Replace("/open/site/{id}/log/limit"), params)
}

// SetLogLimit - Set log limit config (POST /open/site/{id}/log/limit)
func (r *SitesResource) SetLogLimit(id string, data interface{}) (map[string]interface{}, error) {
	return r.client.Post(strings.NewReplacer("{id}", id).Replace("/open/site/{id}/log/limit"), data)
}

// GetNginxConfig - Get generated nginx config (GET /open/site/{id}/nginx_config)
func (r *SitesResource) GetNginxConfig(id string, params map[string]string) (map[string]interface{}, error) {
	return r.client.Get(strings.NewReplacer("{id}", id).Replace("/open/site/{id}/nginx_config"), params)
}

// UpdateNginxConfig - Update custom nginx config (PUT /open/site/{id}/nginx_config)
func (r *SitesResource) UpdateNginxConfig(id string, data interface{}) (map[string]interface{}, error) {
	return r.client.Put(strings.NewReplacer("{id}", id).Replace("/open/site/{id}/nginx_config"), data)
}

// GetProxy - Get reverse-proxy / origin config (GET /open/site/{id}/proxy)
func (r *SitesResource) GetProxy(id string, params map[string]string) (map[string]interface{}, error) {
	return r.client.Get(strings.NewReplacer("{id}", id).Replace("/open/site/{id}/proxy"), params)
}

// UpdateProxy - Update reverse-proxy / origin config (PUT /open/site/{id}/proxy)
func (r *SitesResource) UpdateProxy(id string, data interface{}) (map[string]interface{}, error) {
	return r.client.Put(strings.NewReplacer("{id}", id).Replace("/open/site/{id}/proxy"), data)
}

// GetResources - List static resources (GET /open/site/{id}/resources)
func (r *SitesResource) GetResources(id string, params map[string]string) (map[string]interface{}, error) {
	return r.client.Get(strings.NewReplacer("{id}", id).Replace("/open/site/{id}/resources"), params)
}

// DeleteResources - Delete static resources (DELETE /open/site/{id}/resources)
func (r *SitesResource) DeleteResources(id string, data interface{}) (map[string]interface{}, error) {
	return r.client.Delete(strings.NewReplacer("{id}", id).Replace("/open/site/{id}/resources"), data)
}

// Sort - Sort a site (PUT /open/site/{id}/sort)
func (r *SitesResource) Sort(id string, data interface{}) (map[string]interface{}, error) {
	return r.client.Put(strings.NewReplacer("{id}", id).Replace("/open/site/{id}/sort"), data)
}

// GetStatic - List static-site rules (GET /open/site/{id}/static)
func (r *SitesResource) GetStatic(id string, params map[string]string) (map[string]interface{}, error) {
	return r.client.Get(strings.NewReplacer("{id}", id).Replace("/open/site/{id}/static"), params)
}

// AddStatic - Add a static-site rule (POST /open/site/{id}/static)
func (r *SitesResource) AddStatic(id string, data interface{}) (map[string]interface{}, error) {
	return r.client.Post(strings.NewReplacer("{id}", id).Replace("/open/site/{id}/static"), data)
}

// UpdateStatic - Update a static-site rule (PUT /open/site/{id}/static)
func (r *SitesResource) UpdateStatic(id string, data interface{}) (map[string]interface{}, error) {
	return r.client.Put(strings.NewReplacer("{id}", id).Replace("/open/site/{id}/static"), data)
}

// DeleteStatic - Delete a static-site rule (DELETE /open/site/{id}/static)
func (r *SitesResource) DeleteStatic(id string, data interface{}) (map[string]interface{}, error) {
	return r.client.Delete(strings.NewReplacer("{id}", id).Replace("/open/site/{id}/static"), data)
}

// ListForwardingRules - List forwarding rules (GET /open/site/{id}/forwarding_rules)
func (r *SitesResource) ListForwardingRules(id string, params map[string]string) (map[string]interface{}, error) {
	return r.client.Get(strings.NewReplacer("{id}", id).Replace("/open/site/{id}/forwarding_rules"), params)
}

// CreateForwardingRule - Create a forwarding rule (POST /open/site/{id}/forwarding_rules)
func (r *SitesResource) CreateForwardingRule(id string, data interface{}) (map[string]interface{}, error) {
	return r.client.Post(strings.NewReplacer("{id}", id).Replace("/open/site/{id}/forwarding_rules"), data)
}

// UpdateForwardingRule - Update a forwarding rule (PUT /open/site/{id}/forwarding_rules/{rule_id})
func (r *SitesResource) UpdateForwardingRule(id string, ruleId string, data interface{}) (map[string]interface{}, error) {
	return r.client.Put(strings.NewReplacer("{id}", id, "{rule_id}", ruleId).Replace("/open/site/{id}/forwarding_rules/{rule_id}"), data)
}

// DeleteForwardingRule - Delete a forwarding rule (DELETE /open/site/{id}/forwarding_rules/{rule_id})
func (r *SitesResource) DeleteForwardingRule(id string, ruleId string, data interface{}) (map[string]interface{}, error) {
	return r.client.Delete(strings.NewReplacer("{id}", id, "{rule_id}", ruleId).Replace("/open/site/{id}/forwarding_rules/{rule_id}"), data)
}

// SwitchForwardingRule - Enable / disable a forwarding rule (PUT /open/site/{id}/forwarding_rules/{rule_id}/switch)
func (r *SitesResource) SwitchForwardingRule(id string, ruleId string, data interface{}) (map[string]interface{}, error) {
	return r.client.Put(strings.NewReplacer("{id}", id, "{rule_id}", ruleId).Replace("/open/site/{id}/forwarding_rules/{rule_id}/switch"), data)
}
