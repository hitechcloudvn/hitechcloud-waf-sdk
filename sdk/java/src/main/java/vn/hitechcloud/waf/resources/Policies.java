package vn.hitechcloud.waf.resources;

import java.util.Map;
import vn.hitechcloud.waf.HttpClient;

/** Policies & Rules: Custom protection policies and rule ordering. */
public class Policies extends BaseResource {

    public Policies(HttpClient http) {
        super(http);
    }

    /** GET /open/policy - List policies */
    public Map<String, Object> list(Map<String, String> params) {
        return http.get("/open/policy", params);
    }

    /** POST /open/policy - Create a policy */
    public Map<String, Object> create(Object data) {
        return http.post("/open/policy", data);
    }

    /** PUT /open/policy - Update a policy */
    public Map<String, Object> update(Object data) {
        return http.put("/open/policy", data);
    }

    /** DELETE /open/policy - Delete a policy */
    public Map<String, Object> delete(Object data) {
        return http.delete("/open/policy", data);
    }

    /** GET /open/policy/detail - Get policy detail */
    public Map<String, Object> detail(Map<String, String> params) {
        return http.get("/open/policy/detail", params);
    }

    /** PUT /open/policy/order - Set a policy order */
    public Map<String, Object> setOrder(Object data) {
        return http.put("/open/policy/order", data);
    }

    /** PUT /open/policy/order/value - Set a policy order value */
    public Map<String, Object> setOrderValue(Object data) {
        return http.put("/open/policy/order/value", data);
    }

    /** PUT /open/policy/orders - Bulk-set policy orders */
    public Map<String, Object> setOrders(Object data) {
        return http.put("/open/policy/orders", data);
    }

    /** PUT /open/policy/switch - Toggle a policy on/off */
    public Map<String, Object> setSwitch(Object data) {
        return http.put("/open/policy/switch", data);
    }

}