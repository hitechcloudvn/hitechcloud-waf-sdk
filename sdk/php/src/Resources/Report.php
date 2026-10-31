<?php

namespace HiTechCloud\WAF\Resources;

/**
 * Reports: Scheduled / on-demand security reports.
 */
class Report extends BaseResource
{
    /** GET /business/report - List reports */
    public function list(array $params = []): ?array
    {
        return $this->http->get('/business/report', $params);
    }

    /** POST /business/report - Create a report */
    public function create(?array $data = null): ?array
    {
        return $this->http->post('/business/report', $data);
    }

    /** GET /business/report/{id} - Get a report */
    public function get(string $id, array $params = []): ?array
    {
        return $this->http->get('/business/report/' . $id, $params);
    }

    /** DELETE /business/report/{id} - Delete a report */
    public function delete(string $id, ?array $data = null): ?array
    {
        return $this->http->delete('/business/report/' . $id, $data);
    }

}