package waf

// PoliciesResource handles Policies & Rules endpoints.
type PoliciesResource struct {
	client *HTTPClient
}

// List - List policies (GET /open/policy)
func (r *PoliciesResource) List(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/open/policy", params)
}

// Create - Create a policy (POST /open/policy)
func (r *PoliciesResource) Create(data interface{}) (map[string]interface{}, error) {
	return r.client.Post("/open/policy", data)
}

// Update - Update a policy (PUT /open/policy)
func (r *PoliciesResource) Update(data interface{}) (map[string]interface{}, error) {
	return r.client.Put("/open/policy", data)
}

// Delete - Delete a policy (DELETE /open/policy)
func (r *PoliciesResource) Delete(data interface{}) (map[string]interface{}, error) {
	return r.client.Delete("/open/policy", data)
}

// Detail - Get policy detail (GET /open/policy/detail)
func (r *PoliciesResource) Detail(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/open/policy/detail", params)
}

// SetOrder - Set a policy order (PUT /open/policy/order)
func (r *PoliciesResource) SetOrder(data interface{}) (map[string]interface{}, error) {
	return r.client.Put("/open/policy/order", data)
}

// SetOrderValue - Set a policy order value (PUT /open/policy/order/value)
func (r *PoliciesResource) SetOrderValue(data interface{}) (map[string]interface{}, error) {
	return r.client.Put("/open/policy/order/value", data)
}

// SetOrders - Bulk-set policy orders (PUT /open/policy/orders)
func (r *PoliciesResource) SetOrders(data interface{}) (map[string]interface{}, error) {
	return r.client.Put("/open/policy/orders", data)
}

// SetSwitch - Toggle a policy on/off (PUT /open/policy/switch)
func (r *PoliciesResource) SetSwitch(data interface{}) (map[string]interface{}, error) {
	return r.client.Put("/open/policy/switch", data)
}
