package waf

// AlarmResource handles Alarms endpoints.
type AlarmResource struct {
	client *HTTPClient
}

// Get - Get alarm config (GET /commercial/alarm)
func (r *AlarmResource) Get(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/commercial/alarm", params)
}

// Update - Update alarm config (PUT /commercial/alarm)
func (r *AlarmResource) Update(data interface{}) (map[string]interface{}, error) {
	return r.client.Put("/commercial/alarm", data)
}

// Test - Send a test alarm (POST /commercial/alarm/test)
func (r *AlarmResource) Test(data interface{}) (map[string]interface{}, error) {
	return r.client.Post("/commercial/alarm/test", data)
}
