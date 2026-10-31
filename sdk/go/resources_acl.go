package waf

import "strings"

// AclResource handles ACL & Rate-Limit endpoints.
type AclResource struct {
	client *HTTPClient
}

// Relieve - Relieve (unblock) an ACL-blocked client (PUT /acl/relieve)
func (r *AclResource) Relieve(data interface{}) (map[string]interface{}, error) {
	return r.client.Put("/acl/relieve", data)
}

// Records - ACL hit records (GET /open/records/acl)
func (r *AclResource) Records(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/open/records/acl", params)
}

// List - List ACL / rate-limit rules for a site (GET /open/site/{id}/acl)
func (r *AclResource) List(id string, params map[string]string) (map[string]interface{}, error) {
	return r.client.Get(strings.NewReplacer("{id}", id).Replace("/open/site/{id}/acl"), params)
}

// Create - Create an ACL rule (POST /open/site/{id}/acl)
func (r *AclResource) Create(id string, data interface{}) (map[string]interface{}, error) {
	return r.client.Post(strings.NewReplacer("{id}", id).Replace("/open/site/{id}/acl"), data)
}

// Update - Update ACL rules for a site (PUT /open/site/{id}/acl)
func (r *AclResource) Update(id string, data interface{}) (map[string]interface{}, error) {
	return r.client.Put(strings.NewReplacer("{id}", id).Replace("/open/site/{id}/acl"), data)
}

// UpdateRule - Update a single ACL rule (PUT /open/site/{id}/acl/{rule_id})
func (r *AclResource) UpdateRule(id string, ruleId string, data interface{}) (map[string]interface{}, error) {
	return r.client.Put(strings.NewReplacer("{id}", id, "{rule_id}", ruleId).Replace("/open/site/{id}/acl/{rule_id}"), data)
}

// DeleteRule - Delete a single ACL rule (DELETE /open/site/{id}/acl/{rule_id})
func (r *AclResource) DeleteRule(id string, ruleId string, data interface{}) (map[string]interface{}, error) {
	return r.client.Delete(strings.NewReplacer("{id}", id, "{rule_id}", ruleId).Replace("/open/site/{id}/acl/{rule_id}"), data)
}
