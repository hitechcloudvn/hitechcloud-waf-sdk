package vn.hitechcloud.waf.resources;

import java.util.Map;
import vn.hitechcloud.waf.HttpClient;

/** MCP: Model Context Protocol bridge endpoints. */
public class Mcp extends BaseResource {

    public Mcp(HttpClient http) {
        super(http);
    }

    /** GET /open/mcp - Get MCP capability metadata */
    public Map<String, Object> get(Map<String, String> params) {
        return http.get("/open/mcp", params);
    }

    /** POST /open/mcp - Invoke an MCP action */
    public Map<String, Object> call(Object data) {
        return http.post("/open/mcp", data);
    }

}