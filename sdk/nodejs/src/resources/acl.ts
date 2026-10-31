import { BaseResource } from "./base";
import { Params, Body } from "../http-client";

/** ACL & Rate-Limit: Access-control / rate-limit rules and block relief. */
export class AclResource extends BaseResource {
  /** PUT /acl/relieve - Relieve (unblock) an ACL-blocked client */
  relieve(data?: Body): Promise<any> {
    return this.http.put("/acl/relieve", data);
  }
  /** GET /open/records/acl - ACL hit records */
  records(params?: Params): Promise<any> {
    return this.http.get("/open/records/acl", params);
  }
  /** GET /open/site/{id}/acl - List ACL / rate-limit rules for a site */
  list(id: string, params?: Params): Promise<any> {
    return this.http.get(`/open/site/${id}/acl`, params);
  }
  /** POST /open/site/{id}/acl - Create an ACL rule */
  create(id: string, data?: Body): Promise<any> {
    return this.http.post(`/open/site/${id}/acl`, data);
  }
  /** PUT /open/site/{id}/acl - Update ACL rules for a site */
  update(id: string, data?: Body): Promise<any> {
    return this.http.put(`/open/site/${id}/acl`, data);
  }
  /** PUT /open/site/{id}/acl/{rule_id} - Update a single ACL rule */
  updateRule(id: string, ruleId: string, data?: Body): Promise<any> {
    return this.http.put(`/open/site/${id}/acl/${ruleId}`, data);
  }
  /** DELETE /open/site/{id}/acl/{rule_id} - Delete a single ACL rule */
  deleteRule(id: string, ruleId: string, data?: Body): Promise<any> {
    return this.http.delete(`/open/site/${id}/acl/${ruleId}`, data);
  }
}
