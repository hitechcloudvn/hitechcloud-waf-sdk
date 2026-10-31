package waf

import "strings"

// AntiTamperResource handles Anti-Tamper endpoints.
type AntiTamperResource struct {
	client *HTTPClient
}

// Get - Get an anti-tamper config (GET /business/anti_tamper/{id})
func (r *AntiTamperResource) Get(id string, params map[string]string) (map[string]interface{}, error) {
	return r.client.Get(strings.NewReplacer("{id}", id).Replace("/business/anti_tamper/{id}"), params)
}

// Update - Update an anti-tamper config (PUT /business/anti_tamper/{id})
func (r *AntiTamperResource) Update(id string, data interface{}) (map[string]interface{}, error) {
	return r.client.Put(strings.NewReplacer("{id}", id).Replace("/business/anti_tamper/{id}"), data)
}

// Delete - Delete an anti-tamper config (DELETE /business/anti_tamper/{id})
func (r *AntiTamperResource) Delete(id string, data interface{}) (map[string]interface{}, error) {
	return r.client.Delete(strings.NewReplacer("{id}", id).Replace("/business/anti_tamper/{id}"), data)
}

// GetPage - Get protected pages for a config (GET /business/anti_tamper/{id}/page)
func (r *AntiTamperResource) GetPage(id string, params map[string]string) (map[string]interface{}, error) {
	return r.client.Get(strings.NewReplacer("{id}", id).Replace("/business/anti_tamper/{id}/page"), params)
}

// ListForSite - List anti-tamper configs for a site (GET /business/site/{site_id}/anti_tamper)
func (r *AntiTamperResource) ListForSite(siteId string, params map[string]string) (map[string]interface{}, error) {
	return r.client.Get(strings.NewReplacer("{site_id}", siteId).Replace("/business/site/{site_id}/anti_tamper"), params)
}

// CreateForSite - Create an anti-tamper config for a site (POST /business/site/{site_id}/anti_tamper)
func (r *AntiTamperResource) CreateForSite(siteId string, data interface{}) (map[string]interface{}, error) {
	return r.client.Post(strings.NewReplacer("{site_id}", siteId).Replace("/business/site/{site_id}/anti_tamper"), data)
}

// UpdateForSite - Update anti-tamper for a site (PUT /business/site/{site_id}/anti_tamper)
func (r *AntiTamperResource) UpdateForSite(siteId string, data interface{}) (map[string]interface{}, error) {
	return r.client.Put(strings.NewReplacer("{site_id}", siteId).Replace("/business/site/{site_id}/anti_tamper"), data)
}

// GetResource - Get a protected resource (GET /business/site/{site_id}/anti_tamper/resource/{resource_id})
func (r *AntiTamperResource) GetResource(siteId string, resourceId string, params map[string]string) (map[string]interface{}, error) {
	return r.client.Get(strings.NewReplacer("{site_id}", siteId, "{resource_id}", resourceId).Replace("/business/site/{site_id}/anti_tamper/resource/{resource_id}"), params)
}
