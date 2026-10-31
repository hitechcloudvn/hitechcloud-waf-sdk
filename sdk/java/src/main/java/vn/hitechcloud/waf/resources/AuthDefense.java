package vn.hitechcloud.waf.resources;

import java.util.Map;
import vn.hitechcloud.waf.HttpClient;

/** Dynamic Authentication Defense: Authentication-defense sources, protected users, reviews and records. */
public class AuthDefense extends BaseResource {

    public AuthDefense(HttpClient http) {
        super(http);
    }

    /** GET /open/auth_defense/source - List authentication-defense sources */
    public Map<String, Object> listSources(Map<String, String> params) {
        return http.get("/open/auth_defense/source", params);
    }

    /** POST /open/auth_defense/source - Create a source */
    public Map<String, Object> createSource(Object data) {
        return http.post("/open/auth_defense/source", data);
    }

    /** GET /open/auth_defense/source/password - Get the shared source password policy */
    public Map<String, Object> getSourcePassword(Map<String, String> params) {
        return http.get("/open/auth_defense/source/password", params);
    }

    /** PUT /open/auth_defense/source/password - Update the shared source password policy */
    public Map<String, Object> updateSourcePassword(Object data) {
        return http.put("/open/auth_defense/source/password", data);
    }

    /** GET /open/auth_defense/source/{id} - Get a source */
    public Map<String, Object> getSource(String id, Map<String, String> params) {
        return http.get("/open/auth_defense/source/{id}".replace("{id}", id), params);
    }

    /** PUT /open/auth_defense/source/{id} - Update a source */
    public Map<String, Object> updateSource(String id, Object data) {
        return http.put("/open/auth_defense/source/{id}".replace("{id}", id), data);
    }

    /** DELETE /open/auth_defense/source/{id} - Delete a source */
    public Map<String, Object> deleteSource(String id, Object data) {
        return http.delete("/open/auth_defense/source/{id}".replace("{id}", id), data);
    }

    /** GET /open/auth_defense/source/{id}/secret - Get a source secret */
    public Map<String, Object> getSourceSecret(String id, Map<String, String> params) {
        return http.get("/open/auth_defense/source/{id}/secret".replace("{id}", id), params);
    }

    /** GET /open/auth_defense/source/{id}/user - List users bound to a source */
    public Map<String, Object> listSourceUsers(String id, Map<String, String> params) {
        return http.get("/open/auth_defense/source/{id}/user".replace("{id}", id), params);
    }

    /** PUT /open/auth_defense/source/{id}/user/{user_id} - Update a source user */
    public Map<String, Object> updateSourceUser(String id, String userId, Object data) {
        return http.put("/open/auth_defense/source/{id}/user/{user_id}".replace("{id}", id).replace("{user_id}", userId), data);
    }

    /** DELETE /open/auth_defense/source/{id}/user/{user_id} - Remove a user from a source */
    public Map<String, Object> deleteSourceUser(String id, String userId, Object data) {
        return http.delete("/open/auth_defense/source/{id}/user/{user_id}".replace("{id}", id).replace("{user_id}", userId), data);
    }

    /** GET /open/auth_defense/user - List protected users */
    public Map<String, Object> listUsers(Map<String, String> params) {
        return http.get("/open/auth_defense/user", params);
    }

    /** POST /open/auth_defense/user - Create a protected user */
    public Map<String, Object> createUser(Object data) {
        return http.post("/open/auth_defense/user", data);
    }

    /** DELETE /open/auth_defense/user - Delete protected users (bulk) */
    public Map<String, Object> deleteUsers(Object data) {
        return http.delete("/open/auth_defense/user", data);
    }

    /** POST /open/auth_defense/user/merge - Merge protected users */
    public Map<String, Object> mergeUsers(Object data) {
        return http.post("/open/auth_defense/user/merge", data);
    }

    /** GET /open/auth_defense/user/review - List pending user reviews */
    public Map<String, Object> listUserReviews(Map<String, String> params) {
        return http.get("/open/auth_defense/user/review", params);
    }

    /** PUT /open/auth_defense/user/review/{review_id} - Approve / reject a user review */
    public Map<String, Object> updateUserReview(String reviewId, Object data) {
        return http.put("/open/auth_defense/user/review/{review_id}".replace("{review_id}", reviewId), data);
    }

    /** GET /open/auth_defense/user/{user_id} - Get a protected user */
    public Map<String, Object> getUser(String userId, Map<String, String> params) {
        return http.get("/open/auth_defense/user/{user_id}".replace("{user_id}", userId), params);
    }

    /** PUT /open/auth_defense/user/{user_id} - Update a protected user */
    public Map<String, Object> updateUser(String userId, Object data) {
        return http.put("/open/auth_defense/user/{user_id}".replace("{user_id}", userId), data);
    }

    /** DELETE /open/auth_defense/user/{user_id} - Delete a protected user */
    public Map<String, Object> deleteUser(String userId, Object data) {
        return http.delete("/open/auth_defense/user/{user_id}".replace("{user_id}", userId), data);
    }

    /** PUT /open/auth_defense/user/{user_id}/reset_totp - Reset a user's TOTP */
    public Map<String, Object> resetUserTotp(String userId, Object data) {
        return http.put("/open/auth_defense/user/{user_id}/reset_totp".replace("{user_id}", userId), data);
    }

    /** PUT /open/auth_defense/user/{user_id}/unbind - Unbind a user's authenticator */
    public Map<String, Object> unbindUser(String userId, Object data) {
        return http.put("/open/auth_defense/user/{user_id}/unbind".replace("{user_id}", userId), data);
    }

    /** GET /open/records/auth_defense - Authentication-defense event records */
    public Map<String, Object> records(Map<String, String> params) {
        return http.get("/open/records/auth_defense", params);
    }

    /** GET /open/v2/records/auth_defense - Authentication-defense event records (v2) */
    public Map<String, Object> recordsV2(Map<String, String> params) {
        return http.get("/open/v2/records/auth_defense", params);
    }

    /** POST /business/users/{id}/unlock - Unlock a locked business user */
    public Map<String, Object> unlockUser(String id, Object data) {
        return http.post("/business/users/{id}/unlock".replace("{id}", id), data);
    }

}