<?php

namespace HiTechCloud\WAF\Resources;

/**
 * Users & Account: Manage administrative users, TOTP and the current manager profile.
 */
class Users extends BaseResource
{
    /** GET /open/users - List users */
    public function list(array $params = []): ?array
    {
        return $this->http->get('/open/users', $params);
    }

    /** POST /open/users - Create a user */
    public function create(?array $data = null): ?array
    {
        return $this->http->post('/open/users', $data);
    }

    /** PUT /open/users - Update a user */
    public function update(?array $data = null): ?array
    {
        return $this->http->put('/open/users', $data);
    }

    /** DELETE /open/users - Delete a user */
    public function delete(?array $data = null): ?array
    {
        return $this->http->delete('/open/users', $data);
    }

    /** POST /open/users/{id}/totp - Enable / bind TOTP for a user */
    public function setupTotp(string $id, ?array $data = null): ?array
    {
        return $this->http->post('/open/users/' . $id . '/totp', $data);
    }

    /** GET /ManagerInfo - Get the current manager (admin) profile */
    public function getManagerInfo(array $params = []): ?array
    {
        return $this->http->get('/ManagerInfo', $params);
    }

    /** PUT /ManagerInfo - Update the current manager (admin) profile */
    public function updateManagerInfo(?array $data = null): ?array
    {
        return $this->http->put('/ManagerInfo', $data);
    }

}