package waf

// SystemResource handles System endpoints.
type SystemResource struct {
	client *HTTPClient
}

// Get - Get system settings (GET /open/system)
func (r *SystemResource) Get(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/open/system", params)
}

// Update - Update system settings (PUT /open/system)
func (r *SystemResource) Update(data interface{}) (map[string]interface{}, error) {
	return r.client.Put("/open/system", data)
}

// GetArch - Get the cluster architecture (GET /open/system/arch)
func (r *SystemResource) GetArch(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/open/system/arch", params)
}

// GetCustomerService - Get customer-service config (GET /open/system/customer_service)
func (r *SystemResource) GetCustomerService(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/open/system/customer_service", params)
}

// SetCustomerService - Set customer-service config (PUT /open/system/customer_service)
func (r *SystemResource) SetCustomerService(data interface{}) (map[string]interface{}, error) {
	return r.client.Put("/open/system/customer_service", data)
}

// GetEdition - Get the product edition (GET /open/system/edition)
func (r *SystemResource) GetEdition(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/open/system/edition", params)
}

// GetKey - Get the system public key (GET /open/system/key)
func (r *SystemResource) GetKey(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/open/system/key", params)
}

// GetLoginMethod - Get the system login method (GET /open/system/login_method)
func (r *SystemResource) GetLoginMethod(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/open/system/login_method", params)
}

// GetLoginType - Get the system login type (GET /open/system/login_type)
func (r *SystemResource) GetLoginType(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/open/system/login_type", params)
}

// GetNetworkProxy - Get the network proxy (GET /open/system/network_proxy)
func (r *SystemResource) GetNetworkProxy(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/open/system/network_proxy", params)
}

// SetNetworkProxy - Set the network proxy (POST /open/system/network_proxy)
func (r *SystemResource) SetNetworkProxy(data interface{}) (map[string]interface{}, error) {
	return r.client.Post("/open/system/network_proxy", data)
}

// GetProtocol - Get supported protocols (GET /open/system/protocol)
func (r *SystemResource) GetProtocol(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/open/system/protocol", params)
}

// SetProtocol - Configure supported protocols (POST /open/system/protocol)
func (r *SystemResource) SetProtocol(data interface{}) (map[string]interface{}, error) {
	return r.client.Post("/open/system/protocol", data)
}

// DashboardToken - Issue a dashboard token (POST /commercial/dashboard/token)
func (r *SystemResource) DashboardToken(data interface{}) (map[string]interface{}, error) {
	return r.client.Post("/commercial/dashboard/token", data)
}

// SafepointConsult - Consult the safepoint advisor (POST /open/safepoint/consult)
func (r *SystemResource) SafepointConsult(data interface{}) (map[string]interface{}, error) {
	return r.client.Post("/open/safepoint/consult", data)
}
