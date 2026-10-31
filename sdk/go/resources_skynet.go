package waf

// SkynetResource handles Skynet AI Rules endpoints.
type SkynetResource struct {
	client *HTTPClient
}

// GetRules - Get Skynet rules (GET /open/skynet/rule)
func (r *SkynetResource) GetRules(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/open/skynet/rule", params)
}

// UpdateRules - Update Skynet rules (PUT /open/skynet/rule)
func (r *SkynetResource) UpdateRules(data interface{}) (map[string]interface{}, error) {
	return r.client.Put("/open/skynet/rule", data)
}

// GetRuleSwitch - Get Skynet rule switch state (GET /open/skynet/rule/switch)
func (r *SkynetResource) GetRuleSwitch(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/open/skynet/rule/switch", params)
}

// SetRuleSwitch - Toggle Skynet rules (PUT /open/skynet/rule/switch)
func (r *SkynetResource) SetRuleSwitch(data interface{}) (map[string]interface{}, error) {
	return r.client.Put("/open/skynet/rule/switch", data)
}
