package vn.hitechcloud.waf.resources;

import java.util.Map;
import vn.hitechcloud.waf.HttpClient;

/** Statistics: Advanced access/attack analytics and QPS. */
public class Statistics extends BaseResource {

    public Statistics(HttpClient http) {
        super(http);
    }

    /** GET /stat/advance/access - Access statistics */
    public Map<String, Object> access(Map<String, String> params) {
        return http.get("/stat/advance/access", params);
    }

    /** GET /stat/advance/attack - Attack statistics */
    public Map<String, Object> attack(Map<String, String> params) {
        return http.get("/stat/advance/attack", params);
    }

    /** GET /stat/advance/client - Client statistics */
    public Map<String, Object> client(Map<String, String> params) {
        return http.get("/stat/advance/client", params);
    }

    /** GET /stat/advance/domain - Per-domain statistics */
    public Map<String, Object> domain(Map<String, String> params) {
        return http.get("/stat/advance/domain", params);
    }

    /** GET /stat/advance/error_status_code - Error status-code statistics */
    public Map<String, Object> errorStatusCode(Map<String, String> params) {
        return http.get("/stat/advance/error_status_code", params);
    }

    /** GET /stat/advance/location - Geo-location statistics */
    public Map<String, Object> location(Map<String, String> params) {
        return http.get("/stat/advance/location", params);
    }

    /** GET /stat/advance/page - Per-page statistics */
    public Map<String, Object> page(Map<String, String> params) {
        return http.get("/stat/advance/page", params);
    }

    /** GET /stat/advance/status_code - Status-code statistics */
    public Map<String, Object> statusCode(Map<String, String> params) {
        return http.get("/stat/advance/status_code", params);
    }

    /** GET /stat/advance/trend/access - Access trend */
    public Map<String, Object> trendAccess(Map<String, String> params) {
        return http.get("/stat/advance/trend/access", params);
    }

    /** GET /stat/advance/trend/intercept - Interception trend */
    public Map<String, Object> trendIntercept(Map<String, String> params) {
        return http.get("/stat/advance/trend/intercept", params);
    }

    /** GET /stat/qps - Current QPS */
    public Map<String, Object> qps(Map<String, String> params) {
        return http.get("/stat/qps", params);
    }

}