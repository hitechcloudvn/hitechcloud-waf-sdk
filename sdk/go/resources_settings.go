package waf

// SettingsResource handles Settings endpoints.
type SettingsResource struct {
	client *HTTPClient
}

// SetLoginMethod - Set the admin login method (PUT /business/system/login_method)
func (r *SettingsResource) SetLoginMethod(data interface{}) (map[string]interface{}, error) {
	return r.client.Put("/business/system/login_method", data)
}

// TestLoginMethod - Test the admin login method (POST /business/system/login_method/test)
func (r *SettingsResource) TestLoginMethod(data interface{}) (map[string]interface{}, error) {
	return r.client.Post("/business/system/login_method/test", data)
}

// GetSyslog - Get syslog config (GET /commercial/syslog)
func (r *SettingsResource) GetSyslog(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/commercial/syslog", params)
}

// SetSyslog - Set syslog config (POST /commercial/syslog)
func (r *SettingsResource) SetSyslog(data interface{}) (map[string]interface{}, error) {
	return r.client.Post("/commercial/syslog", data)
}

// TestSyslog - Test syslog delivery (POST /commercial/syslog/test)
func (r *SettingsResource) TestSyslog(data interface{}) (map[string]interface{}, error) {
	return r.client.Post("/commercial/syslog/test", data)
}

// GetDowngrade - Get downgrade config (GET /open/Commercial/downgrade)
func (r *SettingsResource) GetDowngrade(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/open/Commercial/downgrade", params)
}

// SetDowngrade - Set downgrade config (PUT /open/Commercial/downgrade)
func (r *SettingsResource) SetDowngrade(data interface{}) (map[string]interface{}, error) {
	return r.client.Put("/open/Commercial/downgrade", data)
}

// AuditLog - Get the admin audit log (GET /open/audit_log)
func (r *SettingsResource) AuditLog(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/open/audit_log", params)
}

// GetDetector - Get detector config (GET /open/detector)
func (r *SettingsResource) GetDetector(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/open/detector", params)
}

// SetDetector - Set detector config (POST /open/detector)
func (r *SettingsResource) SetDetector(data interface{}) (map[string]interface{}, error) {
	return r.client.Post("/open/detector", data)
}
