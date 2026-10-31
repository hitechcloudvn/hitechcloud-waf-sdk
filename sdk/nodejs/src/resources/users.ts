import { BaseResource } from "./base";
import { Params, Body } from "../http-client";

/** Users & Account: Manage administrative users, TOTP and the current manager profile. */
export class UsersResource extends BaseResource {
  /** GET /open/users - List users */
  list(params?: Params): Promise<any> {
    return this.http.get("/open/users", params);
  }
  /** POST /open/users - Create a user */
  create(data?: Body): Promise<any> {
    return this.http.post("/open/users", data);
  }
  /** PUT /open/users - Update a user */
  update(data?: Body): Promise<any> {
    return this.http.put("/open/users", data);
  }
  /** DELETE /open/users - Delete a user */
  delete(data?: Body): Promise<any> {
    return this.http.delete("/open/users", data);
  }
  /** POST /open/users/{id}/totp - Enable / bind TOTP for a user */
  setupTotp(id: string, data?: Body): Promise<any> {
    return this.http.post(`/open/users/${id}/totp`, data);
  }
  /** GET /ManagerInfo - Get the current manager (admin) profile */
  getManagerInfo(params?: Params): Promise<any> {
    return this.http.get("/ManagerInfo", params);
  }
  /** PUT /ManagerInfo - Update the current manager (admin) profile */
  updateManagerInfo(data?: Body): Promise<any> {
    return this.http.put("/ManagerInfo", data);
  }
}
