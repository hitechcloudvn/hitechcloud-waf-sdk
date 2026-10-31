import { BaseResource } from "./base";
import { Params, Body } from "../http-client";

/** License & Authorization: Authorization records and license status / lifecycle. */
export class LicenseResource extends BaseResource {
  /** GET /open/system/authorize - Get authorization records */
  getAuthorize(params?: Params): Promise<any> {
    return this.http.get("/open/system/authorize", params);
  }
  /** POST /open/system/authorize - Create an authorization */
  createAuthorize(data?: Body): Promise<any> {
    return this.http.post("/open/system/authorize", data);
  }
  /** PUT /open/system/authorize - Update an authorization */
  updateAuthorize(data?: Body): Promise<any> {
    return this.http.put("/open/system/authorize", data);
  }
  /** DELETE /open/system/authorize - Delete an authorization */
  deleteAuthorize(data?: Body): Promise<any> {
    return this.http.delete("/open/system/authorize", data);
  }
  /** POST /open/system/license/expired/confirm-downgrade - Confirm downgrade after license expiry */
  confirmDowngrade(data?: Body): Promise<any> {
    return this.http.post("/open/system/license/expired/confirm-downgrade", data);
  }
  /** POST /open/system/license/expiry/acknowledge - Acknowledge an upcoming license expiry */
  acknowledgeExpiry(data?: Body): Promise<any> {
    return this.http.post("/open/system/license/expiry/acknowledge", data);
  }
  /** GET /open/system/license/status - Get the license status */
  status(params?: Params): Promise<any> {
    return this.http.get("/open/system/license/status", params);
  }
}
