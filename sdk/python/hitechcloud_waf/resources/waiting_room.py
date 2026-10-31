"""Waiting Room resource (3 endpoints).

Virtual waiting-room queue for traffic peaks
"""

from .base import BaseResource


class WaitingRoomResource(BaseResource):
    """Waiting Room: Virtual waiting-room queue for traffic peaks."""

    def records(self, params=None):
        """GET /open/records/waiting - Waiting-room event records."""
        return self._http.get(f"/open/records/waiting", params)

    def get(self, id, params=None):
        """GET /open/site/{id}/waiting - Get waiting-room config for a site."""
        return self._http.get(f"/open/site/{id}/waiting", params)

    def set(self, id, data=None):
        """POST /open/site/{id}/waiting - Set waiting-room config for a site."""
        return self._http.post(f"/open/site/{id}/waiting", data)
