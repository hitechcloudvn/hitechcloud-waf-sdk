package vn.hitechcloud.waf.resources;

import java.util.Map;
import vn.hitechcloud.waf.HttpClient;

/** JA4 Fingerprint: JA4/JA4+ TLS client-fingerprint rules. */
public class Ja4 extends BaseResource {

    public Ja4(HttpClient http) {
        super(http);
    }

    /** GET /open/ja4 - Get JA4 fingerprint rules */
    public Map<String, Object> get(Map<String, String> params) {
        return http.get("/open/ja4", params);
    }

    /** PUT /open/ja4 - Update JA4 fingerprint rules */
    public Map<String, Object> update(Object data) {
        return http.put("/open/ja4", data);
    }

}