import { BaseResource } from "./base";
import { Params, Body } from "../http-client";

/** Skynet AI Rules: AI-driven Skynet detection rules and switches. */
export class SkynetResource extends BaseResource {
  /** GET /open/skynet/rule - Get Skynet rules */
  getRules(params?: Params): Promise<any> {
    return this.http.get("/open/skynet/rule", params);
  }
  /** PUT /open/skynet/rule - Update Skynet rules */
  updateRules(data?: Body): Promise<any> {
    return this.http.put("/open/skynet/rule", data);
  }
  /** GET /open/skynet/rule/switch - Get Skynet rule switch state */
  getRuleSwitch(params?: Params): Promise<any> {
    return this.http.get("/open/skynet/rule/switch", params);
  }
  /** PUT /open/skynet/rule/switch - Toggle Skynet rules */
  setRuleSwitch(data?: Body): Promise<any> {
    return this.http.put("/open/skynet/rule/switch", data);
  }
}
