package waf

// StatisticsResource handles Statistics endpoints.
type StatisticsResource struct {
	client *HTTPClient
}

// Access - Access statistics (GET /stat/advance/access)
func (r *StatisticsResource) Access(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/stat/advance/access", params)
}

// Attack - Attack statistics (GET /stat/advance/attack)
func (r *StatisticsResource) Attack(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/stat/advance/attack", params)
}

// Client - Client statistics (GET /stat/advance/client)
func (r *StatisticsResource) Client(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/stat/advance/client", params)
}

// Domain - Per-domain statistics (GET /stat/advance/domain)
func (r *StatisticsResource) Domain(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/stat/advance/domain", params)
}

// ErrorStatusCode - Error status-code statistics (GET /stat/advance/error_status_code)
func (r *StatisticsResource) ErrorStatusCode(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/stat/advance/error_status_code", params)
}

// Location - Geo-location statistics (GET /stat/advance/location)
func (r *StatisticsResource) Location(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/stat/advance/location", params)
}

// Page - Per-page statistics (GET /stat/advance/page)
func (r *StatisticsResource) Page(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/stat/advance/page", params)
}

// StatusCode - Status-code statistics (GET /stat/advance/status_code)
func (r *StatisticsResource) StatusCode(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/stat/advance/status_code", params)
}

// TrendAccess - Access trend (GET /stat/advance/trend/access)
func (r *StatisticsResource) TrendAccess(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/stat/advance/trend/access", params)
}

// TrendIntercept - Interception trend (GET /stat/advance/trend/intercept)
func (r *StatisticsResource) TrendIntercept(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/stat/advance/trend/intercept", params)
}

// Qps - Current QPS (GET /stat/qps)
func (r *StatisticsResource) Qps(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/stat/qps", params)
}
