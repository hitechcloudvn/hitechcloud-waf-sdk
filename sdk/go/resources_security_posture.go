package waf

// SecurityPostureResource handles Security Posture endpoints.
type SecurityPostureResource struct {
	client *HTTPClient
}

// Realtime - Realtime security posture (GET /open/security_posture/realtime)
func (r *SecurityPostureResource) Realtime(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/open/security_posture/realtime", params)
}

// Site - Security posture for selected sites (POST /open/security_posture/site)
func (r *SecurityPostureResource) Site(data interface{}) (map[string]interface{}, error) {
	return r.client.Post("/open/security_posture/site", data)
}

// Statistics - Security-posture statistics (GET /open/security_posture/statistics)
func (r *SecurityPostureResource) Statistics(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/open/security_posture/statistics", params)
}

// Trends - Security-posture trends (GET /open/security_posture/trends)
func (r *SecurityPostureResource) Trends(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/open/security_posture/trends", params)
}
