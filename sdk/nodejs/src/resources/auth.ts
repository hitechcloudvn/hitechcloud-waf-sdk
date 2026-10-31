import { BaseResource } from "./base";
import { Params, Body } from "../http-client";

/** Authentication & Token: CSRF, login, TFA, logout and API token lifecycle. */
export class AuthResource extends BaseResource {
  /** GET /open/auth/csrf - Get a CSRF token for the login step */
  csrf(params?: Params): Promise<any> {
    return this.http.get("/open/auth/csrf", params);
  }
  /** POST /open/auth/login - Authenticate with username/password (+csrf_token) and obtain a JWT */
  login(data?: Body): Promise<any> {
    return this.http.post("/open/auth/login", data);
  }
  /** POST /open/auth/logout - End the current session */
  logout(data?: Body): Promise<any> {
    return this.http.post("/open/auth/logout", data);
  }
  /** POST /open/auth/tfa - Verify a two-factor (OTP) code during login */
  tfa(data?: Body): Promise<any> {
    return this.http.post("/open/auth/tfa", data);
  }
  /** GET /open/auth/token - Get the current API token */
  getToken(params?: Params): Promise<any> {
    return this.http.get("/open/auth/token", params);
  }
  /** PUT /open/auth/token - Rotate / refresh the API token */
  refreshToken(data?: Body): Promise<any> {
    return this.http.put("/open/auth/token", data);
  }
  /** DELETE /open/auth/token - Revoke the API token */
  revokeToken(data?: Body): Promise<any> {
    return this.http.delete("/open/auth/token", data);
  }
}
