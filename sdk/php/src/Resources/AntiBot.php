<?php

namespace HiTechCloud\WAF\Resources;

/**
 * Anti-Bot & Intelligence: Threat intelligence, IP groups, crawlers and behaviour/fingerprint sharing.
 */
class AntiBot extends BaseResource
{
    /** GET /open/intelligence - Get threat-intelligence config */
    public function getIntelligence(array $params = []): ?array
    {
        return $this->http->get('/open/intelligence', $params);
    }

    /** POST /open/intelligence - Set threat-intelligence config */
    public function setIntelligence(?array $data = null): ?array
    {
        return $this->http->post('/open/intelligence', $data);
    }

    /** POST /open/intelligence/ip_lib - Upload a custom IP library */
    public function uploadIpLib(?array $data = null): ?array
    {
        return $this->http->post('/open/intelligence/ip_lib', $data);
    }

    /** GET /open/ipgroup - List IP groups */
    public function listIpGroups(array $params = []): ?array
    {
        return $this->http->get('/open/ipgroup', $params);
    }

    /** POST /open/ipgroup - Create an IP group */
    public function createIpGroup(?array $data = null): ?array
    {
        return $this->http->post('/open/ipgroup', $data);
    }

    /** PUT /open/ipgroup - Update an IP group */
    public function updateIpGroup(?array $data = null): ?array
    {
        return $this->http->put('/open/ipgroup', $data);
    }

    /** DELETE /open/ipgroup - Delete an IP group */
    public function deleteIpGroup(?array $data = null): ?array
    {
        return $this->http->delete('/open/ipgroup', $data);
    }

    /** POST /open/ipgroup/append - Append entries to an IP group */
    public function appendIpGroup(?array $data = null): ?array
    {
        return $this->http->post('/open/ipgroup/append', $data);
    }

    /** GET /open/ipgroup/crawler - Get known-crawler groups */
    public function getCrawlers(array $params = []): ?array
    {
        return $this->http->get('/open/ipgroup/crawler', $params);
    }

    /** POST /open/ipgroup/crawler - Configure known-crawler groups */
    public function setCrawlers(?array $data = null): ?array
    {
        return $this->http->post('/open/ipgroup/crawler', $data);
    }

    /** GET /open/ipgroup/detail - Get IP-group detail */
    public function getIpGroupDetail(array $params = []): ?array
    {
        return $this->http->get('/open/ipgroup/detail', $params);
    }

    /** GET /open/ipgroup/link - Get IP-group link bindings */
    public function getIpGroupLink(array $params = []): ?array
    {
        return $this->http->get('/open/ipgroup/link', $params);
    }

    /** POST /open/ipgroup/link - Set IP-group link bindings */
    public function setIpGroupLink(?array $data = null): ?array
    {
        return $this->http->post('/open/ipgroup/link', $data);
    }

    /** GET /open/ipgroup/tserver - Get trusted-server IP groups */
    public function getTrustServers(array $params = []): ?array
    {
        return $this->http->get('/open/ipgroup/tserver', $params);
    }

    /** POST /open/ipgroup/tserver - Set trusted-server IP groups */
    public function setTrustServers(?array $data = null): ?array
    {
        return $this->http->post('/open/ipgroup/tserver', $data);
    }

    /** GET /open/share_behaviour - Get behaviour-sharing config */
    public function getShareBehaviour(array $params = []): ?array
    {
        return $this->http->get('/open/share_behaviour', $params);
    }

    /** POST /open/share_behaviour - Set behaviour-sharing config */
    public function setShareBehaviour(?array $data = null): ?array
    {
        return $this->http->post('/open/share_behaviour', $data);
    }

    /** GET /open/share_fingerprint - Get fingerprint-sharing config */
    public function getShareFingerprint(array $params = []): ?array
    {
        return $this->http->get('/open/share_fingerprint', $params);
    }

    /** POST /open/share_fingerprint - Set fingerprint-sharing config */
    public function setShareFingerprint(?array $data = null): ?array
    {
        return $this->http->post('/open/share_fingerprint', $data);
    }

}