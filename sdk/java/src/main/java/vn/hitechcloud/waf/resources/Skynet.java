package vn.hitechcloud.waf.resources;

import java.util.Map;
import vn.hitechcloud.waf.HttpClient;

/** Skynet AI Rules: AI-driven Skynet detection rules and switches. */
public class Skynet extends BaseResource {

    public Skynet(HttpClient http) {
        super(http);
    }

    /** GET /open/skynet/rule - Get Skynet rules */
    public Map<String, Object> getRules(Map<String, String> params) {
        return http.get("/open/skynet/rule", params);
    }

    /** PUT /open/skynet/rule - Update Skynet rules */
    public Map<String, Object> updateRules(Object data) {
        return http.put("/open/skynet/rule", data);
    }

    /** GET /open/skynet/rule/switch - Get Skynet rule switch state */
    public Map<String, Object> getRuleSwitch(Map<String, String> params) {
        return http.get("/open/skynet/rule/switch", params);
    }

    /** PUT /open/skynet/rule/switch - Toggle Skynet rules */
    public Map<String, Object> setRuleSwitch(Object data) {
        return http.put("/open/skynet/rule/switch", data);
    }

}