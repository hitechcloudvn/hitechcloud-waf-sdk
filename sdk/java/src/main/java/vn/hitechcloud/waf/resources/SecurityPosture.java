package vn.hitechcloud.waf.resources;

import java.util.Map;
import vn.hitechcloud.waf.HttpClient;

/** Security Posture: Realtime posture, per-site scoring, statistics and trends. */
public class SecurityPosture extends BaseResource {

    public SecurityPosture(HttpClient http) {
        super(http);
    }

    /** GET /open/security_posture/realtime - Realtime security posture */
    public Map<String, Object> realtime(Map<String, String> params) {
        return http.get("/open/security_posture/realtime", params);
    }

    /** POST /open/security_posture/site - Security posture for selected sites */
    public Map<String, Object> site(Object data) {
        return http.post("/open/security_posture/site", data);
    }

    /** GET /open/security_posture/statistics - Security-posture statistics */
    public Map<String, Object> statistics(Map<String, String> params) {
        return http.get("/open/security_posture/statistics", params);
    }

    /** GET /open/security_posture/trends - Security-posture trends */
    public Map<String, Object> trends(Map<String, String> params) {
        return http.get("/open/security_posture/trends", params);
    }

}