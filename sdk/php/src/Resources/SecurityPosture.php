<?php

namespace HiTechCloud\WAF\Resources;

/**
 * Security Posture: Realtime posture, per-site scoring, statistics and trends.
 */
class SecurityPosture extends BaseResource
{
    /** GET /open/security_posture/realtime - Realtime security posture */
    public function realtime(array $params = []): ?array
    {
        return $this->http->get('/open/security_posture/realtime', $params);
    }

    /** POST /open/security_posture/site - Security posture for selected sites */
    public function site(?array $data = null): ?array
    {
        return $this->http->post('/open/security_posture/site', $data);
    }

    /** GET /open/security_posture/statistics - Security-posture statistics */
    public function statistics(array $params = []): ?array
    {
        return $this->http->get('/open/security_posture/statistics', $params);
    }

    /** GET /open/security_posture/trends - Security-posture trends */
    public function trends(array $params = []): ?array
    {
        return $this->http->get('/open/security_posture/trends', $params);
    }

}