package waf

import "strings"

// SslResource handles SSL Certificates endpoints.
type SslResource struct {
	client *HTTPClient
}

// List - List certificates (GET /open/cert)
func (r *SslResource) List(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/open/cert", params)
}

// Create - Upload / create a certificate (POST /open/cert)
func (r *SslResource) Create(data interface{}) (map[string]interface{}, error) {
	return r.client.Post("/open/cert", data)
}

// Get - Get a certificate (GET /open/cert/{id})
func (r *SslResource) Get(id string, params map[string]string) (map[string]interface{}, error) {
	return r.client.Get(strings.NewReplacer("{id}", id).Replace("/open/cert/{id}"), params)
}

// Delete - Delete a certificate (DELETE /open/cert/{id})
func (r *SslResource) Delete(id string, data interface{}) (map[string]interface{}, error) {
	return r.client.Delete(strings.NewReplacer("{id}", id).Replace("/open/cert/{id}"), data)
}
