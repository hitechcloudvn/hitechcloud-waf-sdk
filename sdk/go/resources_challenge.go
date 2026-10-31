package waf

import "strings"

// ChallengeResource handles Challenge & Semantics endpoints.
type ChallengeResource struct {
	client *HTTPClient
}

// GetConfig - Get challenge config (GET /open/challenge/config)
func (r *ChallengeResource) GetConfig(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/open/challenge/config", params)
}

// SetConfig - Set challenge config (POST /open/challenge/config)
func (r *ChallengeResource) SetConfig(data interface{}) (map[string]interface{}, error) {
	return r.client.Post("/open/challenge/config", data)
}

// Records - Challenge event records (GET /open/records/challenge)
func (r *ChallengeResource) Records(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/open/records/challenge", params)
}

// GetSemantics - Get semantic-analysis config for a site (GET /open/site/{id}/semantics)
func (r *ChallengeResource) GetSemantics(id string, params map[string]string) (map[string]interface{}, error) {
	return r.client.Get(strings.NewReplacer("{id}", id).Replace("/open/site/{id}/semantics"), params)
}

// SetSemantics - Set semantic-analysis config for a site (PUT /open/site/{id}/semantics)
func (r *ChallengeResource) SetSemantics(id string, data interface{}) (map[string]interface{}, error) {
	return r.client.Put(strings.NewReplacer("{id}", id).Replace("/open/site/{id}/semantics"), data)
}

// GetChaos - Get chaos/dynamic-protection config for a site (GET /open/site/{id}/chaos)
func (r *ChallengeResource) GetChaos(id string, params map[string]string) (map[string]interface{}, error) {
	return r.client.Get(strings.NewReplacer("{id}", id).Replace("/open/site/{id}/chaos"), params)
}

// SetChaos - Set chaos/dynamic-protection config for a site (POST /open/site/{id}/chaos)
func (r *ChallengeResource) SetChaos(id string, data interface{}) (map[string]interface{}, error) {
	return r.client.Post(strings.NewReplacer("{id}", id).Replace("/open/site/{id}/chaos"), data)
}
