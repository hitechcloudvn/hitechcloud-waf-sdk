<?php

namespace HiTechCloud\WAF\Resources;

/**
 * Policies & Rules: Custom protection policies and rule ordering.
 */
class Policies extends BaseResource
{
    /** GET /open/policy - List policies */
    public function list(array $params = []): ?array
    {
        return $this->http->get('/open/policy', $params);
    }

    /** POST /open/policy - Create a policy */
    public function create(?array $data = null): ?array
    {
        return $this->http->post('/open/policy', $data);
    }

    /** PUT /open/policy - Update a policy */
    public function update(?array $data = null): ?array
    {
        return $this->http->put('/open/policy', $data);
    }

    /** DELETE /open/policy - Delete a policy */
    public function delete(?array $data = null): ?array
    {
        return $this->http->delete('/open/policy', $data);
    }

    /** GET /open/policy/detail - Get policy detail */
    public function detail(array $params = []): ?array
    {
        return $this->http->get('/open/policy/detail', $params);
    }

    /** PUT /open/policy/order - Set a policy order */
    public function setOrder(?array $data = null): ?array
    {
        return $this->http->put('/open/policy/order', $data);
    }

    /** PUT /open/policy/order/value - Set a policy order value */
    public function setOrderValue(?array $data = null): ?array
    {
        return $this->http->put('/open/policy/order/value', $data);
    }

    /** PUT /open/policy/orders - Bulk-set policy orders */
    public function setOrders(?array $data = null): ?array
    {
        return $this->http->put('/open/policy/orders', $data);
    }

    /** PUT /open/policy/switch - Toggle a policy on/off */
    public function setSwitch(?array $data = null): ?array
    {
        return $this->http->put('/open/policy/switch', $data);
    }

}