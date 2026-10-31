"""Policies & Rules resource (9 endpoints).

Custom protection policies and rule ordering
"""

from .base import BaseResource


class PoliciesResource(BaseResource):
    """Policies & Rules: Custom protection policies and rule ordering."""

    def list(self, params=None):
        """GET /open/policy - List policies."""
        return self._http.get(f"/open/policy", params)

    def create(self, data=None):
        """POST /open/policy - Create a policy."""
        return self._http.post(f"/open/policy", data)

    def update(self, data=None):
        """PUT /open/policy - Update a policy."""
        return self._http.put(f"/open/policy", data)

    def delete(self, data=None):
        """DELETE /open/policy - Delete a policy."""
        return self._http.delete(f"/open/policy", data)

    def detail(self, params=None):
        """GET /open/policy/detail - Get policy detail."""
        return self._http.get(f"/open/policy/detail", params)

    def set_order(self, data=None):
        """PUT /open/policy/order - Set a policy order."""
        return self._http.put(f"/open/policy/order", data)

    def set_order_value(self, data=None):
        """PUT /open/policy/order/value - Set a policy order value."""
        return self._http.put(f"/open/policy/order/value", data)

    def set_orders(self, data=None):
        """PUT /open/policy/orders - Bulk-set policy orders."""
        return self._http.put(f"/open/policy/orders", data)

    def set_switch(self, data=None):
        """PUT /open/policy/switch - Toggle a policy on/off."""
        return self._http.put(f"/open/policy/switch", data)
