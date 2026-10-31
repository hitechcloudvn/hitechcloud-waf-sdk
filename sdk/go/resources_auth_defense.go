package waf

import "strings"

// AuthDefenseResource handles Dynamic Authentication Defense endpoints.
type AuthDefenseResource struct {
	client *HTTPClient
}

// ListSources - List authentication-defense sources (GET /open/auth_defense/source)
func (r *AuthDefenseResource) ListSources(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/open/auth_defense/source", params)
}

// CreateSource - Create a source (POST /open/auth_defense/source)
func (r *AuthDefenseResource) CreateSource(data interface{}) (map[string]interface{}, error) {
	return r.client.Post("/open/auth_defense/source", data)
}

// GetSourcePassword - Get the shared source password policy (GET /open/auth_defense/source/password)
func (r *AuthDefenseResource) GetSourcePassword(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/open/auth_defense/source/password", params)
}

// UpdateSourcePassword - Update the shared source password policy (PUT /open/auth_defense/source/password)
func (r *AuthDefenseResource) UpdateSourcePassword(data interface{}) (map[string]interface{}, error) {
	return r.client.Put("/open/auth_defense/source/password", data)
}

// GetSource - Get a source (GET /open/auth_defense/source/{id})
func (r *AuthDefenseResource) GetSource(id string, params map[string]string) (map[string]interface{}, error) {
	return r.client.Get(strings.NewReplacer("{id}", id).Replace("/open/auth_defense/source/{id}"), params)
}

// UpdateSource - Update a source (PUT /open/auth_defense/source/{id})
func (r *AuthDefenseResource) UpdateSource(id string, data interface{}) (map[string]interface{}, error) {
	return r.client.Put(strings.NewReplacer("{id}", id).Replace("/open/auth_defense/source/{id}"), data)
}

// DeleteSource - Delete a source (DELETE /open/auth_defense/source/{id})
func (r *AuthDefenseResource) DeleteSource(id string, data interface{}) (map[string]interface{}, error) {
	return r.client.Delete(strings.NewReplacer("{id}", id).Replace("/open/auth_defense/source/{id}"), data)
}

// GetSourceSecret - Get a source secret (GET /open/auth_defense/source/{id}/secret)
func (r *AuthDefenseResource) GetSourceSecret(id string, params map[string]string) (map[string]interface{}, error) {
	return r.client.Get(strings.NewReplacer("{id}", id).Replace("/open/auth_defense/source/{id}/secret"), params)
}

// ListSourceUsers - List users bound to a source (GET /open/auth_defense/source/{id}/user)
func (r *AuthDefenseResource) ListSourceUsers(id string, params map[string]string) (map[string]interface{}, error) {
	return r.client.Get(strings.NewReplacer("{id}", id).Replace("/open/auth_defense/source/{id}/user"), params)
}

// UpdateSourceUser - Update a source user (PUT /open/auth_defense/source/{id}/user/{user_id})
func (r *AuthDefenseResource) UpdateSourceUser(id string, userId string, data interface{}) (map[string]interface{}, error) {
	return r.client.Put(strings.NewReplacer("{id}", id, "{user_id}", userId).Replace("/open/auth_defense/source/{id}/user/{user_id}"), data)
}

// DeleteSourceUser - Remove a user from a source (DELETE /open/auth_defense/source/{id}/user/{user_id})
func (r *AuthDefenseResource) DeleteSourceUser(id string, userId string, data interface{}) (map[string]interface{}, error) {
	return r.client.Delete(strings.NewReplacer("{id}", id, "{user_id}", userId).Replace("/open/auth_defense/source/{id}/user/{user_id}"), data)
}

// ListUsers - List protected users (GET /open/auth_defense/user)
func (r *AuthDefenseResource) ListUsers(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/open/auth_defense/user", params)
}

// CreateUser - Create a protected user (POST /open/auth_defense/user)
func (r *AuthDefenseResource) CreateUser(data interface{}) (map[string]interface{}, error) {
	return r.client.Post("/open/auth_defense/user", data)
}

// DeleteUsers - Delete protected users (bulk) (DELETE /open/auth_defense/user)
func (r *AuthDefenseResource) DeleteUsers(data interface{}) (map[string]interface{}, error) {
	return r.client.Delete("/open/auth_defense/user", data)
}

// MergeUsers - Merge protected users (POST /open/auth_defense/user/merge)
func (r *AuthDefenseResource) MergeUsers(data interface{}) (map[string]interface{}, error) {
	return r.client.Post("/open/auth_defense/user/merge", data)
}

// ListUserReviews - List pending user reviews (GET /open/auth_defense/user/review)
func (r *AuthDefenseResource) ListUserReviews(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/open/auth_defense/user/review", params)
}

// UpdateUserReview - Approve / reject a user review (PUT /open/auth_defense/user/review/{review_id})
func (r *AuthDefenseResource) UpdateUserReview(reviewId string, data interface{}) (map[string]interface{}, error) {
	return r.client.Put(strings.NewReplacer("{review_id}", reviewId).Replace("/open/auth_defense/user/review/{review_id}"), data)
}

// GetUser - Get a protected user (GET /open/auth_defense/user/{user_id})
func (r *AuthDefenseResource) GetUser(userId string, params map[string]string) (map[string]interface{}, error) {
	return r.client.Get(strings.NewReplacer("{user_id}", userId).Replace("/open/auth_defense/user/{user_id}"), params)
}

// UpdateUser - Update a protected user (PUT /open/auth_defense/user/{user_id})
func (r *AuthDefenseResource) UpdateUser(userId string, data interface{}) (map[string]interface{}, error) {
	return r.client.Put(strings.NewReplacer("{user_id}", userId).Replace("/open/auth_defense/user/{user_id}"), data)
}

// DeleteUser - Delete a protected user (DELETE /open/auth_defense/user/{user_id})
func (r *AuthDefenseResource) DeleteUser(userId string, data interface{}) (map[string]interface{}, error) {
	return r.client.Delete(strings.NewReplacer("{user_id}", userId).Replace("/open/auth_defense/user/{user_id}"), data)
}

// ResetUserTotp - Reset a user's TOTP (PUT /open/auth_defense/user/{user_id}/reset_totp)
func (r *AuthDefenseResource) ResetUserTotp(userId string, data interface{}) (map[string]interface{}, error) {
	return r.client.Put(strings.NewReplacer("{user_id}", userId).Replace("/open/auth_defense/user/{user_id}/reset_totp"), data)
}

// UnbindUser - Unbind a user's authenticator (PUT /open/auth_defense/user/{user_id}/unbind)
func (r *AuthDefenseResource) UnbindUser(userId string, data interface{}) (map[string]interface{}, error) {
	return r.client.Put(strings.NewReplacer("{user_id}", userId).Replace("/open/auth_defense/user/{user_id}/unbind"), data)
}

// Records - Authentication-defense event records (GET /open/records/auth_defense)
func (r *AuthDefenseResource) Records(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/open/records/auth_defense", params)
}

// RecordsV2 - Authentication-defense event records (v2) (GET /open/v2/records/auth_defense)
func (r *AuthDefenseResource) RecordsV2(params map[string]string) (map[string]interface{}, error) {
	return r.client.Get("/open/v2/records/auth_defense", params)
}

// UnlockUser - Unlock a locked business user (POST /business/users/{id}/unlock)
func (r *AuthDefenseResource) UnlockUser(id string, data interface{}) (map[string]interface{}, error) {
	return r.client.Post(strings.NewReplacer("{id}", id).Replace("/business/users/{id}/unlock"), data)
}
