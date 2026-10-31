package vn.hitechcloud.waf.resources;

import java.util.Map;
import vn.hitechcloud.waf.HttpClient;

/** SSL Certificates: TLS/SSL certificate management. */
public class Ssl extends BaseResource {

    public Ssl(HttpClient http) {
        super(http);
    }

    /** GET /open/cert - List certificates */
    public Map<String, Object> list(Map<String, String> params) {
        return http.get("/open/cert", params);
    }

    /** POST /open/cert - Upload / create a certificate */
    public Map<String, Object> create(Object data) {
        return http.post("/open/cert", data);
    }

    /** GET /open/cert/{id} - Get a certificate */
    public Map<String, Object> get(String id, Map<String, String> params) {
        return http.get("/open/cert/{id}".replace("{id}", id), params);
    }

    /** DELETE /open/cert/{id} - Delete a certificate */
    public Map<String, Object> delete(String id, Object data) {
        return http.delete("/open/cert/{id}".replace("{id}", id), data);
    }

}