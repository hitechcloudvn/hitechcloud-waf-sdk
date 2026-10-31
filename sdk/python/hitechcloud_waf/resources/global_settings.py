"""Global Mode & Proxy resource (6 endpoints).

Global protection mode, outbound proxy and log cleanup
"""

from .base import BaseResource


class GlobalSettingsResource(BaseResource):
    """Global Mode & Proxy: Global protection mode, outbound proxy and log cleanup."""

    def get_mode(self, params=None):
        """GET /open/global/mode - Get the global protection mode."""
        return self._http.get(f"/open/global/mode", params)

    def set_mode(self, data=None):
        """PUT /open/global/mode - Set the global protection mode."""
        return self._http.put(f"/open/global/mode", data)

    def get_proxy(self, params=None):
        """GET /open/global/proxy - Get the global proxy config."""
        return self._http.get(f"/open/global/proxy", params)

    def set_proxy(self, data=None):
        """PUT /open/global/proxy - Set the global proxy config."""
        return self._http.put(f"/open/global/proxy", data)

    def get_log_clean(self, params=None):
        """GET /open/global/log_clean - Get the log-cleanup policy."""
        return self._http.get(f"/open/global/log_clean", params)

    def set_log_clean(self, data=None):
        """POST /open/global/log_clean - Set the log-cleanup policy."""
        return self._http.post(f"/open/global/log_clean", data)
