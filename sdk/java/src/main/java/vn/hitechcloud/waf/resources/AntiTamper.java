package vn.hitechcloud.waf.resources;

import java.util.Map;
import vn.hitechcloud.waf.HttpClient;

/** Anti-Tamper: Content anti-tamper protection per site and resource. */
public class AntiTamper extends BaseResource {

    public AntiTamper(HttpClient http) {
        super(http);
    }

    /** GET /business/anti_tamper/{id} - Get an anti-tamper config */
    public Map<String, Object> get(String id, Map<String, String> params) {
        return http.get("/business/anti_tamper/{id}".replace("{id}", id), params);
    }

    /** PUT /business/anti_tamper/{id} - Update an anti-tamper config */
    public Map<String, Object> update(String id, Object data) {
        return http.put("/business/anti_tamper/{id}".replace("{id}", id), data);
    }

    /** DELETE /business/anti_tamper/{id} - Delete an anti-tamper config */
    public Map<String, Object> delete(String id, Object data) {
        return http.delete("/business/anti_tamper/{id}".replace("{id}", id), data);
    }

    /** GET /business/anti_tamper/{id}/page - Get protected pages for a config */
    public Map<String, Object> getPage(String id, Map<String, String> params) {
        return http.get("/business/anti_tamper/{id}/page".replace("{id}", id), params);
    }

    /** GET /business/site/{site_id}/anti_tamper - List anti-tamper configs for a site */
    public Map<String, Object> listForSite(String siteId, Map<String, String> params) {
        return http.get("/business/site/{site_id}/anti_tamper".replace("{site_id}", siteId), params);
    }

    /** POST /business/site/{site_id}/anti_tamper - Create an anti-tamper config for a site */
    public Map<String, Object> createForSite(String siteId, Object data) {
        return http.post("/business/site/{site_id}/anti_tamper".replace("{site_id}", siteId), data);
    }

    /** PUT /business/site/{site_id}/anti_tamper - Update anti-tamper for a site */
    public Map<String, Object> updateForSite(String siteId, Object data) {
        return http.put("/business/site/{site_id}/anti_tamper".replace("{site_id}", siteId), data);
    }

    /** GET /business/site/{site_id}/anti_tamper/resource/{resource_id} - Get a protected resource */
    public Map<String, Object> getResource(String siteId, String resourceId, Map<String, String> params) {
        return http.get("/business/site/{site_id}/anti_tamper/resource/{resource_id}".replace("{site_id}", siteId).replace("{resource_id}", resourceId), params);
    }

}