package vn.hitechcloud.waf.resources;

import java.util.Map;
import vn.hitechcloud.waf.HttpClient;

/** ACL & Rate-Limit: Access-control / rate-limit rules and block relief. */
public class Acl extends BaseResource {

    public Acl(HttpClient http) {
        super(http);
    }

    /** PUT /acl/relieve - Relieve (unblock) an ACL-blocked client */
    public Map<String, Object> relieve(Object data) {
        return http.put("/acl/relieve", data);
    }

    /** GET /open/records/acl - ACL hit records */
    public Map<String, Object> records(Map<String, String> params) {
        return http.get("/open/records/acl", params);
    }

    /** GET /open/site/{id}/acl - List ACL / rate-limit rules for a site */
    public Map<String, Object> list(String id, Map<String, String> params) {
        return http.get("/open/site/{id}/acl".replace("{id}", id), params);
    }

    /** POST /open/site/{id}/acl - Create an ACL rule */
    public Map<String, Object> create(String id, Object data) {
        return http.post("/open/site/{id}/acl".replace("{id}", id), data);
    }

    /** PUT /open/site/{id}/acl - Update ACL rules for a site */
    public Map<String, Object> update(String id, Object data) {
        return http.put("/open/site/{id}/acl".replace("{id}", id), data);
    }

    /** PUT /open/site/{id}/acl/{rule_id} - Update a single ACL rule */
    public Map<String, Object> updateRule(String id, String ruleId, Object data) {
        return http.put("/open/site/{id}/acl/{rule_id}".replace("{id}", id).replace("{rule_id}", ruleId), data);
    }

    /** DELETE /open/site/{id}/acl/{rule_id} - Delete a single ACL rule */
    public Map<String, Object> deleteRule(String id, String ruleId, Object data) {
        return http.delete("/open/site/{id}/acl/{rule_id}".replace("{id}", id).replace("{rule_id}", ruleId), data);
    }

}