package vn.hitechcloud.waf.resources;

import java.util.Map;
import vn.hitechcloud.waf.HttpClient;

/** Portal & Branding: Self-service portal, block pages, branding and accounts. */
public class Portal extends BaseResource {

    public Portal(HttpClient http) {
        super(http);
    }

    /** PUT /commercial/portal/style - Update portal style (commercial) */
    public Map<String, Object> updateStyle(Object data) {
        return http.put("/commercial/portal/style", data);
    }

    /** GET /open/portal - Get portal config */
    public Map<String, Object> get(Map<String, String> params) {
        return http.get("/open/portal", params);
    }

    /** PUT /open/portal - Update portal config */
    public Map<String, Object> update(Object data) {
        return http.put("/open/portal", data);
    }

    /** GET /open/portal/proxy_config - Get portal proxy config */
    public Map<String, Object> getProxyConfig(Map<String, String> params) {
        return http.get("/open/portal/proxy_config", params);
    }

    /** PUT /open/portal/proxy_config - Set portal proxy config */
    public Map<String, Object> setProxyConfig(Object data) {
        return http.put("/open/portal/proxy_config", data);
    }

    /** GET /open/portal/style - Get portal style */
    public Map<String, Object> getStyle(Map<String, String> params) {
        return http.get("/open/portal/style", params);
    }

    /** PUT /commercial/block_page - Update the block page (commercial) */
    public Map<String, Object> updateBlockPage(Object data) {
        return http.put("/commercial/block_page", data);
    }

    /** GET /open/block_page - Get the active block page */
    public Map<String, Object> getBlockPage(Map<String, String> params) {
        return http.get("/open/block_page", params);
    }

    /** GET /open/block_page_list - List available block pages */
    public Map<String, Object> listBlockPages(Map<String, String> params) {
        return http.get("/open/block_page_list", params);
    }

    /** GET /business/frontend_style - Get frontend style */
    public Map<String, Object> getFrontendStyle(Map<String, String> params) {
        return http.get("/business/frontend_style", params);
    }

    /** PUT /business/frontend_style - Set frontend style */
    public Map<String, Object> setFrontendStyle(Object data) {
        return http.put("/business/frontend_style", data);
    }

    /** GET /business/account - Get portal account(s) */
    public Map<String, Object> getAccount(Map<String, String> params) {
        return http.get("/business/account", params);
    }

    /** POST /business/account - Create a portal account */
    public Map<String, Object> createAccount(Object data) {
        return http.post("/business/account", data);
    }

}