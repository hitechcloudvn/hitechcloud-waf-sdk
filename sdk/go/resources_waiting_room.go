package waf

import "strings"

// WaitingRoomResource handles Waiting Room endpoints.
type WaitingRoomResource struct {
	client *HTTPClient
}

// Records - Waiting-room event records (GET /open/records/waiting)
func (r *WaitingRoomResource) Records(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/open/records/waiting", params)
}

// Get - Get waiting-room config for a site (GET /open/site/{id}/waiting)
func (r *WaitingRoomResource) Get(id string, params map[string]string) (map[string]interface{}, error) {
	return r.client.Get(strings.NewReplacer("{id}", id).Replace("/open/site/{id}/waiting"), params)
}

// Set - Set waiting-room config for a site (POST /open/site/{id}/waiting)
func (r *WaitingRoomResource) Set(id string, data interface{}) (map[string]interface{}, error) {
	return r.client.Post(strings.NewReplacer("{id}", id).Replace("/open/site/{id}/waiting"), data)
}
