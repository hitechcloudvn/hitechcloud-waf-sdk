package vn.hitechcloud.waf.resources;

import java.util.Map;
import vn.hitechcloud.waf.HttpClient;

/** Attacks & Events: Attack events, matched rules, raw records and export. */
public class Events extends BaseResource {

    public Events(HttpClient http) {
        super(http);
    }

    /** GET /open/events - List attack events */
    public Map<String, Object> list(Map<String, String> params) {
        return http.get("/open/events", params);
    }

    /** GET /open/events/rule - List events grouped by rule */
    public Map<String, Object> listRules(Map<String, String> params) {
        return http.get("/open/events/rule", params);
    }

    /** GET /open/record/{id} - Get a single event record */
    public Map<String, Object> getRecord(String id, Map<String, String> params) {
        return http.get("/open/record/{id}".replace("{id}", id), params);
    }

    /** GET /open/record/rule/{id} - Get the rule for a record */
    public Map<String, Object> getRecordRule(String id, Map<String, String> params) {
        return http.get("/open/record/rule/{id}".replace("{id}", id), params);
    }

    /** GET /open/records - List raw records */
    public Map<String, Object> records(Map<String, String> params) {
        return http.get("/open/records", params);
    }

    /** GET /open/records/rule - List raw records grouped by rule */
    public Map<String, Object> recordsRule(Map<String, String> params) {
        return http.get("/open/records/rule", params);
    }

    /** GET /commercial/record/export - Export records (commercial) */
    public Map<String, Object> export(Map<String, String> params) {
        return http.get("/commercial/record/export", params);
    }

}