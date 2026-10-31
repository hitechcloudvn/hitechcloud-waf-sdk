package waf

// McpResource handles MCP endpoints.
type McpResource struct {
	client *HTTPClient
}

// Get - Get MCP capability metadata (GET /open/mcp)
func (r *McpResource) Get(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/open/mcp", params)
}

// Call - Invoke an MCP action (POST /open/mcp)
func (r *McpResource) Call(data interface{}) (map[string]interface{}, error) {
	return r.client.Post("/open/mcp", data)
}
