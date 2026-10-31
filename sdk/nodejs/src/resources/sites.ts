import { BaseResource } from "./base";
import { Params, Body } from "../http-client";

/** Sites (Web Services): Protected sites, groups, proxy/origin config, logs and forwarding rules. */
export class SitesResource extends BaseResource {
  /** GET /open/site - List protected sites */
  list(params?: Params): Promise<any> {
    return this.http.get("/open/site", params);
  }
  /** POST /open/site - Create a protected site */
  create(data?: Body): Promise<any> {
    return this.http.post("/open/site", data);
  }
  /** PUT /open/site - Update a site */
  update(data?: Body): Promise<any> {
    return this.http.put("/open/site", data);
  }
  /** DELETE /open/site - Delete a site */
  delete(data?: Body): Promise<any> {
    return this.http.delete("/open/site", data);
  }
  /** PUT /open/site/challenge - Toggle challenge mode for sites */
  setChallenge(data?: Body): Promise<any> {
    return this.http.put("/open/site/challenge", data);
  }
  /** PUT /open/site/defense - Toggle defense mode for sites */
  setDefense(data?: Body): Promise<any> {
    return this.http.put("/open/site/defense", data);
  }
  /** GET /open/site/group - List site groups */
  listGroups(params?: Params): Promise<any> {
    return this.http.get("/open/site/group", params);
  }
  /** POST /open/site/group - Create a site group */
  createGroup(data?: Body): Promise<any> {
    return this.http.post("/open/site/group", data);
  }
  /** GET /open/site/group/switch - Get site-group switch state */
  getGroupSwitch(params?: Params): Promise<any> {
    return this.http.get("/open/site/group/switch", params);
  }
  /** PUT /open/site/group/{id} - Update a site group */
  updateGroup(id: string, data?: Body): Promise<any> {
    return this.http.put(`/open/site/group/${id}`, data);
  }
  /** DELETE /open/site/group/{id} - Delete a site group */
  deleteGroup(id: string, data?: Body): Promise<any> {
    return this.http.delete(`/open/site/group/${id}`, data);
  }
  /** PUT /open/site/group/{id}/sort - Sort a site group */
  sortGroup(id: string, data?: Body): Promise<any> {
    return this.http.put(`/open/site/group/${id}/sort`, data);
  }
  /** PUT /commercial/site/group/switch - Switch the active site group (commercial) */
  switchGroup(data?: Body): Promise<any> {
    return this.http.put("/commercial/site/group/switch", data);
  }
  /** POST /open/site/healthcheck - Run an origin health check */
  healthcheck(data?: Body): Promise<any> {
    return this.http.post("/open/site/healthcheck", data);
  }
  /** PUT /open/site/mode - Set the protection mode for a site */
  setMode(data?: Body): Promise<any> {
    return this.http.put("/open/site/mode", data);
  }
  /** GET /open/site/{id} - Get a site */
  get(id: string, params?: Params): Promise<any> {
    return this.http.get(`/open/site/${id}`, params);
  }
  /** PUT /open/site/{id}/basic_info - Update site basic info */
  updateBasicInfo(id: string, data?: Body): Promise<any> {
    return this.http.put(`/open/site/${id}/basic_info`, data);
  }
  /** GET /open/site/{id}/excludes - List exclusion rules */
  getExcludes(id: string, params?: Params): Promise<any> {
    return this.http.get(`/open/site/${id}/excludes`, params);
  }
  /** POST /open/site/{id}/excludes - Add an exclusion rule */
  addExclude(id: string, data?: Body): Promise<any> {
    return this.http.post(`/open/site/${id}/excludes`, data);
  }
  /** PUT /open/site/{id}/group - Move a site into a group */
  setGroup(id: string, data?: Body): Promise<any> {
    return this.http.put(`/open/site/${id}/group`, data);
  }
  /** GET /open/site/{id}/log - Get site logs */
  getLog(id: string, params?: Params): Promise<any> {
    return this.http.get(`/open/site/${id}/log`, params);
  }
  /** GET /open/site/{id}/log/detail - Get a log detail */
  getLogDetail(id: string, params?: Params): Promise<any> {
    return this.http.get(`/open/site/${id}/log/detail`, params);
  }
  /** GET /open/site/{id}/log/download - Download site logs */
  downloadLog(id: string, params?: Params): Promise<any> {
    return this.http.get(`/open/site/${id}/log/download`, params);
  }
  /** GET /open/site/{id}/log/limit - Get log limit config */
  getLogLimit(id: string, params?: Params): Promise<any> {
    return this.http.get(`/open/site/${id}/log/limit`, params);
  }
  /** POST /open/site/{id}/log/limit - Set log limit config */
  setLogLimit(id: string, data?: Body): Promise<any> {
    return this.http.post(`/open/site/${id}/log/limit`, data);
  }
  /** GET /open/site/{id}/nginx_config - Get generated nginx config */
  getNginxConfig(id: string, params?: Params): Promise<any> {
    return this.http.get(`/open/site/${id}/nginx_config`, params);
  }
  /** PUT /open/site/{id}/nginx_config - Update custom nginx config */
  updateNginxConfig(id: string, data?: Body): Promise<any> {
    return this.http.put(`/open/site/${id}/nginx_config`, data);
  }
  /** GET /open/site/{id}/proxy - Get reverse-proxy / origin config */
  getProxy(id: string, params?: Params): Promise<any> {
    return this.http.get(`/open/site/${id}/proxy`, params);
  }
  /** PUT /open/site/{id}/proxy - Update reverse-proxy / origin config */
  updateProxy(id: string, data?: Body): Promise<any> {
    return this.http.put(`/open/site/${id}/proxy`, data);
  }
  /** GET /open/site/{id}/resources - List static resources */
  getResources(id: string, params?: Params): Promise<any> {
    return this.http.get(`/open/site/${id}/resources`, params);
  }
  /** DELETE /open/site/{id}/resources - Delete static resources */
  deleteResources(id: string, data?: Body): Promise<any> {
    return this.http.delete(`/open/site/${id}/resources`, data);
  }
  /** PUT /open/site/{id}/sort - Sort a site */
  sort(id: string, data?: Body): Promise<any> {
    return this.http.put(`/open/site/${id}/sort`, data);
  }
  /** GET /open/site/{id}/static - List static-site rules */
  getStatic(id: string, params?: Params): Promise<any> {
    return this.http.get(`/open/site/${id}/static`, params);
  }
  /** POST /open/site/{id}/static - Add a static-site rule */
  addStatic(id: string, data?: Body): Promise<any> {
    return this.http.post(`/open/site/${id}/static`, data);
  }
  /** PUT /open/site/{id}/static - Update a static-site rule */
  updateStatic(id: string, data?: Body): Promise<any> {
    return this.http.put(`/open/site/${id}/static`, data);
  }
  /** DELETE /open/site/{id}/static - Delete a static-site rule */
  deleteStatic(id: string, data?: Body): Promise<any> {
    return this.http.delete(`/open/site/${id}/static`, data);
  }
  /** GET /open/site/{id}/forwarding_rules - List forwarding rules */
  listForwardingRules(id: string, params?: Params): Promise<any> {
    return this.http.get(`/open/site/${id}/forwarding_rules`, params);
  }
  /** POST /open/site/{id}/forwarding_rules - Create a forwarding rule */
  createForwardingRule(id: string, data?: Body): Promise<any> {
    return this.http.post(`/open/site/${id}/forwarding_rules`, data);
  }
  /** PUT /open/site/{id}/forwarding_rules/{rule_id} - Update a forwarding rule */
  updateForwardingRule(id: string, ruleId: string, data?: Body): Promise<any> {
    return this.http.put(`/open/site/${id}/forwarding_rules/${ruleId}`, data);
  }
  /** DELETE /open/site/{id}/forwarding_rules/{rule_id} - Delete a forwarding rule */
  deleteForwardingRule(id: string, ruleId: string, data?: Body): Promise<any> {
    return this.http.delete(`/open/site/${id}/forwarding_rules/${ruleId}`, data);
  }
  /** PUT /open/site/{id}/forwarding_rules/{rule_id}/switch - Enable / disable a forwarding rule */
  switchForwardingRule(id: string, ruleId: string, data?: Body): Promise<any> {
    return this.http.put(`/open/site/${id}/forwarding_rules/${ruleId}/switch`, data);
  }
}
