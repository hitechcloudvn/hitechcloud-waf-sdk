"""Dynamic Authentication Defense resource (25 endpoints).

Authentication-defense sources, protected users, reviews and records
"""

from .base import BaseResource


class AuthDefenseResource(BaseResource):
    """Dynamic Authentication Defense: Authentication-defense sources, protected users, reviews and records."""

    def list_sources(self, params=None):
        """GET /open/auth_defense/source - List authentication-defense sources."""
        return self._http.get(f"/open/auth_defense/source", params)

    def create_source(self, data=None):
        """POST /open/auth_defense/source - Create a source."""
        return self._http.post(f"/open/auth_defense/source", data)

    def get_source_password(self, params=None):
        """GET /open/auth_defense/source/password - Get the shared source password policy."""
        return self._http.get(f"/open/auth_defense/source/password", params)

    def update_source_password(self, data=None):
        """PUT /open/auth_defense/source/password - Update the shared source password policy."""
        return self._http.put(f"/open/auth_defense/source/password", data)

    def get_source(self, id, params=None):
        """GET /open/auth_defense/source/{id} - Get a source."""
        return self._http.get(f"/open/auth_defense/source/{id}", params)

    def update_source(self, id, data=None):
        """PUT /open/auth_defense/source/{id} - Update a source."""
        return self._http.put(f"/open/auth_defense/source/{id}", data)

    def delete_source(self, id, data=None):
        """DELETE /open/auth_defense/source/{id} - Delete a source."""
        return self._http.delete(f"/open/auth_defense/source/{id}", data)

    def get_source_secret(self, id, params=None):
        """GET /open/auth_defense/source/{id}/secret - Get a source secret."""
        return self._http.get(f"/open/auth_defense/source/{id}/secret", params)

    def list_source_users(self, id, params=None):
        """GET /open/auth_defense/source/{id}/user - List users bound to a source."""
        return self._http.get(f"/open/auth_defense/source/{id}/user", params)

    def update_source_user(self, id, user_id, data=None):
        """PUT /open/auth_defense/source/{id}/user/{user_id} - Update a source user."""
        return self._http.put(f"/open/auth_defense/source/{id}/user/{user_id}", data)

    def delete_source_user(self, id, user_id, data=None):
        """DELETE /open/auth_defense/source/{id}/user/{user_id} - Remove a user from a source."""
        return self._http.delete(f"/open/auth_defense/source/{id}/user/{user_id}", data)

    def list_users(self, params=None):
        """GET /open/auth_defense/user - List protected users."""
        return self._http.get(f"/open/auth_defense/user", params)

    def create_user(self, data=None):
        """POST /open/auth_defense/user - Create a protected user."""
        return self._http.post(f"/open/auth_defense/user", data)

    def delete_users(self, data=None):
        """DELETE /open/auth_defense/user - Delete protected users (bulk)."""
        return self._http.delete(f"/open/auth_defense/user", data)

    def merge_users(self, data=None):
        """POST /open/auth_defense/user/merge - Merge protected users."""
        return self._http.post(f"/open/auth_defense/user/merge", data)

    def list_user_reviews(self, params=None):
        """GET /open/auth_defense/user/review - List pending user reviews."""
        return self._http.get(f"/open/auth_defense/user/review", params)

    def update_user_review(self, review_id, data=None):
        """PUT /open/auth_defense/user/review/{review_id} - Approve / reject a user review."""
        return self._http.put(f"/open/auth_defense/user/review/{review_id}", data)

    def get_user(self, user_id, params=None):
        """GET /open/auth_defense/user/{user_id} - Get a protected user."""
        return self._http.get(f"/open/auth_defense/user/{user_id}", params)

    def update_user(self, user_id, data=None):
        """PUT /open/auth_defense/user/{user_id} - Update a protected user."""
        return self._http.put(f"/open/auth_defense/user/{user_id}", data)

    def delete_user(self, user_id, data=None):
        """DELETE /open/auth_defense/user/{user_id} - Delete a protected user."""
        return self._http.delete(f"/open/auth_defense/user/{user_id}", data)

    def reset_user_totp(self, user_id, data=None):
        """PUT /open/auth_defense/user/{user_id}/reset_totp - Reset a user's TOTP."""
        return self._http.put(f"/open/auth_defense/user/{user_id}/reset_totp", data)

    def unbind_user(self, user_id, data=None):
        """PUT /open/auth_defense/user/{user_id}/unbind - Unbind a user's authenticator."""
        return self._http.put(f"/open/auth_defense/user/{user_id}/unbind", data)

    def records(self, params=None):
        """GET /open/records/auth_defense - Authentication-defense event records."""
        return self._http.get(f"/open/records/auth_defense", params)

    def records_v_2(self, params=None):
        """GET /open/v2/records/auth_defense - Authentication-defense event records (v2)."""
        return self._http.get(f"/open/v2/records/auth_defense", params)

    def unlock_user(self, id, data=None):
        """POST /business/users/{id}/unlock - Unlock a locked business user."""
        return self._http.post(f"/business/users/{id}/unlock", data)
