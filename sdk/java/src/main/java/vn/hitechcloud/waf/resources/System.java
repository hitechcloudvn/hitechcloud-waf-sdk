package vn.hitechcloud.waf.resources;

import java.util.Map;
import vn.hitechcloud.waf.HttpClient;

/** System: System info, edition, protocols, network proxy and helpers. */
public class System extends BaseResource {

    public System(HttpClient http) {
        super(http);
    }

    /** GET /open/system - Get system settings */
    public Map<String, Object> get(Map<String, String> params) {
        return http.get("/open/system", params);
    }

    /** PUT /open/system - Update system settings */
    public Map<String, Object> update(Object data) {
        return http.put("/open/system", data);
    }

    /** GET /open/system/arch - Get the cluster architecture */
    public Map<String, Object> getArch(Map<String, String> params) {
        return http.get("/open/system/arch", params);
    }

    /** GET /open/system/customer_service - Get customer-service config */
    public Map<String, Object> getCustomerService(Map<String, String> params) {
        return http.get("/open/system/customer_service", params);
    }

    /** PUT /open/system/customer_service - Set customer-service config */
    public Map<String, Object> setCustomerService(Object data) {
        return http.put("/open/system/customer_service", data);
    }

    /** GET /open/system/edition - Get the product edition */
    public Map<String, Object> getEdition(Map<String, String> params) {
        return http.get("/open/system/edition", params);
    }

    /** GET /open/system/key - Get the system public key */
    public Map<String, Object> getKey(Map<String, String> params) {
        return http.get("/open/system/key", params);
    }

    /** GET /open/system/login_method - Get the system login method */
    public Map<String, Object> getLoginMethod(Map<String, String> params) {
        return http.get("/open/system/login_method", params);
    }

    /** GET /open/system/login_type - Get the system login type */
    public Map<String, Object> getLoginType(Map<String, String> params) {
        return http.get("/open/system/login_type", params);
    }

    /** GET /open/system/network_proxy - Get the network proxy */
    public Map<String, Object> getNetworkProxy(Map<String, String> params) {
        return http.get("/open/system/network_proxy", params);
    }

    /** POST /open/system/network_proxy - Set the network proxy */
    public Map<String, Object> setNetworkProxy(Object data) {
        return http.post("/open/system/network_proxy", data);
    }

    /** GET /open/system/protocol - Get supported protocols */
    public Map<String, Object> getProtocol(Map<String, String> params) {
        return http.get("/open/system/protocol", params);
    }

    /** POST /open/system/protocol - Configure supported protocols */
    public Map<String, Object> setProtocol(Object data) {
        return http.post("/open/system/protocol", data);
    }

    /** POST /commercial/dashboard/token - Issue a dashboard token */
    public Map<String, Object> dashboardToken(Object data) {
        return http.post("/commercial/dashboard/token", data);
    }

    /** POST /open/safepoint/consult - Consult the safepoint advisor */
    public Map<String, Object> safepointConsult(Object data) {
        return http.post("/open/safepoint/consult", data);
    }

}