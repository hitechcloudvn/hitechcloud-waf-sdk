package vn.hitechcloud.waf.resources;

import java.util.Map;
import vn.hitechcloud.waf.HttpClient;

/** Users & Account: Manage administrative users, TOTP and the current manager profile. */
public class Users extends BaseResource {

    public Users(HttpClient http) {
        super(http);
    }

    /** GET /open/users - List users */
    public Map<String, Object> list(Map<String, String> params) {
        return http.get("/open/users", params);
    }

    /** POST /open/users - Create a user */
    public Map<String, Object> create(Object data) {
        return http.post("/open/users", data);
    }

    /** PUT /open/users - Update a user */
    public Map<String, Object> update(Object data) {
        return http.put("/open/users", data);
    }

    /** DELETE /open/users - Delete a user */
    public Map<String, Object> delete(Object data) {
        return http.delete("/open/users", data);
    }

    /** POST /open/users/{id}/totp - Enable / bind TOTP for a user */
    public Map<String, Object> setupTotp(String id, Object data) {
        return http.post("/open/users/{id}/totp".replace("{id}", id), data);
    }

    /** GET /ManagerInfo - Get the current manager (admin) profile */
    public Map<String, Object> getManagerInfo(Map<String, String> params) {
        return http.get("/ManagerInfo", params);
    }

    /** PUT /ManagerInfo - Update the current manager (admin) profile */
    public Map<String, Object> updateManagerInfo(Object data) {
        return http.put("/ManagerInfo", data);
    }

}