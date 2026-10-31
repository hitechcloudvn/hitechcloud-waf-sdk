import { BaseResource } from "./base";
import { Params, Body } from "../http-client";

/** Reports: Scheduled / on-demand security reports. */
export class ReportResource extends BaseResource {
  /** GET /business/report - List reports */
  list(params?: Params): Promise<any> {
    return this.http.get("/business/report", params);
  }
  /** POST /business/report - Create a report */
  create(data?: Body): Promise<any> {
    return this.http.post("/business/report", data);
  }
  /** GET /business/report/{id} - Get a report */
  get(id: string, params?: Params): Promise<any> {
    return this.http.get(`/business/report/${id}`, params);
  }
  /** DELETE /business/report/{id} - Delete a report */
  delete(id: string, data?: Body): Promise<any> {
    return this.http.delete(`/business/report/${id}`, data);
  }
}
