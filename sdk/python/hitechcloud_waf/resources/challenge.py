"""Challenge & Semantics resource (7 endpoints).

Human-verification challenge, semantic analysis and chaos protection
"""

from .base import BaseResource


class ChallengeResource(BaseResource):
    """Challenge & Semantics: Human-verification challenge, semantic analysis and chaos protection."""

    def get_config(self, params=None):
        """GET /open/challenge/config - Get challenge config."""
        return self._http.get(f"/open/challenge/config", params)

    def set_config(self, data=None):
        """POST /open/challenge/config - Set challenge config."""
        return self._http.post(f"/open/challenge/config", data)

    def records(self, params=None):
        """GET /open/records/challenge - Challenge event records."""
        return self._http.get(f"/open/records/challenge", params)

    def get_semantics(self, id, params=None):
        """GET /open/site/{id}/semantics - Get semantic-analysis config for a site."""
        return self._http.get(f"/open/site/{id}/semantics", params)

    def set_semantics(self, id, data=None):
        """PUT /open/site/{id}/semantics - Set semantic-analysis config for a site."""
        return self._http.put(f"/open/site/{id}/semantics", data)

    def get_chaos(self, id, params=None):
        """GET /open/site/{id}/chaos - Get chaos/dynamic-protection config for a site."""
        return self._http.get(f"/open/site/{id}/chaos", params)

    def set_chaos(self, id, data=None):
        """POST /open/site/{id}/chaos - Set chaos/dynamic-protection config for a site."""
        return self._http.post(f"/open/site/{id}/chaos", data)
