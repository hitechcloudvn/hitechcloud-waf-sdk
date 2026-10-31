package waf

import "strings"

// ReportResource handles Reports endpoints.
type ReportResource struct {
	client *HTTPClient
}

// List - List reports (GET /business/report)
func (r *ReportResource) List(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/business/report", params)
}

// Create - Create a report (POST /business/report)
func (r *ReportResource) Create(data interface{}) (map[string]interface{}, error) {
	return r.client.Post("/business/report", data)
}

// Get - Get a report (GET /business/report/{id})
func (r *ReportResource) Get(id string, params map[string]string) (map[string]interface{}, error) {
	return r.client.Get(strings.NewReplacer("{id}", id).Replace("/business/report/{id}"), params)
}

// Delete - Delete a report (DELETE /business/report/{id})
func (r *ReportResource) Delete(id string, data interface{}) (map[string]interface{}, error) {
	return r.client.Delete(strings.NewReplacer("{id}", id).Replace("/business/report/{id}"), data)
}
