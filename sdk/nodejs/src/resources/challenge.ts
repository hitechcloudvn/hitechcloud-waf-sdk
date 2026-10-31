import { BaseResource } from "./base";
import { Params, Body } from "../http-client";

/** Challenge & Semantics: Human-verification challenge, semantic analysis and chaos protection. */
export class ChallengeResource extends BaseResource {
  /** GET /open/challenge/config - Get challenge config */
  getConfig(params?: Params): Promise<any> {
    return this.http.get("/open/challenge/config", params);
  }
  /** POST /open/challenge/config - Set challenge config */
  setConfig(data?: Body): Promise<any> {
    return this.http.post("/open/challenge/config", data);
  }
  /** GET /open/records/challenge - Challenge event records */
  records(params?: Params): Promise<any> {
    return this.http.get("/open/records/challenge", params);
  }
  /** GET /open/site/{id}/semantics - Get semantic-analysis config for a site */
  getSemantics(id: string, params?: Params): Promise<any> {
    return this.http.get(`/open/site/${id}/semantics`, params);
  }
  /** PUT /open/site/{id}/semantics - Set semantic-analysis config for a site */
  setSemantics(id: string, data?: Body): Promise<any> {
    return this.http.put(`/open/site/${id}/semantics`, data);
  }
  /** GET /open/site/{id}/chaos - Get chaos/dynamic-protection config for a site */
  getChaos(id: string, params?: Params): Promise<any> {
    return this.http.get(`/open/site/${id}/chaos`, params);
  }
  /** POST /open/site/{id}/chaos - Set chaos/dynamic-protection config for a site */
  setChaos(id: string, data?: Body): Promise<any> {
    return this.http.post(`/open/site/${id}/chaos`, data);
  }
}
