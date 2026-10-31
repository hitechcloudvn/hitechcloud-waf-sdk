import { BaseResource } from "./base";
import { Params, Body } from "../http-client";

/** Settings: Login methods, syslog, detector, downgrade and audit log. */
export class SettingsResource extends BaseResource {
  /** PUT /business/system/login_method - Set the admin login method */
  setLoginMethod(data?: Body): Promise<any> {
    return this.http.put("/business/system/login_method", data);
  }
  /** POST /business/system/login_method/test - Test the admin login method */
  testLoginMethod(data?: Body): Promise<any> {
    return this.http.post("/business/system/login_method/test", data);
  }
  /** GET /commercial/syslog - Get syslog config */
  getSyslog(params?: Params): Promise<any> {
    return this.http.get("/commercial/syslog", params);
  }
  /** POST /commercial/syslog - Set syslog config */
  setSyslog(data?: Body): Promise<any> {
    return this.http.post("/commercial/syslog", data);
  }
  /** POST /commercial/syslog/test - Test syslog delivery */
  testSyslog(data?: Body): Promise<any> {
    return this.http.post("/commercial/syslog/test", data);
  }
  /** GET /open/Commercial/downgrade - Get downgrade config */
  getDowngrade(params?: Params): Promise<any> {
    return this.http.get("/open/Commercial/downgrade", params);
  }
  /** PUT /open/Commercial/downgrade - Set downgrade config */
  setDowngrade(data?: Body): Promise<any> {
    return this.http.put("/open/Commercial/downgrade", data);
  }
  /** GET /open/audit_log - Get the admin audit log */
  auditLog(params?: Params): Promise<any> {
    return this.http.get("/open/audit_log", params);
  }
  /** GET /open/detector - Get detector config */
  getDetector(params?: Params): Promise<any> {
    return this.http.get("/open/detector", params);
  }
  /** POST /open/detector - Set detector config */
  setDetector(data?: Body): Promise<any> {
    return this.http.post("/open/detector", data);
  }
}
