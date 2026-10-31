"""System resource (15 endpoints).

System info, edition, protocols, network proxy and helpers
"""

from .base import BaseResource


class SystemResource(BaseResource):
    """System: System info, edition, protocols, network proxy and helpers."""

    def get(self, params=None):
        """GET /open/system - Get system settings."""
        return self._http.get(f"/open/system", params)

    def update(self, data=None):
        """PUT /open/system - Update system settings."""
        return self._http.put(f"/open/system", data)

    def get_arch(self, params=None):
        """GET /open/system/arch - Get the cluster architecture."""
        return self._http.get(f"/open/system/arch", params)

    def get_customer_service(self, params=None):
        """GET /open/system/customer_service - Get customer-service config."""
        return self._http.get(f"/open/system/customer_service", params)

    def set_customer_service(self, data=None):
        """PUT /open/system/customer_service - Set customer-service config."""
        return self._http.put(f"/open/system/customer_service", data)

    def get_edition(self, params=None):
        """GET /open/system/edition - Get the product edition."""
        return self._http.get(f"/open/system/edition", params)

    def get_key(self, params=None):
        """GET /open/system/key - Get the system public key."""
        return self._http.get(f"/open/system/key", params)

    def get_login_method(self, params=None):
        """GET /open/system/login_method - Get the system login method."""
        return self._http.get(f"/open/system/login_method", params)

    def get_login_type(self, params=None):
        """GET /open/system/login_type - Get the system login type."""
        return self._http.get(f"/open/system/login_type", params)

    def get_network_proxy(self, params=None):
        """GET /open/system/network_proxy - Get the network proxy."""
        return self._http.get(f"/open/system/network_proxy", params)

    def set_network_proxy(self, data=None):
        """POST /open/system/network_proxy - Set the network proxy."""
        return self._http.post(f"/open/system/network_proxy", data)

    def get_protocol(self, params=None):
        """GET /open/system/protocol - Get supported protocols."""
        return self._http.get(f"/open/system/protocol", params)

    def set_protocol(self, data=None):
        """POST /open/system/protocol - Configure supported protocols."""
        return self._http.post(f"/open/system/protocol", data)

    def dashboard_token(self, data=None):
        """POST /commercial/dashboard/token - Issue a dashboard token."""
        return self._http.post(f"/commercial/dashboard/token", data)

    def safepoint_consult(self, data=None):
        """POST /open/safepoint/consult - Consult the safepoint advisor."""
        return self._http.post(f"/open/safepoint/consult", data)
