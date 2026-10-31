import { BaseResource } from "./base";
import { Params, Body } from "../http-client";

/** Global Mode & Proxy: Global protection mode, outbound proxy and log cleanup. */
export class GlobalSettingsResource extends BaseResource {
  /** GET /open/global/mode - Get the global protection mode */
  getMode(params?: Params): Promise<any> {
    return this.http.get("/open/global/mode", params);
  }
  /** PUT /open/global/mode - Set the global protection mode */
  setMode(data?: Body): Promise<any> {
    return this.http.put("/open/global/mode", data);
  }
  /** GET /open/global/proxy - Get the global proxy config */
  getProxy(params?: Params): Promise<any> {
    return this.http.get("/open/global/proxy", params);
  }
  /** PUT /open/global/proxy - Set the global proxy config */
  setProxy(data?: Body): Promise<any> {
    return this.http.put("/open/global/proxy", data);
  }
  /** GET /open/global/log_clean - Get the log-cleanup policy */
  getLogClean(params?: Params): Promise<any> {
    return this.http.get("/open/global/log_clean", params);
  }
  /** POST /open/global/log_clean - Set the log-cleanup policy */
  setLogClean(data?: Body): Promise<any> {
    return this.http.post("/open/global/log_clean", data);
  }
}
