import { BaseResource } from "./base";
import { Params, Body } from "../http-client";

/** Alarms: Alert / notification configuration. */
export class AlarmResource extends BaseResource {
  /** GET /commercial/alarm - Get alarm config */
  get(params?: Params): Promise<any> {
    return this.http.get("/commercial/alarm", params);
  }
  /** PUT /commercial/alarm - Update alarm config */
  update(data?: Body): Promise<any> {
    return this.http.put("/commercial/alarm", data);
  }
  /** POST /commercial/alarm/test - Send a test alarm */
  test(data?: Body): Promise<any> {
    return this.http.post("/commercial/alarm/test", data);
  }
}
