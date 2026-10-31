package vn.hitechcloud.waf.resources;

import java.util.Map;
import vn.hitechcloud.waf.HttpClient;

/** Global Mode & Proxy: Global protection mode, outbound proxy and log cleanup. */
public class GlobalSettings extends BaseResource {

    public GlobalSettings(HttpClient http) {
        super(http);
    }

    /** GET /open/global/mode - Get the global protection mode */
    public Map<String, Object> getMode(Map<String, String> params) {
        return http.get("/open/global/mode", params);
    }

    /** PUT /open/global/mode - Set the global protection mode */
    public Map<String, Object> setMode(Object data) {
        return http.put("/open/global/mode", data);
    }

    /** GET /open/global/proxy - Get the global proxy config */
    public Map<String, Object> getProxy(Map<String, String> params) {
        return http.get("/open/global/proxy", params);
    }

    /** PUT /open/global/proxy - Set the global proxy config */
    public Map<String, Object> setProxy(Object data) {
        return http.put("/open/global/proxy", data);
    }

    /** GET /open/global/log_clean - Get the log-cleanup policy */
    public Map<String, Object> getLogClean(Map<String, String> params) {
        return http.get("/open/global/log_clean", params);
    }

    /** POST /open/global/log_clean - Set the log-cleanup policy */
    public Map<String, Object> setLogClean(Object data) {
        return http.post("/open/global/log_clean", data);
    }

}