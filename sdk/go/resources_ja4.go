package waf

// Ja4Resource handles JA4 Fingerprint endpoints.
type Ja4Resource struct {
	client *HTTPClient
}

// Get - Get JA4 fingerprint rules (GET /open/ja4)
func (r *Ja4Resource) Get(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/open/ja4", params)
}

// Update - Update JA4 fingerprint rules (PUT /open/ja4)
func (r *Ja4Resource) Update(data interface{}) (map[string]interface{}, error) {
	return r.client.Put("/open/ja4", data)
}
