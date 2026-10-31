"""ACL & Rate-Limit resource (7 endpoints).

Access-control / rate-limit rules and block relief
"""

from .base import BaseResource


class AclResource(BaseResource):
    """ACL & Rate-Limit: Access-control / rate-limit rules and block relief."""

    def relieve(self, data=None):
        """PUT /acl/relieve - Relieve (unblock) an ACL-blocked client."""
        return self._http.put(f"/acl/relieve", data)

    def records(self, params=None):
        """GET /open/records/acl - ACL hit records."""
        return self._http.get(f"/open/records/acl", params)

    def list(self, id, params=None):
        """GET /open/site/{id}/acl - List ACL / rate-limit rules for a site."""
        return self._http.get(f"/open/site/{id}/acl", params)

    def create(self, id, data=None):
        """POST /open/site/{id}/acl - Create an ACL rule."""
        return self._http.post(f"/open/site/{id}/acl", data)

    def update(self, id, data=None):
        """PUT /open/site/{id}/acl - Update ACL rules for a site."""
        return self._http.put(f"/open/site/{id}/acl", data)

    def update_rule(self, id, rule_id, data=None):
        """PUT /open/site/{id}/acl/{rule_id} - Update a single ACL rule."""
        return self._http.put(f"/open/site/{id}/acl/{rule_id}", data)

    def delete_rule(self, id, rule_id, data=None):
        """DELETE /open/site/{id}/acl/{rule_id} - Delete a single ACL rule."""
        return self._http.delete(f"/open/site/{id}/acl/{rule_id}", data)
