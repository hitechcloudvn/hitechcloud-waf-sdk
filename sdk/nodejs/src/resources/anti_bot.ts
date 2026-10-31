import { BaseResource } from "./base";
import { Params, Body } from "../http-client";

/** Anti-Bot & Intelligence: Threat intelligence, IP groups, crawlers and behaviour/fingerprint sharing. */
export class AntiBotResource extends BaseResource {
  /** GET /open/intelligence - Get threat-intelligence config */
  getIntelligence(params?: Params): Promise<any> {
    return this.http.get("/open/intelligence", params);
  }
  /** POST /open/intelligence - Set threat-intelligence config */
  setIntelligence(data?: Body): Promise<any> {
    return this.http.post("/open/intelligence", data);
  }
  /** POST /open/intelligence/ip_lib - Upload a custom IP library */
  uploadIpLib(data?: Body): Promise<any> {
    return this.http.post("/open/intelligence/ip_lib", data);
  }
  /** GET /open/ipgroup - List IP groups */
  listIpGroups(params?: Params): Promise<any> {
    return this.http.get("/open/ipgroup", params);
  }
  /** POST /open/ipgroup - Create an IP group */
  createIpGroup(data?: Body): Promise<any> {
    return this.http.post("/open/ipgroup", data);
  }
  /** PUT /open/ipgroup - Update an IP group */
  updateIpGroup(data?: Body): Promise<any> {
    return this.http.put("/open/ipgroup", data);
  }
  /** DELETE /open/ipgroup - Delete an IP group */
  deleteIpGroup(data?: Body): Promise<any> {
    return this.http.delete("/open/ipgroup", data);
  }
  /** POST /open/ipgroup/append - Append entries to an IP group */
  appendIpGroup(data?: Body): Promise<any> {
    return this.http.post("/open/ipgroup/append", data);
  }
  /** GET /open/ipgroup/crawler - Get known-crawler groups */
  getCrawlers(params?: Params): Promise<any> {
    return this.http.get("/open/ipgroup/crawler", params);
  }
  /** POST /open/ipgroup/crawler - Configure known-crawler groups */
  setCrawlers(data?: Body): Promise<any> {
    return this.http.post("/open/ipgroup/crawler", data);
  }
  /** GET /open/ipgroup/detail - Get IP-group detail */
  getIpGroupDetail(params?: Params): Promise<any> {
    return this.http.get("/open/ipgroup/detail", params);
  }
  /** GET /open/ipgroup/link - Get IP-group link bindings */
  getIpGroupLink(params?: Params): Promise<any> {
    return this.http.get("/open/ipgroup/link", params);
  }
  /** POST /open/ipgroup/link - Set IP-group link bindings */
  setIpGroupLink(data?: Body): Promise<any> {
    return this.http.post("/open/ipgroup/link", data);
  }
  /** GET /open/ipgroup/tserver - Get trusted-server IP groups */
  getTrustServers(params?: Params): Promise<any> {
    return this.http.get("/open/ipgroup/tserver", params);
  }
  /** POST /open/ipgroup/tserver - Set trusted-server IP groups */
  setTrustServers(data?: Body): Promise<any> {
    return this.http.post("/open/ipgroup/tserver", data);
  }
  /** GET /open/share_behaviour - Get behaviour-sharing config */
  getShareBehaviour(params?: Params): Promise<any> {
    return this.http.get("/open/share_behaviour", params);
  }
  /** POST /open/share_behaviour - Set behaviour-sharing config */
  setShareBehaviour(data?: Body): Promise<any> {
    return this.http.post("/open/share_behaviour", data);
  }
  /** GET /open/share_fingerprint - Get fingerprint-sharing config */
  getShareFingerprint(params?: Params): Promise<any> {
    return this.http.get("/open/share_fingerprint", params);
  }
  /** POST /open/share_fingerprint - Set fingerprint-sharing config */
  setShareFingerprint(data?: Body): Promise<any> {
    return this.http.post("/open/share_fingerprint", data);
  }
}
