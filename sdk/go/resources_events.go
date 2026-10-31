package waf

import "strings"

// EventsResource handles Attacks & Events endpoints.
type EventsResource struct {
	client *HTTPClient
}

// List - List attack events (GET /open/events)
func (r *EventsResource) List(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/open/events", params)
}

// ListRules - List events grouped by rule (GET /open/events/rule)
func (r *EventsResource) ListRules(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/open/events/rule", params)
}

// GetRecord - Get a single event record (GET /open/record/{id})
func (r *EventsResource) GetRecord(id string, params map[string]string) (map[string]interface{}, error) {
	return r.client.Get(strings.NewReplacer("{id}", id).Replace("/open/record/{id}"), params)
}

// GetRecordRule - Get the rule for a record (GET /open/record/rule/{id})
func (r *EventsResource) GetRecordRule(id string, params map[string]string) (map[string]interface{}, error) {
	return r.client.Get(strings.NewReplacer("{id}", id).Replace("/open/record/rule/{id}"), params)
}

// Records - List raw records (GET /open/records)
func (r *EventsResource) Records(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/open/records", params)
}

// RecordsRule - List raw records grouped by rule (GET /open/records/rule)
func (r *EventsResource) RecordsRule(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/open/records/rule", params)
}

// Export - Export records (commercial) (GET /commercial/record/export)
func (r *EventsResource) Export(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/commercial/record/export", params)
}
