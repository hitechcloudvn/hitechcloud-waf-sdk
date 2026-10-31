package vn.hitechcloud.waf.resources;

import java.util.Map;
import vn.hitechcloud.waf.HttpClient;

/** License & Authorization: Authorization records and license status / lifecycle. */
public class License extends BaseResource {

    public License(HttpClient http) {
        super(http);
    }

    /** GET /open/system/authorize - Get authorization records */
    public Map<String, Object> getAuthorize(Map<String, String> params) {
        return http.get("/open/system/authorize", params);
    }

    /** POST /open/system/authorize - Create an authorization */
    public Map<String, Object> createAuthorize(Object data) {
        return http.post("/open/system/authorize", data);
    }

    /** PUT /open/system/authorize - Update an authorization */
    public Map<String, Object> updateAuthorize(Object data) {
        return http.put("/open/system/authorize", data);
    }

    /** DELETE /open/system/authorize - Delete an authorization */
    public Map<String, Object> deleteAuthorize(Object data) {
        return http.delete("/open/system/authorize", data);
    }

    /** POST /open/system/license/expired/confirm-downgrade - Confirm downgrade after license expiry */
    public Map<String, Object> confirmDowngrade(Object data) {
        return http.post("/open/system/license/expired/confirm-downgrade", data);
    }

    /** POST /open/system/license/expiry/acknowledge - Acknowledge an upcoming license expiry */
    public Map<String, Object> acknowledgeExpiry(Object data) {
        return http.post("/open/system/license/expiry/acknowledge", data);
    }

    /** GET /open/system/license/status - Get the license status */
    public Map<String, Object> status(Map<String, String> params) {
        return http.get("/open/system/license/status", params);
    }

}