"""Skynet AI Rules resource (4 endpoints).

AI-driven Skynet detection rules and switches
"""

from .base import BaseResource


class SkynetResource(BaseResource):
    """Skynet AI Rules: AI-driven Skynet detection rules and switches."""

    def get_rules(self, params=None):
        """GET /open/skynet/rule - Get Skynet rules."""
        return self._http.get(f"/open/skynet/rule", params)

    def update_rules(self, data=None):
        """PUT /open/skynet/rule - Update Skynet rules."""
        return self._http.put(f"/open/skynet/rule", data)

    def get_rule_switch(self, params=None):
        """GET /open/skynet/rule/switch - Get Skynet rule switch state."""
        return self._http.get(f"/open/skynet/rule/switch", params)

    def set_rule_switch(self, data=None):
        """PUT /open/skynet/rule/switch - Toggle Skynet rules."""
        return self._http.put(f"/open/skynet/rule/switch", data)
