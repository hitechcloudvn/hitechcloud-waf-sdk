package vn.hitechcloud.waf.resources;

import java.util.Map;
import vn.hitechcloud.waf.HttpClient;

/** Anti-Bot & Intelligence: Threat intelligence, IP groups, crawlers and behaviour/fingerprint sharing. */
public class AntiBot extends BaseResource {

    public AntiBot(HttpClient http) {
        super(http);
    }

    /** GET /open/intelligence - Get threat-intelligence config */
    public Map<String, Object> getIntelligence(Map<String, String> params) {
        return http.get("/open/intelligence", params);
    }

    /** POST /open/intelligence - Set threat-intelligence config */
    public Map<String, Object> setIntelligence(Object data) {
        return http.post("/open/intelligence", data);
    }

    /** POST /open/intelligence/ip_lib - Upload a custom IP library */
    public Map<String, Object> uploadIpLib(Object data) {
        return http.post("/open/intelligence/ip_lib", data);
    }

    /** GET /open/ipgroup - List IP groups */
    public Map<String, Object> listIpGroups(Map<String, String> params) {
        return http.get("/open/ipgroup", params);
    }

    /** POST /open/ipgroup - Create an IP group */
    public Map<String, Object> createIpGroup(Object data) {
        return http.post("/open/ipgroup", data);
    }

    /** PUT /open/ipgroup - Update an IP group */
    public Map<String, Object> updateIpGroup(Object data) {
        return http.put("/open/ipgroup", data);
    }

    /** DELETE /open/ipgroup - Delete an IP group */
    public Map<String, Object> deleteIpGroup(Object data) {
        return http.delete("/open/ipgroup", data);
    }

    /** POST /open/ipgroup/append - Append entries to an IP group */
    public Map<String, Object> appendIpGroup(Object data) {
        return http.post("/open/ipgroup/append", data);
    }

    /** GET /open/ipgroup/crawler - Get known-crawler groups */
    public Map<String, Object> getCrawlers(Map<String, String> params) {
        return http.get("/open/ipgroup/crawler", params);
    }

    /** POST /open/ipgroup/crawler - Configure known-crawler groups */
    public Map<String, Object> setCrawlers(Object data) {
        return http.post("/open/ipgroup/crawler", data);
    }

    /** GET /open/ipgroup/detail - Get IP-group detail */
    public Map<String, Object> getIpGroupDetail(Map<String, String> params) {
        return http.get("/open/ipgroup/detail", params);
    }

    /** GET /open/ipgroup/link - Get IP-group link bindings */
    public Map<String, Object> getIpGroupLink(Map<String, String> params) {
        return http.get("/open/ipgroup/link", params);
    }

    /** POST /open/ipgroup/link - Set IP-group link bindings */
    public Map<String, Object> setIpGroupLink(Object data) {
        return http.post("/open/ipgroup/link", data);
    }

    /** GET /open/ipgroup/tserver - Get trusted-server IP groups */
    public Map<String, Object> getTrustServers(Map<String, String> params) {
        return http.get("/open/ipgroup/tserver", params);
    }

    /** POST /open/ipgroup/tserver - Set trusted-server IP groups */
    public Map<String, Object> setTrustServers(Object data) {
        return http.post("/open/ipgroup/tserver", data);
    }

    /** GET /open/share_behaviour - Get behaviour-sharing config */
    public Map<String, Object> getShareBehaviour(Map<String, String> params) {
        return http.get("/open/share_behaviour", params);
    }

    /** POST /open/share_behaviour - Set behaviour-sharing config */
    public Map<String, Object> setShareBehaviour(Object data) {
        return http.post("/open/share_behaviour", data);
    }

    /** GET /open/share_fingerprint - Get fingerprint-sharing config */
    public Map<String, Object> getShareFingerprint(Map<String, String> params) {
        return http.get("/open/share_fingerprint", params);
    }

    /** POST /open/share_fingerprint - Set fingerprint-sharing config */
    public Map<String, Object> setShareFingerprint(Object data) {
        return http.post("/open/share_fingerprint", data);
    }

}