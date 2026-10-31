package vn.hitechcloud.waf.resources;

import java.util.Map;
import vn.hitechcloud.waf.HttpClient;

/** Waiting Room: Virtual waiting-room queue for traffic peaks. */
public class WaitingRoom extends BaseResource {

    public WaitingRoom(HttpClient http) {
        super(http);
    }

    /** GET /open/records/waiting - Waiting-room event records */
    public Map<String, Object> records(Map<String, String> params) {
        return http.get("/open/records/waiting", params);
    }

    /** GET /open/site/{id}/waiting - Get waiting-room config for a site */
    public Map<String, Object> get(String id, Map<String, String> params) {
        return http.get("/open/site/{id}/waiting".replace("{id}", id), params);
    }

    /** POST /open/site/{id}/waiting - Set waiting-room config for a site */
    public Map<String, Object> set(String id, Object data) {
        return http.post("/open/site/{id}/waiting".replace("{id}", id), data);
    }

}