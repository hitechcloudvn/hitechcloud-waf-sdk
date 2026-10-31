"""Anti-Tamper resource (8 endpoints).

Content anti-tamper protection per site and resource
"""

from .base import BaseResource


class AntiTamperResource(BaseResource):
    """Anti-Tamper: Content anti-tamper protection per site and resource."""

    def get(self, id, params=None):
        """GET /business/anti_tamper/{id} - Get an anti-tamper config."""
        return self._http.get(f"/business/anti_tamper/{id}", params)

    def update(self, id, data=None):
        """PUT /business/anti_tamper/{id} - Update an anti-tamper config."""
        return self._http.put(f"/business/anti_tamper/{id}", data)

    def delete(self, id, data=None):
        """DELETE /business/anti_tamper/{id} - Delete an anti-tamper config."""
        return self._http.delete(f"/business/anti_tamper/{id}", data)

    def get_page(self, id, params=None):
        """GET /business/anti_tamper/{id}/page - Get protected pages for a config."""
        return self._http.get(f"/business/anti_tamper/{id}/page", params)

    def list_for_site(self, site_id, params=None):
        """GET /business/site/{site_id}/anti_tamper - List anti-tamper configs for a site."""
        return self._http.get(f"/business/site/{site_id}/anti_tamper", params)

    def create_for_site(self, site_id, data=None):
        """POST /business/site/{site_id}/anti_tamper - Create an anti-tamper config for a site."""
        return self._http.post(f"/business/site/{site_id}/anti_tamper", data)

    def update_for_site(self, site_id, data=None):
        """PUT /business/site/{site_id}/anti_tamper - Update anti-tamper for a site."""
        return self._http.put(f"/business/site/{site_id}/anti_tamper", data)

    def get_resource(self, site_id, resource_id, params=None):
        """GET /business/site/{site_id}/anti_tamper/resource/{resource_id} - Get a protected resource."""
        return self._http.get(f"/business/site/{site_id}/anti_tamper/resource/{resource_id}", params)
