package vn.hitechcloud.waf.resources;

import java.util.Map;
import vn.hitechcloud.waf.HttpClient;

/** Alarms: Alert / notification configuration. */
public class Alarm extends BaseResource {

    public Alarm(HttpClient http) {
        super(http);
    }

    /** GET /commercial/alarm - Get alarm config */
    public Map<String, Object> get(Map<String, String> params) {
        return http.get("/commercial/alarm", params);
    }

    /** PUT /commercial/alarm - Update alarm config */
    public Map<String, Object> update(Object data) {
        return http.put("/commercial/alarm", data);
    }

    /** POST /commercial/alarm/test - Send a test alarm */
    public Map<String, Object> test(Object data) {
        return http.post("/commercial/alarm/test", data);
    }

}