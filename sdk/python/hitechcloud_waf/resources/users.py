"""Users & Account resource (7 endpoints).

Manage administrative users, TOTP and the current manager profile
"""

from .base import BaseResource


class UsersResource(BaseResource):
    """Users & Account: Manage administrative users, TOTP and the current manager profile."""

    def list(self, params=None):
        """GET /open/users - List users."""
        return self._http.get(f"/open/users", params)

    def create(self, data=None):
        """POST /open/users - Create a user."""
        return self._http.post(f"/open/users", data)

    def update(self, data=None):
        """PUT /open/users - Update a user."""
        return self._http.put(f"/open/users", data)

    def delete(self, data=None):
        """DELETE /open/users - Delete a user."""
        return self._http.delete(f"/open/users", data)

    def setup_totp(self, id, data=None):
        """POST /open/users/{id}/totp - Enable / bind TOTP for a user."""
        return self._http.post(f"/open/users/{id}/totp", data)

    def get_manager_info(self, params=None):
        """GET /ManagerInfo - Get the current manager (admin) profile."""
        return self._http.get(f"/ManagerInfo", params)

    def update_manager_info(self, data=None):
        """PUT /ManagerInfo - Update the current manager (admin) profile."""
        return self._http.put(f"/ManagerInfo", data)
