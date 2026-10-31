<?php

namespace HiTechCloud\WAF\Resources;

/**
 * Dynamic Authentication Defense: Authentication-defense sources, protected users, reviews and records.
 */
class AuthDefense extends BaseResource
{
    /** GET /open/auth_defense/source - List authentication-defense sources */
    public function listSources(array $params = []): ?array
    {
        return $this->http->get('/open/auth_defense/source', $params);
    }

    /** POST /open/auth_defense/source - Create a source */
    public function createSource(?array $data = null): ?array
    {
        return $this->http->post('/open/auth_defense/source', $data);
    }

    /** GET /open/auth_defense/source/password - Get the shared source password policy */
    public function getSourcePassword(array $params = []): ?array
    {
        return $this->http->get('/open/auth_defense/source/password', $params);
    }

    /** PUT /open/auth_defense/source/password - Update the shared source password policy */
    public function updateSourcePassword(?array $data = null): ?array
    {
        return $this->http->put('/open/auth_defense/source/password', $data);
    }

    /** GET /open/auth_defense/source/{id} - Get a source */
    public function getSource(string $id, array $params = []): ?array
    {
        return $this->http->get('/open/auth_defense/source/' . $id, $params);
    }

    /** PUT /open/auth_defense/source/{id} - Update a source */
    public function updateSource(string $id, ?array $data = null): ?array
    {
        return $this->http->put('/open/auth_defense/source/' . $id, $data);
    }

    /** DELETE /open/auth_defense/source/{id} - Delete a source */
    public function deleteSource(string $id, ?array $data = null): ?array
    {
        return $this->http->delete('/open/auth_defense/source/' . $id, $data);
    }

    /** GET /open/auth_defense/source/{id}/secret - Get a source secret */
    public function getSourceSecret(string $id, array $params = []): ?array
    {
        return $this->http->get('/open/auth_defense/source/' . $id . '/secret', $params);
    }

    /** GET /open/auth_defense/source/{id}/user - List users bound to a source */
    public function listSourceUsers(string $id, array $params = []): ?array
    {
        return $this->http->get('/open/auth_defense/source/' . $id . '/user', $params);
    }

    /** PUT /open/auth_defense/source/{id}/user/{user_id} - Update a source user */
    public function updateSourceUser(string $id, string $userId, ?array $data = null): ?array
    {
        return $this->http->put('/open/auth_defense/source/' . $id . '/user/' . $userId, $data);
    }

    /** DELETE /open/auth_defense/source/{id}/user/{user_id} - Remove a user from a source */
    public function deleteSourceUser(string $id, string $userId, ?array $data = null): ?array
    {
        return $this->http->delete('/open/auth_defense/source/' . $id . '/user/' . $userId, $data);
    }

    /** GET /open/auth_defense/user - List protected users */
    public function listUsers(array $params = []): ?array
    {
        return $this->http->get('/open/auth_defense/user', $params);
    }

    /** POST /open/auth_defense/user - Create a protected user */
    public function createUser(?array $data = null): ?array
    {
        return $this->http->post('/open/auth_defense/user', $data);
    }

    /** DELETE /open/auth_defense/user - Delete protected users (bulk) */
    public function deleteUsers(?array $data = null): ?array
    {
        return $this->http->delete('/open/auth_defense/user', $data);
    }

    /** POST /open/auth_defense/user/merge - Merge protected users */
    public function mergeUsers(?array $data = null): ?array
    {
        return $this->http->post('/open/auth_defense/user/merge', $data);
    }

    /** GET /open/auth_defense/user/review - List pending user reviews */
    public function listUserReviews(array $params = []): ?array
    {
        return $this->http->get('/open/auth_defense/user/review', $params);
    }

    /** PUT /open/auth_defense/user/review/{review_id} - Approve / reject a user review */
    public function updateUserReview(string $reviewId, ?array $data = null): ?array
    {
        return $this->http->put('/open/auth_defense/user/review/' . $reviewId, $data);
    }

    /** GET /open/auth_defense/user/{user_id} - Get a protected user */
    public function getUser(string $userId, array $params = []): ?array
    {
        return $this->http->get('/open/auth_defense/user/' . $userId, $params);
    }

    /** PUT /open/auth_defense/user/{user_id} - Update a protected user */
    public function updateUser(string $userId, ?array $data = null): ?array
    {
        return $this->http->put('/open/auth_defense/user/' . $userId, $data);
    }

    /** DELETE /open/auth_defense/user/{user_id} - Delete a protected user */
    public function deleteUser(string $userId, ?array $data = null): ?array
    {
        return $this->http->delete('/open/auth_defense/user/' . $userId, $data);
    }

    /** PUT /open/auth_defense/user/{user_id}/reset_totp - Reset a user's TOTP */
    public function resetUserTotp(string $userId, ?array $data = null): ?array
    {
        return $this->http->put('/open/auth_defense/user/' . $userId . '/reset_totp', $data);
    }

    /** PUT /open/auth_defense/user/{user_id}/unbind - Unbind a user's authenticator */
    public function unbindUser(string $userId, ?array $data = null): ?array
    {
        return $this->http->put('/open/auth_defense/user/' . $userId . '/unbind', $data);
    }

    /** GET /open/records/auth_defense - Authentication-defense event records */
    public function records(array $params = []): ?array
    {
        return $this->http->get('/open/records/auth_defense', $params);
    }

    /** GET /open/v2/records/auth_defense - Authentication-defense event records (v2) */
    public function recordsV2(array $params = []): ?array
    {
        return $this->http->get('/open/v2/records/auth_defense', $params);
    }

    /** POST /business/users/{id}/unlock - Unlock a locked business user */
    public function unlockUser(string $id, ?array $data = null): ?array
    {
        return $this->http->post('/business/users/' . $id . '/unlock', $data);
    }

}