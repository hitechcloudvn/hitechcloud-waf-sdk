"""Portal & Branding resource (13 endpoints).

Self-service portal, block pages, branding and accounts
"""

from .base import BaseResource


class PortalResource(BaseResource):
    """Portal & Branding: Self-service portal, block pages, branding and accounts."""

    def update_style(self, data=None):
        """PUT /commercial/portal/style - Update portal style (commercial)."""
        return self._http.put(f"/commercial/portal/style", data)

    def get(self, params=None):
        """GET /open/portal - Get portal config."""
        return self._http.get(f"/open/portal", params)

    def update(self, data=None):
        """PUT /open/portal - Update portal config."""
        return self._http.put(f"/open/portal", data)

    def get_proxy_config(self, params=None):
        """GET /open/portal/proxy_config - Get portal proxy config."""
        return self._http.get(f"/open/portal/proxy_config", params)

    def set_proxy_config(self, data=None):
        """PUT /open/portal/proxy_config - Set portal proxy config."""
        return self._http.put(f"/open/portal/proxy_config", data)

    def get_style(self, params=None):
        """GET /open/portal/style - Get portal style."""
        return self._http.get(f"/open/portal/style", params)

    def update_block_page(self, data=None):
        """PUT /commercial/block_page - Update the block page (commercial)."""
        return self._http.put(f"/commercial/block_page", data)

    def get_block_page(self, params=None):
        """GET /open/block_page - Get the active block page."""
        return self._http.get(f"/open/block_page", params)

    def list_block_pages(self, params=None):
        """GET /open/block_page_list - List available block pages."""
        return self._http.get(f"/open/block_page_list", params)

    def get_frontend_style(self, params=None):
        """GET /business/frontend_style - Get frontend style."""
        return self._http.get(f"/business/frontend_style", params)

    def set_frontend_style(self, data=None):
        """PUT /business/frontend_style - Set frontend style."""
        return self._http.put(f"/business/frontend_style", data)

    def get_account(self, params=None):
        """GET /business/account - Get portal account(s)."""
        return self._http.get(f"/business/account", params)

    def create_account(self, data=None):
        """POST /business/account - Create a portal account."""
        return self._http.post(f"/business/account", data)
