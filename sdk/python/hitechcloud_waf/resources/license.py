"""License & Authorization resource (7 endpoints).

Authorization records and license status / lifecycle
"""

from .base import BaseResource


class LicenseResource(BaseResource):
    """License & Authorization: Authorization records and license status / lifecycle."""

    def get_authorize(self, params=None):
        """GET /open/system/authorize - Get authorization records."""
        return self._http.get(f"/open/system/authorize", params)

    def create_authorize(self, data=None):
        """POST /open/system/authorize - Create an authorization."""
        return self._http.post(f"/open/system/authorize", data)

    def update_authorize(self, data=None):
        """PUT /open/system/authorize - Update an authorization."""
        return self._http.put(f"/open/system/authorize", data)

    def delete_authorize(self, data=None):
        """DELETE /open/system/authorize - Delete an authorization."""
        return self._http.delete(f"/open/system/authorize", data)

    def confirm_downgrade(self, data=None):
        """POST /open/system/license/expired/confirm-downgrade - Confirm downgrade after license expiry."""
        return self._http.post(f"/open/system/license/expired/confirm-downgrade", data)

    def acknowledge_expiry(self, data=None):
        """POST /open/system/license/expiry/acknowledge - Acknowledge an upcoming license expiry."""
        return self._http.post(f"/open/system/license/expiry/acknowledge", data)

    def status(self, params=None):
        """GET /open/system/license/status - Get the license status."""
        return self._http.get(f"/open/system/license/status", params)
