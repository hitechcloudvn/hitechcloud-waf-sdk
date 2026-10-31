package vn.hitechcloud.waf.resources;

import java.util.Map;
import vn.hitechcloud.waf.HttpClient;

/** Settings: Login methods, syslog, detector, downgrade and audit log. */
public class Settings extends BaseResource {

    public Settings(HttpClient http) {
        super(http);
    }

    /** PUT /business/system/login_method - Set the admin login method */
    public Map<String, Object> setLoginMethod(Object data) {
        return http.put("/business/system/login_method", data);
    }

    /** POST /business/system/login_method/test - Test the admin login method */
    public Map<String, Object> testLoginMethod(Object data) {
        return http.post("/business/system/login_method/test", data);
    }

    /** GET /commercial/syslog - Get syslog config */
    public Map<String, Object> getSyslog(Map<String, String> params) {
        return http.get("/commercial/syslog", params);
    }

    /** POST /commercial/syslog - Set syslog config */
    public Map<String, Object> setSyslog(Object data) {
        return http.post("/commercial/syslog", data);
    }

    /** POST /commercial/syslog/test - Test syslog delivery */
    public Map<String, Object> testSyslog(Object data) {
        return http.post("/commercial/syslog/test", data);
    }

    /** GET /open/Commercial/downgrade - Get downgrade config */
    public Map<String, Object> getDowngrade(Map<String, String> params) {
        return http.get("/open/Commercial/downgrade", params);
    }

    /** PUT /open/Commercial/downgrade - Set downgrade config */
    public Map<String, Object> setDowngrade(Object data) {
        return http.put("/open/Commercial/downgrade", data);
    }

    /** GET /open/audit_log - Get the admin audit log */
    public Map<String, Object> auditLog(Map<String, String> params) {
        return http.get("/open/audit_log", params);
    }

    /** GET /open/detector - Get detector config */
    public Map<String, Object> getDetector(Map<String, String> params) {
        return http.get("/open/detector", params);
    }

    /** POST /open/detector - Set detector config */
    public Map<String, Object> setDetector(Object data) {
        return http.post("/open/detector", data);
    }

}