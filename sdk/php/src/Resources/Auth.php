<?php

namespace HiTechCloud\WAF\Resources;

/**
 * Authentication & Token: CSRF, login, TFA, logout and API token lifecycle.
 */
class Auth extends BaseResource
{
    /** GET /open/auth/csrf - Get a CSRF token for the login step */
    public function csrf(array $params = []): ?array
    {
        return $this->http->get('/open/auth/csrf', $params);
    }

    /** POST /open/auth/login - Authenticate with username/password (+csrf_token) and obtain a JWT */
    public function login(?array $data = null): ?array
    {
        return $this->http->post('/open/auth/login', $data);
    }

    /** POST /open/auth/logout - End the current session */
    public function logout(?array $data = null): ?array
    {
        return $this->http->post('/open/auth/logout', $data);
    }

    /** POST /open/auth/tfa - Verify a two-factor (OTP) code during login */
    public function tfa(?array $data = null): ?array
    {
        return $this->http->post('/open/auth/tfa', $data);
    }

    /** GET /open/auth/token - Get the current API token */
    public function getToken(array $params = []): ?array
    {
        return $this->http->get('/open/auth/token', $params);
    }

    /** PUT /open/auth/token - Rotate / refresh the API token */
    public function refreshToken(?array $data = null): ?array
    {
        return $this->http->put('/open/auth/token', $data);
    }

    /** DELETE /open/auth/token - Revoke the API token */
    public function revokeToken(?array $data = null): ?array
    {
        return $this->http->delete('/open/auth/token', $data);
    }

}