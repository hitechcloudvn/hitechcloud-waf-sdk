<?php

namespace HiTechCloud\WAF\Resources;

/**
 * System: System info, edition, protocols, network proxy and helpers.
 */
class System extends BaseResource
{
    /** GET /open/system - Get system settings */
    public function get(array $params = []): ?array
    {
        return $this->http->get('/open/system', $params);
    }

    /** PUT /open/system - Update system settings */
    public function update(?array $data = null): ?array
    {
        return $this->http->put('/open/system', $data);
    }

    /** GET /open/system/arch - Get the cluster architecture */
    public function getArch(array $params = []): ?array
    {
        return $this->http->get('/open/system/arch', $params);
    }

    /** GET /open/system/customer_service - Get customer-service config */
    public function getCustomerService(array $params = []): ?array
    {
        return $this->http->get('/open/system/customer_service', $params);
    }

    /** PUT /open/system/customer_service - Set customer-service config */
    public function setCustomerService(?array $data = null): ?array
    {
        return $this->http->put('/open/system/customer_service', $data);
    }

    /** GET /open/system/edition - Get the product edition */
    public function getEdition(array $params = []): ?array
    {
        return $this->http->get('/open/system/edition', $params);
    }

    /** GET /open/system/key - Get the system public key */
    public function getKey(array $params = []): ?array
    {
        return $this->http->get('/open/system/key', $params);
    }

    /** GET /open/system/login_method - Get the system login method */
    public function getLoginMethod(array $params = []): ?array
    {
        return $this->http->get('/open/system/login_method', $params);
    }

    /** GET /open/system/login_type - Get the system login type */
    public function getLoginType(array $params = []): ?array
    {
        return $this->http->get('/open/system/login_type', $params);
    }

    /** GET /open/system/network_proxy - Get the network proxy */
    public function getNetworkProxy(array $params = []): ?array
    {
        return $this->http->get('/open/system/network_proxy', $params);
    }

    /** POST /open/system/network_proxy - Set the network proxy */
    public function setNetworkProxy(?array $data = null): ?array
    {
        return $this->http->post('/open/system/network_proxy', $data);
    }

    /** GET /open/system/protocol - Get supported protocols */
    public function getProtocol(array $params = []): ?array
    {
        return $this->http->get('/open/system/protocol', $params);
    }

    /** POST /open/system/protocol - Configure supported protocols */
    public function setProtocol(?array $data = null): ?array
    {
        return $this->http->post('/open/system/protocol', $data);
    }

    /** POST /commercial/dashboard/token - Issue a dashboard token */
    public function dashboardToken(?array $data = null): ?array
    {
        return $this->http->post('/commercial/dashboard/token', $data);
    }

    /** POST /open/safepoint/consult - Consult the safepoint advisor */
    public function safepointConsult(?array $data = null): ?array
    {
        return $this->http->post('/open/safepoint/consult', $data);
    }

}