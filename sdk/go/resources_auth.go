package waf

// AuthResource handles Authentication & Token endpoints.
type AuthResource struct {
	client *HTTPClient
}

// Csrf - Get a CSRF token for the login step (GET /open/auth/csrf)
func (r *AuthResource) Csrf(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/open/auth/csrf", params)
}

// Login - Authenticate with username/password (+csrf_token) and obtain a JWT (POST /open/auth/login)
func (r *AuthResource) Login(data interface{}) (map[string]interface{}, error) {
	return r.client.Post("/open/auth/login", data)
}

// Logout - End the current session (POST /open/auth/logout)
func (r *AuthResource) Logout(data interface{}) (map[string]interface{}, error) {
	return r.client.Post("/open/auth/logout", data)
}

// Tfa - Verify a two-factor (OTP) code during login (POST /open/auth/tfa)
func (r *AuthResource) Tfa(data interface{}) (map[string]interface{}, error) {
	return r.client.Post("/open/auth/tfa", data)
}

// GetToken - Get the current API token (GET /open/auth/token)
func (r *AuthResource) GetToken(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/open/auth/token", params)
}

// RefreshToken - Rotate / refresh the API token (PUT /open/auth/token)
func (r *AuthResource) RefreshToken(data interface{}) (map[string]interface{}, error) {
	return r.client.Put("/open/auth/token", data)
}

// RevokeToken - Revoke the API token (DELETE /open/auth/token)
func (r *AuthResource) RevokeToken(data interface{}) (map[string]interface{}, error) {
	return r.client.Delete("/open/auth/token", data)
}
