"""Settings resource (10 endpoints).

Login methods, syslog, detector, downgrade and audit log
"""

from .base import BaseResource


class SettingsResource(BaseResource):
    """Settings: Login methods, syslog, detector, downgrade and audit log."""

    def set_login_method(self, data=None):
        """PUT /business/system/login_method - Set the admin login method."""
        return self._http.put(f"/business/system/login_method", data)

    def test_login_method(self, data=None):
        """POST /business/system/login_method/test - Test the admin login method."""
        return self._http.post(f"/business/system/login_method/test", data)

    def get_syslog(self, params=None):
        """GET /commercial/syslog - Get syslog config."""
        return self._http.get(f"/commercial/syslog", params)

    def set_syslog(self, data=None):
        """POST /commercial/syslog - Set syslog config."""
        return self._http.post(f"/commercial/syslog", data)

    def test_syslog(self, data=None):
        """POST /commercial/syslog/test - Test syslog delivery."""
        return self._http.post(f"/commercial/syslog/test", data)

    def get_downgrade(self, params=None):
        """GET /open/Commercial/downgrade - Get downgrade config."""
        return self._http.get(f"/open/Commercial/downgrade", params)

    def set_downgrade(self, data=None):
        """PUT /open/Commercial/downgrade - Set downgrade config."""
        return self._http.put(f"/open/Commercial/downgrade", data)

    def audit_log(self, params=None):
        """GET /open/audit_log - Get the admin audit log."""
        return self._http.get(f"/open/audit_log", params)

    def get_detector(self, params=None):
        """GET /open/detector - Get detector config."""
        return self._http.get(f"/open/detector", params)

    def set_detector(self, data=None):
        """POST /open/detector - Set detector config."""
        return self._http.post(f"/open/detector", data)
