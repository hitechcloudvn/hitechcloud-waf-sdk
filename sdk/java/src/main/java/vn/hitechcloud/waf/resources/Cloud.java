package vn.hitechcloud.waf.resources;

import java.util.Map;
import vn.hitechcloud.waf.HttpClient;

/** Cloud Policies: Shared cloud threat-intelligence policies. */
public class Cloud extends BaseResource {

    public Cloud(HttpClient http) {
        super(http);
    }

    /** GET /open/cloud/policies - List available cloud policies */
    public Map<String, Object> listPolicies(Map<String, String> params) {
        return http.get("/open/cloud/policies", params);
    }

    /** POST /open/cloud/policies/subscribe - Subscribe to a cloud policy */
    public Map<String, Object> subscribe(Object data) {
        return http.post("/open/cloud/policies/subscribe", data);
    }

}