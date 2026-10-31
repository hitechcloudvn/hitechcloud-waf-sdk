"""Cloud Policies resource (2 endpoints).

Shared cloud threat-intelligence policies
"""

from .base import BaseResource


class CloudResource(BaseResource):
    """Cloud Policies: Shared cloud threat-intelligence policies."""

    def list_policies(self, params=None):
        """GET /open/cloud/policies - List available cloud policies."""
        return self._http.get(f"/open/cloud/policies", params)

    def subscribe(self, data=None):
        """POST /open/cloud/policies/subscribe - Subscribe to a cloud policy."""
        return self._http.post(f"/open/cloud/policies/subscribe", data)
