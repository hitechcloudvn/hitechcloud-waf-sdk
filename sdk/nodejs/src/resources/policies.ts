import { BaseResource } from "./base";
import { Params, Body } from "../http-client";

/** Policies & Rules: Custom protection policies and rule ordering. */
export class PoliciesResource extends BaseResource {
  /** GET /open/policy - List policies */
  list(params?: Params): Promise<any> {
    return this.http.get("/open/policy", params);
  }
  /** POST /open/policy - Create a policy */
  create(data?: Body): Promise<any> {
    return this.http.post("/open/policy", data);
  }
  /** PUT /open/policy - Update a policy */
  update(data?: Body): Promise<any> {
    return this.http.put("/open/policy", data);
  }
  /** DELETE /open/policy - Delete a policy */
  delete(data?: Body): Promise<any> {
    return this.http.delete("/open/policy", data);
  }
  /** GET /open/policy/detail - Get policy detail */
  detail(params?: Params): Promise<any> {
    return this.http.get("/open/policy/detail", params);
  }
  /** PUT /open/policy/order - Set a policy order */
  setOrder(data?: Body): Promise<any> {
    return this.http.put("/open/policy/order", data);
  }
  /** PUT /open/policy/order/value - Set a policy order value */
  setOrderValue(data?: Body): Promise<any> {
    return this.http.put("/open/policy/order/value", data);
  }
  /** PUT /open/policy/orders - Bulk-set policy orders */
  setOrders(data?: Body): Promise<any> {
    return this.http.put("/open/policy/orders", data);
  }
  /** PUT /open/policy/switch - Toggle a policy on/off */
  setSwitch(data?: Body): Promise<any> {
    return this.http.put("/open/policy/switch", data);
  }
}
