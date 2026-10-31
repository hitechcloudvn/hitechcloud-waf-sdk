"""Security Posture resource (4 endpoints).

Realtime posture, per-site scoring, statistics and trends
"""

from .base import BaseResource


class SecurityPostureResource(BaseResource):
    """Security Posture: Realtime posture, per-site scoring, statistics and trends."""

    def realtime(self, params=None):
        """GET /open/security_posture/realtime - Realtime security posture."""
        return self._http.get(f"/open/security_posture/realtime", params)

    def site(self, data=None):
        """POST /open/security_posture/site - Security posture for selected sites."""
        return self._http.post(f"/open/security_posture/site", data)

    def statistics(self, params=None):
        """GET /open/security_posture/statistics - Security-posture statistics."""
        return self._http.get(f"/open/security_posture/statistics", params)

    def trends(self, params=None):
        """GET /open/security_posture/trends - Security-posture trends."""
        return self._http.get(f"/open/security_posture/trends", params)
