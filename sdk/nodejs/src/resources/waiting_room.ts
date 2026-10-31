import { BaseResource } from "./base";
import { Params, Body } from "../http-client";

/** Waiting Room: Virtual waiting-room queue for traffic peaks. */
export class WaitingRoomResource extends BaseResource {
  /** GET /open/records/waiting - Waiting-room event records */
  records(params?: Params): Promise<any> {
    return this.http.get("/open/records/waiting", params);
  }
  /** GET /open/site/{id}/waiting - Get waiting-room config for a site */
  get(id: string, params?: Params): Promise<any> {
    return this.http.get(`/open/site/${id}/waiting`, params);
  }
  /** POST /open/site/{id}/waiting - Set waiting-room config for a site */
  set(id: string, data?: Body): Promise<any> {
    return this.http.post(`/open/site/${id}/waiting`, data);
  }
}
