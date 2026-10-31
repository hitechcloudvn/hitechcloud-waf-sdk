package vn.hitechcloud.waf.resources;

import java.util.Map;
import vn.hitechcloud.waf.HttpClient;

/** Reports: Scheduled / on-demand security reports. */
public class Report extends BaseResource {

    public Report(HttpClient http) {
        super(http);
    }

    /** GET /business/report - List reports */
    public Map<String, Object> list(Map<String, String> params) {
        return http.get("/business/report", params);
    }

    /** POST /business/report - Create a report */
    public Map<String, Object> create(Object data) {
        return http.post("/business/report", data);
    }

    /** GET /business/report/{id} - Get a report */
    public Map<String, Object> get(String id, Map<String, String> params) {
        return http.get("/business/report/{id}".replace("{id}", id), params);
    }

    /** DELETE /business/report/{id} - Delete a report */
    public Map<String, Object> delete(String id, Object data) {
        return http.delete("/business/report/{id}".replace("{id}", id), data);
    }

}