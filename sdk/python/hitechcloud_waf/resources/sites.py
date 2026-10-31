"""Sites (Web Services) resource (41 endpoints).

Protected sites, groups, proxy/origin config, logs and forwarding rules
"""

from .base import BaseResource


class SitesResource(BaseResource):
    """Sites (Web Services): Protected sites, groups, proxy/origin config, logs and forwarding rules."""

    def list(self, params=None):
        """GET /open/site - List protected sites."""
        return self._http.get(f"/open/site", params)

    def create(self, data=None):
        """POST /open/site - Create a protected site."""
        return self._http.post(f"/open/site", data)

    def update(self, data=None):
        """PUT /open/site - Update a site."""
        return self._http.put(f"/open/site", data)

    def delete(self, data=None):
        """DELETE /open/site - Delete a site."""
        return self._http.delete(f"/open/site", data)

    def set_challenge(self, data=None):
        """PUT /open/site/challenge - Toggle challenge mode for sites."""
        return self._http.put(f"/open/site/challenge", data)

    def set_defense(self, data=None):
        """PUT /open/site/defense - Toggle defense mode for sites."""
        return self._http.put(f"/open/site/defense", data)

    def list_groups(self, params=None):
        """GET /open/site/group - List site groups."""
        return self._http.get(f"/open/site/group", params)

    def create_group(self, data=None):
        """POST /open/site/group - Create a site group."""
        return self._http.post(f"/open/site/group", data)

    def get_group_switch(self, params=None):
        """GET /open/site/group/switch - Get site-group switch state."""
        return self._http.get(f"/open/site/group/switch", params)

    def update_group(self, id, data=None):
        """PUT /open/site/group/{id} - Update a site group."""
        return self._http.put(f"/open/site/group/{id}", data)

    def delete_group(self, id, data=None):
        """DELETE /open/site/group/{id} - Delete a site group."""
        return self._http.delete(f"/open/site/group/{id}", data)

    def sort_group(self, id, data=None):
        """PUT /open/site/group/{id}/sort - Sort a site group."""
        return self._http.put(f"/open/site/group/{id}/sort", data)

    def switch_group(self, data=None):
        """PUT /commercial/site/group/switch - Switch the active site group (commercial)."""
        return self._http.put(f"/commercial/site/group/switch", data)

    def healthcheck(self, data=None):
        """POST /open/site/healthcheck - Run an origin health check."""
        return self._http.post(f"/open/site/healthcheck", data)

    def set_mode(self, data=None):
        """PUT /open/site/mode - Set the protection mode for a site."""
        return self._http.put(f"/open/site/mode", data)

    def get(self, id, params=None):
        """GET /open/site/{id} - Get a site."""
        return self._http.get(f"/open/site/{id}", params)

    def update_basic_info(self, id, data=None):
        """PUT /open/site/{id}/basic_info - Update site basic info."""
        return self._http.put(f"/open/site/{id}/basic_info", data)

    def get_excludes(self, id, params=None):
        """GET /open/site/{id}/excludes - List exclusion rules."""
        return self._http.get(f"/open/site/{id}/excludes", params)

    def add_exclude(self, id, data=None):
        """POST /open/site/{id}/excludes - Add an exclusion rule."""
        return self._http.post(f"/open/site/{id}/excludes", data)

    def set_group(self, id, data=None):
        """PUT /open/site/{id}/group - Move a site into a group."""
        return self._http.put(f"/open/site/{id}/group", data)

    def get_log(self, id, params=None):
        """GET /open/site/{id}/log - Get site logs."""
        return self._http.get(f"/open/site/{id}/log", params)

    def get_log_detail(self, id, params=None):
        """GET /open/site/{id}/log/detail - Get a log detail."""
        return self._http.get(f"/open/site/{id}/log/detail", params)

    def download_log(self, id, params=None):
        """GET /open/site/{id}/log/download - Download site logs."""
        return self._http.get(f"/open/site/{id}/log/download", params)

    def get_log_limit(self, id, params=None):
        """GET /open/site/{id}/log/limit - Get log limit config."""
        return self._http.get(f"/open/site/{id}/log/limit", params)

    def set_log_limit(self, id, data=None):
        """POST /open/site/{id}/log/limit - Set log limit config."""
        return self._http.post(f"/open/site/{id}/log/limit", data)

    def get_nginx_config(self, id, params=None):
        """GET /open/site/{id}/nginx_config - Get generated nginx config."""
        return self._http.get(f"/open/site/{id}/nginx_config", params)

    def update_nginx_config(self, id, data=None):
        """PUT /open/site/{id}/nginx_config - Update custom nginx config."""
        return self._http.put(f"/open/site/{id}/nginx_config", data)

    def get_proxy(self, id, params=None):
        """GET /open/site/{id}/proxy - Get reverse-proxy / origin config."""
        return self._http.get(f"/open/site/{id}/proxy", params)

    def update_proxy(self, id, data=None):
        """PUT /open/site/{id}/proxy - Update reverse-proxy / origin config."""
        return self._http.put(f"/open/site/{id}/proxy", data)

    def get_resources(self, id, params=None):
        """GET /open/site/{id}/resources - List static resources."""
        return self._http.get(f"/open/site/{id}/resources", params)

    def delete_resources(self, id, data=None):
        """DELETE /open/site/{id}/resources - Delete static resources."""
        return self._http.delete(f"/open/site/{id}/resources", data)

    def sort(self, id, data=None):
        """PUT /open/site/{id}/sort - Sort a site."""
        return self._http.put(f"/open/site/{id}/sort", data)

    def get_static(self, id, params=None):
        """GET /open/site/{id}/static - List static-site rules."""
        return self._http.get(f"/open/site/{id}/static", params)

    def add_static(self, id, data=None):
        """POST /open/site/{id}/static - Add a static-site rule."""
        return self._http.post(f"/open/site/{id}/static", data)

    def update_static(self, id, data=None):
        """PUT /open/site/{id}/static - Update a static-site rule."""
        return self._http.put(f"/open/site/{id}/static", data)

    def delete_static(self, id, data=None):
        """DELETE /open/site/{id}/static - Delete a static-site rule."""
        return self._http.delete(f"/open/site/{id}/static", data)

    def list_forwarding_rules(self, id, params=None):
        """GET /open/site/{id}/forwarding_rules - List forwarding rules."""
        return self._http.get(f"/open/site/{id}/forwarding_rules", params)

    def create_forwarding_rule(self, id, data=None):
        """POST /open/site/{id}/forwarding_rules - Create a forwarding rule."""
        return self._http.post(f"/open/site/{id}/forwarding_rules", data)

    def update_forwarding_rule(self, id, rule_id, data=None):
        """PUT /open/site/{id}/forwarding_rules/{rule_id} - Update a forwarding rule."""
        return self._http.put(f"/open/site/{id}/forwarding_rules/{rule_id}", data)

    def delete_forwarding_rule(self, id, rule_id, data=None):
        """DELETE /open/site/{id}/forwarding_rules/{rule_id} - Delete a forwarding rule."""
        return self._http.delete(f"/open/site/{id}/forwarding_rules/{rule_id}", data)

    def switch_forwarding_rule(self, id, rule_id, data=None):
        """PUT /open/site/{id}/forwarding_rules/{rule_id}/switch - Enable / disable a forwarding rule."""
        return self._http.put(f"/open/site/{id}/forwarding_rules/{rule_id}/switch", data)
