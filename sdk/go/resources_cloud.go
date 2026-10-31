package waf

// CloudResource handles Cloud Policies endpoints.
type CloudResource struct {
	client *HTTPClient
}

// ListPolicies - List available cloud policies (GET /open/cloud/policies)
func (r *CloudResource) ListPolicies(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/open/cloud/policies", params)
}

// Subscribe - Subscribe to a cloud policy (POST /open/cloud/policies/subscribe)
func (r *CloudResource) Subscribe(data interface{}) (map[string]interface{}, error) {
	return r.client.Post("/open/cloud/policies/subscribe", data)
}
