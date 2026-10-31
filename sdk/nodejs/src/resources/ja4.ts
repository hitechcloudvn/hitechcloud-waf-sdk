import { BaseResource } from "./base";
import { Params, Body } from "../http-client";

/** JA4 Fingerprint: JA4/JA4+ TLS client-fingerprint rules. */
export class Ja4Resource extends BaseResource {
  /** GET /open/ja4 - Get JA4 fingerprint rules */
  get(params?: Params): Promise<any> {
    return this.http.get("/open/ja4", params);
  }
  /** PUT /open/ja4 - Update JA4 fingerprint rules */
  update(data?: Body): Promise<any> {
    return this.http.put("/open/ja4", data);
  }
}
