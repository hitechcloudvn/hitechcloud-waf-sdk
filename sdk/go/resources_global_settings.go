package waf

// GlobalSettingsResource handles Global Mode & Proxy endpoints.
type GlobalSettingsResource struct {
	client *HTTPClient
}

// GetMode - Get the global protection mode (GET /open/global/mode)
func (r *GlobalSettingsResource) GetMode(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/open/global/mode", params)
}

// SetMode - Set the global protection mode (PUT /open/global/mode)
func (r *GlobalSettingsResource) SetMode(data interface{}) (map[string]interface{}, error) {
	return r.client.Put("/open/global/mode", data)
}

// GetProxy - Get the global proxy config (GET /open/global/proxy)
func (r *GlobalSettingsResource) GetProxy(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/open/global/proxy", params)
}

// SetProxy - Set the global proxy config (PUT /open/global/proxy)
func (r *GlobalSettingsResource) SetProxy(data interface{}) (map[string]interface{}, error) {
	return r.client.Put("/open/global/proxy", data)
}

// GetLogClean - Get the log-cleanup policy (GET /open/global/log_clean)
func (r *GlobalSettingsResource) GetLogClean(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/open/global/log_clean", params)
}

// SetLogClean - Set the log-cleanup policy (POST /open/global/log_clean)
func (r *GlobalSettingsResource) SetLogClean(data interface{}) (map[string]interface{}, error) {
	return r.client.Post("/open/global/log_clean", data)
}
