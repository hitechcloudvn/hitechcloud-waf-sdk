package waf

import "strings"

// UsersResource handles Users & Account endpoints.
type UsersResource struct {
	client *HTTPClient
}

// List - List users (GET /open/users)
func (r *UsersResource) List(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/open/users", params)
}

// Create - Create a user (POST /open/users)
func (r *UsersResource) Create(data interface{}) (map[string]interface{}, error) {
	return r.client.Post("/open/users", data)
}

// Update - Update a user (PUT /open/users)
func (r *UsersResource) Update(data interface{}) (map[string]interface{}, error) {
	return r.client.Put("/open/users", data)
}

// Delete - Delete a user (DELETE /open/users)
func (r *UsersResource) Delete(data interface{}) (map[string]interface{}, error) {
	return r.client.Delete("/open/users", data)
}

// SetupTotp - Enable / bind TOTP for a user (POST /open/users/{id}/totp)
func (r *UsersResource) SetupTotp(id string, data interface{}) (map[string]interface{}, error) {
	return r.client.Post(strings.NewReplacer("{id}", id).Replace("/open/users/{id}/totp"), data)
}

// GetManagerInfo - Get the current manager (admin) profile (GET /ManagerInfo)
func (r *UsersResource) GetManagerInfo(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/ManagerInfo", params)
}

// UpdateManagerInfo - Update the current manager (admin) profile (PUT /ManagerInfo)
func (r *UsersResource) UpdateManagerInfo(data interface{}) (map[string]interface{}, error) {
	return r.client.Put("/ManagerInfo", data)
}
