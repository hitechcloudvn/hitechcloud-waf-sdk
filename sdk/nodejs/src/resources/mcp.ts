import { BaseResource } from "./base";
import { Params, Body } from "../http-client";

/** MCP: Model Context Protocol bridge endpoints. */
export class McpResource extends BaseResource {
  /** GET /open/mcp - Get MCP capability metadata */
  get(params?: Params): Promise<any> {
    return this.http.get("/open/mcp", params);
  }
  /** POST /open/mcp - Invoke an MCP action */
  call(data?: Body): Promise<any> {
    return this.http.post("/open/mcp", data);
  }
}
