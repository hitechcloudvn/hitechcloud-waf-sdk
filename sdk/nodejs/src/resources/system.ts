import { BaseResource } from "./base";
import { Params, Body } from "../http-client";

/** System: System info, edition, protocols, network proxy and helpers. */
export class SystemResource extends BaseResource {
  /** GET /open/system - Get system settings */
  get(params?: Params): Promise<any> {
    return this.http.get("/open/system", params);
  }
  /** PUT /open/system - Update system settings */
  update(data?: Body): Promise<any> {
    return this.http.put("/open/system", data);
  }
  /** GET /open/system/arch - Get the cluster architecture */
  getArch(params?: Params): Promise<any> {
    return this.http.get("/open/system/arch", params);
  }
  /** GET /open/system/customer_service - Get customer-service config */
  getCustomerService(params?: Params): Promise<any> {
    return this.http.get("/open/system/customer_service", params);
  }
  /** PUT /open/system/customer_service - Set customer-service config */
  setCustomerService(data?: Body): Promise<any> {
    return this.http.put("/open/system/customer_service", data);
  }
  /** GET /open/system/edition - Get the product edition */
  getEdition(params?: Params): Promise<any> {
    return this.http.get("/open/system/edition", params);
  }
  /** GET /open/system/key - Get the system public key */
  getKey(params?: Params): Promise<any> {
    return this.http.get("/open/system/key", params);
  }
  /** GET /open/system/login_method - Get the system login method */
  getLoginMethod(params?: Params): Promise<any> {
    return this.http.get("/open/system/login_method", params);
  }
  /** GET /open/system/login_type - Get the system login type */
  getLoginType(params?: Params): Promise<any> {
    return this.http.get("/open/system/login_type", params);
  }
  /** GET /open/system/network_proxy - Get the network proxy */
  getNetworkProxy(params?: Params): Promise<any> {
    return this.http.get("/open/system/network_proxy", params);
  }
  /** POST /open/system/network_proxy - Set the network proxy */
  setNetworkProxy(data?: Body): Promise<any> {
    return this.http.post("/open/system/network_proxy", data);
  }
  /** GET /open/system/protocol - Get supported protocols */
  getProtocol(params?: Params): Promise<any> {
    return this.http.get("/open/system/protocol", params);
  }
  /** POST /open/system/protocol - Configure supported protocols */
  setProtocol(data?: Body): Promise<any> {
    return this.http.post("/open/system/protocol", data);
  }
  /** POST /commercial/dashboard/token - Issue a dashboard token */
  dashboardToken(data?: Body): Promise<any> {
    return this.http.post("/commercial/dashboard/token", data);
  }
  /** POST /open/safepoint/consult - Consult the safepoint advisor */
  safepointConsult(data?: Body): Promise<any> {
    return this.http.post("/open/safepoint/consult", data);
  }
}
