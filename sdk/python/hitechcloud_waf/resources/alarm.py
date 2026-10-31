"""Alarms resource (3 endpoints).

Alert / notification configuration
"""

from .base import BaseResource


class AlarmResource(BaseResource):
    """Alarms: Alert / notification configuration."""

    def get(self, params=None):
        """GET /commercial/alarm - Get alarm config."""
        return self._http.get(f"/commercial/alarm", params)

    def update(self, data=None):
        """PUT /commercial/alarm - Update alarm config."""
        return self._http.put(f"/commercial/alarm", data)

    def test(self, data=None):
        """POST /commercial/alarm/test - Send a test alarm."""
        return self._http.post(f"/commercial/alarm/test", data)
