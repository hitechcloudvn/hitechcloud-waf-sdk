<?php

namespace HiTechCloud\WAF\Resources;

/**
 * Attacks & Events: Attack events, matched rules, raw records and export.
 */
class Events extends BaseResource
{
    /** GET /open/events - List attack events */
    public function list(array $params = []): ?array
    {
        return $this->http->get('/open/events', $params);
    }

    /** GET /open/events/rule - List events grouped by rule */
    public function listRules(array $params = []): ?array
    {
        return $this->http->get('/open/events/rule', $params);
    }

    /** GET /open/record/{id} - Get a single event record */
    public function getRecord(string $id, array $params = []): ?array
    {
        return $this->http->get('/open/record/' . $id, $params);
    }

    /** GET /open/record/rule/{id} - Get the rule for a record */
    public function getRecordRule(string $id, array $params = []): ?array
    {
        return $this->http->get('/open/record/rule/' . $id, $params);
    }

    /** GET /open/records - List raw records */
    public function records(array $params = []): ?array
    {
        return $this->http->get('/open/records', $params);
    }

    /** GET /open/records/rule - List raw records grouped by rule */
    public function recordsRule(array $params = []): ?array
    {
        return $this->http->get('/open/records/rule', $params);
    }

    /** GET /commercial/record/export - Export records (commercial) */
    public function export(array $params = []): ?array
    {
        return $this->http->get('/commercial/record/export', $params);
    }

}