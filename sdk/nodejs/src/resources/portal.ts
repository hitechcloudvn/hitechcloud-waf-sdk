import { BaseResource } from "./base";
import { Params, Body } from "../http-client";

/** Portal & Branding: Self-service portal, block pages, branding and accounts. */
export class PortalResource extends BaseResource {
  /** PUT /commercial/portal/style - Update portal style (commercial) */
  updateStyle(data?: Body): Promise<any> {
    return this.http.put("/commercial/portal/style", data);
  }
  /** GET /open/portal - Get portal config */
  get(params?: Params): Promise<any> {
    return this.http.get("/open/portal", params);
  }
  /** PUT /open/portal - Update portal config */
  update(data?: Body): Promise<any> {
    return this.http.put("/open/portal", data);
  }
  /** GET /open/portal/proxy_config - Get portal proxy config */
  getProxyConfig(params?: Params): Promise<any> {
    return this.http.get("/open/portal/proxy_config", params);
  }
  /** PUT /open/portal/proxy_config - Set portal proxy config */
  setProxyConfig(data?: Body): Promise<any> {
    return this.http.put("/open/portal/proxy_config", data);
  }
  /** GET /open/portal/style - Get portal style */
  getStyle(params?: Params): Promise<any> {
    return this.http.get("/open/portal/style", params);
  }
  /** PUT /commercial/block_page - Update the block page (commercial) */
  updateBlockPage(data?: Body): Promise<any> {
    return this.http.put("/commercial/block_page", data);
  }
  /** GET /open/block_page - Get the active block page */
  getBlockPage(params?: Params): Promise<any> {
    return this.http.get("/open/block_page", params);
  }
  /** GET /open/block_page_list - List available block pages */
  listBlockPages(params?: Params): Promise<any> {
    return this.http.get("/open/block_page_list", params);
  }
  /** GET /business/frontend_style - Get frontend style */
  getFrontendStyle(params?: Params): Promise<any> {
    return this.http.get("/business/frontend_style", params);
  }
  /** PUT /business/frontend_style - Set frontend style */
  setFrontendStyle(data?: Body): Promise<any> {
    return this.http.put("/business/frontend_style", data);
  }
  /** GET /business/account - Get portal account(s) */
  getAccount(params?: Params): Promise<any> {
    return this.http.get("/business/account", params);
  }
  /** POST /business/account - Create a portal account */
  createAccount(data?: Body): Promise<any> {
    return this.http.post("/business/account", data);
  }
}
