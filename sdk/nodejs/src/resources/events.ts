import { BaseResource } from "./base";
import { Params, Body } from "../http-client";

/** Attacks & Events: Attack events, matched rules, raw records and export. */
export class EventsResource extends BaseResource {
  /** GET /open/events - List attack events */
  list(params?: Params): Promise<any> {
    return this.http.get("/open/events", params);
  }
  /** GET /open/events/rule - List events grouped by rule */
  listRules(params?: Params): Promise<any> {
    return this.http.get("/open/events/rule", params);
  }
  /** GET /open/record/{id} - Get a single event record */
  getRecord(id: string, params?: Params): Promise<any> {
    return this.http.get(`/open/record/${id}`, params);
  }
  /** GET /open/record/rule/{id} - Get the rule for a record */
  getRecordRule(id: string, params?: Params): Promise<any> {
    return this.http.get(`/open/record/rule/${id}`, params);
  }
  /** GET /open/records - List raw records */
  records(params?: Params): Promise<any> {
    return this.http.get("/open/records", params);
  }
  /** GET /open/records/rule - List raw records grouped by rule */
  recordsRule(params?: Params): Promise<any> {
    return this.http.get("/open/records/rule", params);
  }
  /** GET /commercial/record/export - Export records (commercial) */
  export(params?: Params): Promise<any> {
    return this.http.get("/commercial/record/export", params);
  }
}
