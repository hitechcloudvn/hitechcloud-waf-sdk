<?php

namespace HiTechCloud\WAF\Resources;

/**
 * Statistics: Advanced access/attack analytics and QPS.
 */
class Statistics extends BaseResource
{
    /** GET /stat/advance/access - Access statistics */
    public function access(array $params = []): ?array
    {
        return $this->http->get('/stat/advance/access', $params);
    }

    /** GET /stat/advance/attack - Attack statistics */
    public function attack(array $params = []): ?array
    {
        return $this->http->get('/stat/advance/attack', $params);
    }

    /** GET /stat/advance/client - Client statistics */
    public function client(array $params = []): ?array
    {
        return $this->http->get('/stat/advance/client', $params);
    }

    /** GET /stat/advance/domain - Per-domain statistics */
    public function domain(array $params = []): ?array
    {
        return $this->http->get('/stat/advance/domain', $params);
    }

    /** GET /stat/advance/error_status_code - Error status-code statistics */
    public function errorStatusCode(array $params = []): ?array
    {
        return $this->http->get('/stat/advance/error_status_code', $params);
    }

    /** GET /stat/advance/location - Geo-location statistics */
    public function location(array $params = []): ?array
    {
        return $this->http->get('/stat/advance/location', $params);
    }

    /** GET /stat/advance/page - Per-page statistics */
    public function page(array $params = []): ?array
    {
        return $this->http->get('/stat/advance/page', $params);
    }

    /** GET /stat/advance/status_code - Status-code statistics */
    public function statusCode(array $params = []): ?array
    {
        return $this->http->get('/stat/advance/status_code', $params);
    }

    /** GET /stat/advance/trend/access - Access trend */
    public function trendAccess(array $params = []): ?array
    {
        return $this->http->get('/stat/advance/trend/access', $params);
    }

    /** GET /stat/advance/trend/intercept - Interception trend */
    public function trendIntercept(array $params = []): ?array
    {
        return $this->http->get('/stat/advance/trend/intercept', $params);
    }

    /** GET /stat/qps - Current QPS */
    public function qps(array $params = []): ?array
    {
        return $this->http->get('/stat/qps', $params);
    }

}