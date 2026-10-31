"""Statistics resource (11 endpoints).

Advanced access/attack analytics and QPS
"""

from .base import BaseResource


class StatisticsResource(BaseResource):
    """Statistics: Advanced access/attack analytics and QPS."""

    def access(self, params=None):
        """GET /stat/advance/access - Access statistics."""
        return self._http.get(f"/stat/advance/access", params)

    def attack(self, params=None):
        """GET /stat/advance/attack - Attack statistics."""
        return self._http.get(f"/stat/advance/attack", params)

    def client(self, params=None):
        """GET /stat/advance/client - Client statistics."""
        return self._http.get(f"/stat/advance/client", params)

    def domain(self, params=None):
        """GET /stat/advance/domain - Per-domain statistics."""
        return self._http.get(f"/stat/advance/domain", params)

    def error_status_code(self, params=None):
        """GET /stat/advance/error_status_code - Error status-code statistics."""
        return self._http.get(f"/stat/advance/error_status_code", params)

    def location(self, params=None):
        """GET /stat/advance/location - Geo-location statistics."""
        return self._http.get(f"/stat/advance/location", params)

    def page(self, params=None):
        """GET /stat/advance/page - Per-page statistics."""
        return self._http.get(f"/stat/advance/page", params)

    def status_code(self, params=None):
        """GET /stat/advance/status_code - Status-code statistics."""
        return self._http.get(f"/stat/advance/status_code", params)

    def trend_access(self, params=None):
        """GET /stat/advance/trend/access - Access trend."""
        return self._http.get(f"/stat/advance/trend/access", params)

    def trend_intercept(self, params=None):
        """GET /stat/advance/trend/intercept - Interception trend."""
        return self._http.get(f"/stat/advance/trend/intercept", params)

    def qps(self, params=None):
        """GET /stat/qps - Current QPS."""
        return self._http.get(f"/stat/qps", params)
