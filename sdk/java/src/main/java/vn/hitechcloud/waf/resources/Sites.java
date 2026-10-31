package vn.hitechcloud.waf.resources;

import java.util.Map;
import vn.hitechcloud.waf.HttpClient;

/** Sites (Web Services): Protected sites, groups, proxy/origin config, logs and forwarding rules. */
public class Sites extends BaseResource {

    public Sites(HttpClient http) {
        super(http);
    }

    /** GET /open/site - List protected sites */
    public Map<String, Object> list(Map<String, String> params) {
        return http.get("/open/site", params);
    }

    /** POST /open/site - Create a protected site */
    public Map<String, Object> create(Object data) {
        return http.post("/open/site", data);
    }

    /** PUT /open/site - Update a site */
    public Map<String, Object> update(Object data) {
        return http.put("/open/site", data);
    }

    /** DELETE /open/site - Delete a site */
    public Map<String, Object> delete(Object data) {
        return http.delete("/open/site", data);
    }

    /** PUT /open/site/challenge - Toggle challenge mode for sites */
    public Map<String, Object> setChallenge(Object data) {
        return http.put("/open/site/challenge", data);
    }

    /** PUT /open/site/defense - Toggle defense mode for sites */
    public Map<String, Object> setDefense(Object data) {
        return http.put("/open/site/defense", data);
    }

    /** GET /open/site/group - List site groups */
    public Map<String, Object> listGroups(Map<String, String> params) {
        return http.get("/open/site/group", params);
    }

    /** POST /open/site/group - Create a site group */
    public Map<String, Object> createGroup(Object data) {
        return http.post("/open/site/group", data);
    }

    /** GET /open/site/group/switch - Get site-group switch state */
    public Map<String, Object> getGroupSwitch(Map<String, String> params) {
        return http.get("/open/site/group/switch", params);
    }

    /** PUT /open/site/group/{id} - Update a site group */
    public Map<String, Object> updateGroup(String id, Object data) {
        return http.put("/open/site/group/{id}".replace("{id}", id), data);
    }

    /** DELETE /open/site/group/{id} - Delete a site group */
    public Map<String, Object> deleteGroup(String id, Object data) {
        return http.delete("/open/site/group/{id}".replace("{id}", id), data);
    }

    /** PUT /open/site/group/{id}/sort - Sort a site group */
    public Map<String, Object> sortGroup(String id, Object data) {
        return http.put("/open/site/group/{id}/sort".replace("{id}", id), data);
    }

    /** PUT /commercial/site/group/switch - Switch the active site group (commercial) */
    public Map<String, Object> switchGroup(Object data) {
        return http.put("/commercial/site/group/switch", data);
    }

    /** POST /open/site/healthcheck - Run an origin health check */
    public Map<String, Object> healthcheck(Object data) {
        return http.post("/open/site/healthcheck", data);
    }

    /** PUT /open/site/mode - Set the protection mode for a site */
    public Map<String, Object> setMode(Object data) {
        return http.put("/open/site/mode", data);
    }

    /** GET /open/site/{id} - Get a site */
    public Map<String, Object> get(String id, Map<String, String> params) {
        return http.get("/open/site/{id}".replace("{id}", id), params);
    }

    /** PUT /open/site/{id}/basic_info - Update site basic info */
    public Map<String, Object> updateBasicInfo(String id, Object data) {
        return http.put("/open/site/{id}/basic_info".replace("{id}", id), data);
    }

    /** GET /open/site/{id}/excludes - List exclusion rules */
    public Map<String, Object> getExcludes(String id, Map<String, String> params) {
        return http.get("/open/site/{id}/excludes".replace("{id}", id), params);
    }

    /** POST /open/site/{id}/excludes - Add an exclusion rule */
    public Map<String, Object> addExclude(String id, Object data) {
        return http.post("/open/site/{id}/excludes".replace("{id}", id), data);
    }

    /** PUT /open/site/{id}/group - Move a site into a group */
    public Map<String, Object> setGroup(String id, Object data) {
        return http.put("/open/site/{id}/group".replace("{id}", id), data);
    }

    /** GET /open/site/{id}/log - Get site logs */
    public Map<String, Object> getLog(String id, Map<String, String> params) {
        return http.get("/open/site/{id}/log".replace("{id}", id), params);
    }

    /** GET /open/site/{id}/log/detail - Get a log detail */
    public Map<String, Object> getLogDetail(String id, Map<String, String> params) {
        return http.get("/open/site/{id}/log/detail".replace("{id}", id), params);
    }

    /** GET /open/site/{id}/log/download - Download site logs */
    public Map<String, Object> downloadLog(String id, Map<String, String> params) {
        return http.get("/open/site/{id}/log/download".replace("{id}", id), params);
    }

    /** GET /open/site/{id}/log/limit - Get log limit config */
    public Map<String, Object> getLogLimit(String id, Map<String, String> params) {
        return http.get("/open/site/{id}/log/limit".replace("{id}", id), params);
    }

    /** POST /open/site/{id}/log/limit - Set log limit config */
    public Map<String, Object> setLogLimit(String id, Object data) {
        return http.post("/open/site/{id}/log/limit".replace("{id}", id), data);
    }

    /** GET /open/site/{id}/nginx_config - Get generated nginx config */
    public Map<String, Object> getNginxConfig(String id, Map<String, String> params) {
        return http.get("/open/site/{id}/nginx_config".replace("{id}", id), params);
    }

    /** PUT /open/site/{id}/nginx_config - Update custom nginx config */
    public Map<String, Object> updateNginxConfig(String id, Object data) {
        return http.put("/open/site/{id}/nginx_config".replace("{id}", id), data);
    }

    /** GET /open/site/{id}/proxy - Get reverse-proxy / origin config */
    public Map<String, Object> getProxy(String id, Map<String, String> params) {
        return http.get("/open/site/{id}/proxy".replace("{id}", id), params);
    }

    /** PUT /open/site/{id}/proxy - Update reverse-proxy / origin config */
    public Map<String, Object> updateProxy(String id, Object data) {
        return http.put("/open/site/{id}/proxy".replace("{id}", id), data);
    }

    /** GET /open/site/{id}/resources - List static resources */
    public Map<String, Object> getResources(String id, Map<String, String> params) {
        return http.get("/open/site/{id}/resources".replace("{id}", id), params);
    }

    /** DELETE /open/site/{id}/resources - Delete static resources */
    public Map<String, Object> deleteResources(String id, Object data) {
        return http.delete("/open/site/{id}/resources".replace("{id}", id), data);
    }

    /** PUT /open/site/{id}/sort - Sort a site */
    public Map<String, Object> sort(String id, Object data) {
        return http.put("/open/site/{id}/sort".replace("{id}", id), data);
    }

    /** GET /open/site/{id}/static - List static-site rules */
    public Map<String, Object> getStatic(String id, Map<String, String> params) {
        return http.get("/open/site/{id}/static".replace("{id}", id), params);
    }

    /** POST /open/site/{id}/static - Add a static-site rule */
    public Map<String, Object> addStatic(String id, Object data) {
        return http.post("/open/site/{id}/static".replace("{id}", id), data);
    }

    /** PUT /open/site/{id}/static - Update a static-site rule */
    public Map<String, Object> updateStatic(String id, Object data) {
        return http.put("/open/site/{id}/static".replace("{id}", id), data);
    }

    /** DELETE /open/site/{id}/static - Delete a static-site rule */
    public Map<String, Object> deleteStatic(String id, Object data) {
        return http.delete("/open/site/{id}/static".replace("{id}", id), data);
    }

    /** GET /open/site/{id}/forwarding_rules - List forwarding rules */
    public Map<String, Object> listForwardingRules(String id, Map<String, String> params) {
        return http.get("/open/site/{id}/forwarding_rules".replace("{id}", id), params);
    }

    /** POST /open/site/{id}/forwarding_rules - Create a forwarding rule */
    public Map<String, Object> createForwardingRule(String id, Object data) {
        return http.post("/open/site/{id}/forwarding_rules".replace("{id}", id), data);
    }

    /** PUT /open/site/{id}/forwarding_rules/{rule_id} - Update a forwarding rule */
    public Map<String, Object> updateForwardingRule(String id, String ruleId, Object data) {
        return http.put("/open/site/{id}/forwarding_rules/{rule_id}".replace("{id}", id).replace("{rule_id}", ruleId), data);
    }

    /** DELETE /open/site/{id}/forwarding_rules/{rule_id} - Delete a forwarding rule */
    public Map<String, Object> deleteForwardingRule(String id, String ruleId, Object data) {
        return http.delete("/open/site/{id}/forwarding_rules/{rule_id}".replace("{id}", id).replace("{rule_id}", ruleId), data);
    }

    /** PUT /open/site/{id}/forwarding_rules/{rule_id}/switch - Enable / disable a forwarding rule */
    public Map<String, Object> switchForwardingRule(String id, String ruleId, Object data) {
        return http.put("/open/site/{id}/forwarding_rules/{rule_id}/switch".replace("{id}", id).replace("{rule_id}", ruleId), data);
    }

}