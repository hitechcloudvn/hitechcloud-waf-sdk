import { BaseResource } from "./base";
import { Params, Body } from "../http-client";

/** Security Posture: Realtime posture, per-site scoring, statistics and trends. */
export class SecurityPostureResource extends BaseResource {
  /** GET /open/security_posture/realtime - Realtime security posture */
  realtime(params?: Params): Promise<any> {
    return this.http.get("/open/security_posture/realtime", params);
  }
  /** POST /open/security_posture/site - Security posture for selected sites */
  site(data?: Body): Promise<any> {
    return this.http.post("/open/security_posture/site", data);
  }
  /** GET /open/security_posture/statistics - Security-posture statistics */
  statistics(params?: Params): Promise<any> {
    return this.http.get("/open/security_posture/statistics", params);
  }
  /** GET /open/security_posture/trends - Security-posture trends */
  trends(params?: Params): Promise<any> {
    return this.http.get("/open/security_posture/trends", params);
  }
}
