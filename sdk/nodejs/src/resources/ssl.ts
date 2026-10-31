import { BaseResource } from "./base";
import { Params, Body } from "../http-client";

/** SSL Certificates: TLS/SSL certificate management. */
export class SslResource extends BaseResource {
  /** GET /open/cert - List certificates */
  list(params?: Params): Promise<any> {
    return this.http.get("/open/cert", params);
  }
  /** POST /open/cert - Upload / create a certificate */
  create(data?: Body): Promise<any> {
    return this.http.post("/open/cert", data);
  }
  /** GET /open/cert/{id} - Get a certificate */
  get(id: string, params?: Params): Promise<any> {
    return this.http.get(`/open/cert/${id}`, params);
  }
  /** DELETE /open/cert/{id} - Delete a certificate */
  delete(id: string, data?: Body): Promise<any> {
    return this.http.delete(`/open/cert/${id}`, data);
  }
}
