import { BaseResource } from "./base";
import { Params, Body } from "../http-client";

/** Cloud Policies: Shared cloud threat-intelligence policies. */
export class CloudResource extends BaseResource {
  /** GET /open/cloud/policies - List available cloud policies */
  listPolicies(params?: Params): Promise<any> {
    return this.http.get("/open/cloud/policies", params);
  }
  /** POST /open/cloud/policies/subscribe - Subscribe to a cloud policy */
  subscribe(data?: Body): Promise<any> {
    return this.http.post("/open/cloud/policies/subscribe", data);
  }
}
