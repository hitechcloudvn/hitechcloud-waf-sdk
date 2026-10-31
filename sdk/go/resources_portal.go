package waf

// PortalResource handles Portal & Branding endpoints.
type PortalResource struct {
	client *HTTPClient
}

// UpdateStyle - Update portal style (commercial) (PUT /commercial/portal/style)
func (r *PortalResource) UpdateStyle(data interface{}) (map[string]interface{}, error) {
	return r.client.Put("/commercial/portal/style", data)
}

// Get - Get portal config (GET /open/portal)
func (r *PortalResource) Get(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/open/portal", params)
}

// Update - Update portal config (PUT /open/portal)
func (r *PortalResource) Update(data interface{}) (map[string]interface{}, error) {
	return r.client.Put("/open/portal", data)
}

// GetProxyConfig - Get portal proxy config (GET /open/portal/proxy_config)
func (r *PortalResource) GetProxyConfig(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/open/portal/proxy_config", params)
}

// SetProxyConfig - Set portal proxy config (PUT /open/portal/proxy_config)
func (r *PortalResource) SetProxyConfig(data interface{}) (map[string]interface{}, error) {
	return r.client.Put("/open/portal/proxy_config", data)
}

// GetStyle - Get portal style (GET /open/portal/style)
func (r *PortalResource) GetStyle(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/open/portal/style", params)
}

// UpdateBlockPage - Update the block page (commercial) (PUT /commercial/block_page)
func (r *PortalResource) UpdateBlockPage(data interface{}) (map[string]interface{}, error) {
	return r.client.Put("/commercial/block_page", data)
}

// GetBlockPage - Get the active block page (GET /open/block_page)
func (r *PortalResource) GetBlockPage(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/open/block_page", params)
}

// ListBlockPages - List available block pages (GET /open/block_page_list)
func (r *PortalResource) ListBlockPages(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/open/block_page_list", params)
}

// GetFrontendStyle - Get frontend style (GET /business/frontend_style)
func (r *PortalResource) GetFrontendStyle(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/business/frontend_style", params)
}

// SetFrontendStyle - Set frontend style (PUT /business/frontend_style)
func (r *PortalResource) SetFrontendStyle(data interface{}) (map[string]interface{}, error) {
	return r.client.Put("/business/frontend_style", data)
}

// GetAccount - Get portal account(s) (GET /business/account)
func (r *PortalResource) GetAccount(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/business/account", params)
}

// CreateAccount - Create a portal account (POST /business/account)
func (r *PortalResource) CreateAccount(data interface{}) (map[string]interface{}, error) {
	return r.client.Post("/business/account", data)
}
