package vn.hitechcloud.waf.resources;

import java.util.Map;
import vn.hitechcloud.waf.HttpClient;

/** Authentication & Token: CSRF, login, TFA, logout and API token lifecycle. */
public class Auth extends BaseResource {

    public Auth(HttpClient http) {
        super(http);
    }

    /** GET /open/auth/csrf - Get a CSRF token for the login step */
    public Map<String, Object> csrf(Map<String, String> params) {
        return http.get("/open/auth/csrf", params);
    }

    /** POST /open/auth/login - Authenticate with username/password (+csrf_token) and obtain a JWT */
    public Map<String, Object> login(Object data) {
        return http.post("/open/auth/login", data);
    }

    /** POST /open/auth/logout - End the current session */
    public Map<String, Object> logout(Object data) {
        return http.post("/open/auth/logout", data);
    }

    /** POST /open/auth/tfa - Verify a two-factor (OTP) code during login */
    public Map<String, Object> tfa(Object data) {
        return http.post("/open/auth/tfa", data);
    }

    /** GET /open/auth/token - Get the current API token */
    public Map<String, Object> getToken(Map<String, String> params) {
        return http.get("/open/auth/token", params);
    }

    /** PUT /open/auth/token - Rotate / refresh the API token */
    public Map<String, Object> refreshToken(Object data) {
        return http.put("/open/auth/token", data);
    }

    /** DELETE /open/auth/token - Revoke the API token */
    public Map<String, Object> revokeToken(Object data) {
        return http.delete("/open/auth/token", data);
    }

}