package waf

// LicenseResource handles License & Authorization endpoints.
type LicenseResource struct {
	client *HTTPClient
}

// GetAuthorize - Get authorization records (GET /open/system/authorize)
func (r *LicenseResource) GetAuthorize(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/open/system/authorize", params)
}

// CreateAuthorize - Create an authorization (POST /open/system/authorize)
func (r *LicenseResource) CreateAuthorize(data interface{}) (map[string]interface{}, error) {
	return r.client.Post("/open/system/authorize", data)
}

// UpdateAuthorize - Update an authorization (PUT /open/system/authorize)
func (r *LicenseResource) UpdateAuthorize(data interface{}) (map[string]interface{}, error) {
	return r.client.Put("/open/system/authorize", data)
}

// DeleteAuthorize - Delete an authorization (DELETE /open/system/authorize)
func (r *LicenseResource) DeleteAuthorize(data interface{}) (map[string]interface{}, error) {
	return r.client.Delete("/open/system/authorize", data)
}

// ConfirmDowngrade - Confirm downgrade after license expiry (POST /open/system/license/expired/confirm-downgrade)
func (r *LicenseResource) ConfirmDowngrade(data interface{}) (map[string]interface{}, error) {
	return r.client.Post("/open/system/license/expired/confirm-downgrade", data)
}

// AcknowledgeExpiry - Acknowledge an upcoming license expiry (POST /open/system/license/expiry/acknowledge)
func (r *LicenseResource) AcknowledgeExpiry(data interface{}) (map[string]interface{}, error) {
	return r.client.Post("/open/system/license/expiry/acknowledge", data)
}

// Status - Get the license status (GET /open/system/license/status)
func (r *LicenseResource) Status(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/open/system/license/status", params)
}
