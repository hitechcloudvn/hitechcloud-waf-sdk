import { BaseResource } from "./base";
import { Params, Body } from "../http-client";

/** Statistics: Advanced access/attack analytics and QPS. */
export class StatisticsResource extends BaseResource {
  /** GET /stat/advance/access - Access statistics */
  access(params?: Params): Promise<any> {
    return this.http.get("/stat/advance/access", params);
  }
  /** GET /stat/advance/attack - Attack statistics */
  attack(params?: Params): Promise<any> {
    return this.http.get("/stat/advance/attack", params);
  }
  /** GET /stat/advance/client - Client statistics */
  client(params?: Params): Promise<any> {
    return this.http.get("/stat/advance/client", params);
  }
  /** GET /stat/advance/domain - Per-domain statistics */
  domain(params?: Params): Promise<any> {
    return this.http.get("/stat/advance/domain", params);
  }
  /** GET /stat/advance/error_status_code - Error status-code statistics */
  errorStatusCode(params?: Params): Promise<any> {
    return this.http.get("/stat/advance/error_status_code", params);
  }
  /** GET /stat/advance/location - Geo-location statistics */
  location(params?: Params): Promise<any> {
    return this.http.get("/stat/advance/location", params);
  }
  /** GET /stat/advance/page - Per-page statistics */
  page(params?: Params): Promise<any> {
    return this.http.get("/stat/advance/page", params);
  }
  /** GET /stat/advance/status_code - Status-code statistics */
  statusCode(params?: Params): Promise<any> {
    return this.http.get("/stat/advance/status_code", params);
  }
  /** GET /stat/advance/trend/access - Access trend */
  trendAccess(params?: Params): Promise<any> {
    return this.http.get("/stat/advance/trend/access", params);
  }
  /** GET /stat/advance/trend/intercept - Interception trend */
  trendIntercept(params?: Params): Promise<any> {
    return this.http.get("/stat/advance/trend/intercept", params);
  }
  /** GET /stat/qps - Current QPS */
  qps(params?: Params): Promise<any> {
    return this.http.get("/stat/qps", params);
  }
}
