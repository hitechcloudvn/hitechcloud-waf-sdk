package vn.hitechcloud.waf.resources;

import java.util.Map;
import vn.hitechcloud.waf.HttpClient;

/** Challenge & Semantics: Human-verification challenge, semantic analysis and chaos protection. */
public class Challenge extends BaseResource {

    public Challenge(HttpClient http) {
        super(http);
    }

    /** GET /open/challenge/config - Get challenge config */
    public Map<String, Object> getConfig(Map<String, String> params) {
        return http.get("/open/challenge/config", params);
    }

    /** POST /open/challenge/config - Set challenge config */
    public Map<String, Object> setConfig(Object data) {
        return http.post("/open/challenge/config", data);
    }

    /** GET /open/records/challenge - Challenge event records */
    public Map<String, Object> records(Map<String, String> params) {
        return http.get("/open/records/challenge", params);
    }

    /** GET /open/site/{id}/semantics - Get semantic-analysis config for a site */
    public Map<String, Object> getSemantics(String id, Map<String, String> params) {
        return http.get("/open/site/{id}/semantics".replace("{id}", id), params);
    }

    /** PUT /open/site/{id}/semantics - Set semantic-analysis config for a site */
    public Map<String, Object> setSemantics(String id, Object data) {
        return http.put("/open/site/{id}/semantics".replace("{id}", id), data);
    }

    /** GET /open/site/{id}/chaos - Get chaos/dynamic-protection config for a site */
    public Map<String, Object> getChaos(String id, Map<String, String> params) {
        return http.get("/open/site/{id}/chaos".replace("{id}", id), params);
    }

    /** POST /open/site/{id}/chaos - Set chaos/dynamic-protection config for a site */
    public Map<String, Object> setChaos(String id, Object data) {
        return http.post("/open/site/{id}/chaos".replace("{id}", id), data);
    }

}