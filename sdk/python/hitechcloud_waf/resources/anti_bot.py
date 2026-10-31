"""Anti-Bot & Intelligence resource (19 endpoints).

Threat intelligence, IP groups, crawlers and behaviour/fingerprint sharing
"""

from .base import BaseResource


class AntiBotResource(BaseResource):
    """Anti-Bot & Intelligence: Threat intelligence, IP groups, crawlers and behaviour/fingerprint sharing."""

    def get_intelligence(self, params=None):
        """GET /open/intelligence - Get threat-intelligence config."""
        return self._http.get(f"/open/intelligence", params)

    def set_intelligence(self, data=None):
        """POST /open/intelligence - Set threat-intelligence config."""
        return self._http.post(f"/open/intelligence", data)

    def upload_ip_lib(self, data=None):
        """POST /open/intelligence/ip_lib - Upload a custom IP library."""
        return self._http.post(f"/open/intelligence/ip_lib", data)

    def list_ip_groups(self, params=None):
        """GET /open/ipgroup - List IP groups."""
        return self._http.get(f"/open/ipgroup", params)

    def create_ip_group(self, data=None):
        """POST /open/ipgroup - Create an IP group."""
        return self._http.post(f"/open/ipgroup", data)

    def update_ip_group(self, data=None):
        """PUT /open/ipgroup - Update an IP group."""
        return self._http.put(f"/open/ipgroup", data)

    def delete_ip_group(self, data=None):
        """DELETE /open/ipgroup - Delete an IP group."""
        return self._http.delete(f"/open/ipgroup", data)

    def append_ip_group(self, data=None):
        """POST /open/ipgroup/append - Append entries to an IP group."""
        return self._http.post(f"/open/ipgroup/append", data)

    def get_crawlers(self, params=None):
        """GET /open/ipgroup/crawler - Get known-crawler groups."""
        return self._http.get(f"/open/ipgroup/crawler", params)

    def set_crawlers(self, data=None):
        """POST /open/ipgroup/crawler - Configure known-crawler groups."""
        return self._http.post(f"/open/ipgroup/crawler", data)

    def get_ip_group_detail(self, params=None):
        """GET /open/ipgroup/detail - Get IP-group detail."""
        return self._http.get(f"/open/ipgroup/detail", params)

    def get_ip_group_link(self, params=None):
        """GET /open/ipgroup/link - Get IP-group link bindings."""
        return self._http.get(f"/open/ipgroup/link", params)

    def set_ip_group_link(self, data=None):
        """POST /open/ipgroup/link - Set IP-group link bindings."""
        return self._http.post(f"/open/ipgroup/link", data)

    def get_trust_servers(self, params=None):
        """GET /open/ipgroup/tserver - Get trusted-server IP groups."""
        return self._http.get(f"/open/ipgroup/tserver", params)

    def set_trust_servers(self, data=None):
        """POST /open/ipgroup/tserver - Set trusted-server IP groups."""
        return self._http.post(f"/open/ipgroup/tserver", data)

    def get_share_behaviour(self, params=None):
        """GET /open/share_behaviour - Get behaviour-sharing config."""
        return self._http.get(f"/open/share_behaviour", params)

    def set_share_behaviour(self, data=None):
        """POST /open/share_behaviour - Set behaviour-sharing config."""
        return self._http.post(f"/open/share_behaviour", data)

    def get_share_fingerprint(self, params=None):
        """GET /open/share_fingerprint - Get fingerprint-sharing config."""
        return self._http.get(f"/open/share_fingerprint", params)

    def set_share_fingerprint(self, data=None):
        """POST /open/share_fingerprint - Set fingerprint-sharing config."""
        return self._http.post(f"/open/share_fingerprint", data)
