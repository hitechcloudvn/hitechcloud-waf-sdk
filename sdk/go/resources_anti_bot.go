package waf

// AntiBotResource handles Anti-Bot & Intelligence endpoints.
type AntiBotResource struct {
	client *HTTPClient
}

// GetIntelligence - Get threat-intelligence config (GET /open/intelligence)
func (r *AntiBotResource) GetIntelligence(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/open/intelligence", params)
}

// SetIntelligence - Set threat-intelligence config (POST /open/intelligence)
func (r *AntiBotResource) SetIntelligence(data interface{}) (map[string]interface{}, error) {
	return r.client.Post("/open/intelligence", data)
}

// UploadIpLib - Upload a custom IP library (POST /open/intelligence/ip_lib)
func (r *AntiBotResource) UploadIpLib(data interface{}) (map[string]interface{}, error) {
	return r.client.Post("/open/intelligence/ip_lib", data)
}

// ListIpGroups - List IP groups (GET /open/ipgroup)
func (r *AntiBotResource) ListIpGroups(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/open/ipgroup", params)
}

// CreateIpGroup - Create an IP group (POST /open/ipgroup)
func (r *AntiBotResource) CreateIpGroup(data interface{}) (map[string]interface{}, error) {
	return r.client.Post("/open/ipgroup", data)
}

// UpdateIpGroup - Update an IP group (PUT /open/ipgroup)
func (r *AntiBotResource) UpdateIpGroup(data interface{}) (map[string]interface{}, error) {
	return r.client.Put("/open/ipgroup", data)
}

// DeleteIpGroup - Delete an IP group (DELETE /open/ipgroup)
func (r *AntiBotResource) DeleteIpGroup(data interface{}) (map[string]interface{}, error) {
	return r.client.Delete("/open/ipgroup", data)
}

// AppendIpGroup - Append entries to an IP group (POST /open/ipgroup/append)
func (r *AntiBotResource) AppendIpGroup(data interface{}) (map[string]interface{}, error) {
	return r.client.Post("/open/ipgroup/append", data)
}

// GetCrawlers - Get known-crawler groups (GET /open/ipgroup/crawler)
func (r *AntiBotResource) GetCrawlers(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/open/ipgroup/crawler", params)
}

// SetCrawlers - Configure known-crawler groups (POST /open/ipgroup/crawler)
func (r *AntiBotResource) SetCrawlers(data interface{}) (map[string]interface{}, error) {
	return r.client.Post("/open/ipgroup/crawler", data)
}

// GetIpGroupDetail - Get IP-group detail (GET /open/ipgroup/detail)
func (r *AntiBotResource) GetIpGroupDetail(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/open/ipgroup/detail", params)
}

// GetIpGroupLink - Get IP-group link bindings (GET /open/ipgroup/link)
func (r *AntiBotResource) GetIpGroupLink(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/open/ipgroup/link", params)
}

// SetIpGroupLink - Set IP-group link bindings (POST /open/ipgroup/link)
func (r *AntiBotResource) SetIpGroupLink(data interface{}) (map[string]interface{}, error) {
	return r.client.Post("/open/ipgroup/link", data)
}

// GetTrustServers - Get trusted-server IP groups (GET /open/ipgroup/tserver)
func (r *AntiBotResource) GetTrustServers(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/open/ipgroup/tserver", params)
}

// SetTrustServers - Set trusted-server IP groups (POST /open/ipgroup/tserver)
func (r *AntiBotResource) SetTrustServers(data interface{}) (map[string]interface{}, error) {
	return r.client.Post("/open/ipgroup/tserver", data)
}

// GetShareBehaviour - Get behaviour-sharing config (GET /open/share_behaviour)
func (r *AntiBotResource) GetShareBehaviour(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/open/share_behaviour", params)
}

// SetShareBehaviour - Set behaviour-sharing config (POST /open/share_behaviour)
func (r *AntiBotResource) SetShareBehaviour(data interface{}) (map[string]interface{}, error) {
	return r.client.Post("/open/share_behaviour", data)
}

// GetShareFingerprint - Get fingerprint-sharing config (GET /open/share_fingerprint)
func (r *AntiBotResource) GetShareFingerprint(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/open/share_fingerprint", params)
}

// SetShareFingerprint - Set fingerprint-sharing config (POST /open/share_fingerprint)
func (r *AntiBotResource) SetShareFingerprint(data interface{}) (map[string]interface{}, error) {
	return r.client.Post("/open/share_fingerprint", data)
}
