"""JA4 Fingerprint resource (2 endpoints).

JA4/JA4+ TLS client-fingerprint rules
"""

from .base import BaseResource


class Ja4Resource(BaseResource):
    """JA4 Fingerprint: JA4/JA4+ TLS client-fingerprint rules."""

    def get(self, params=None):
        """GET /open/ja4 - Get JA4 fingerprint rules."""
        return self._http.get(f"/open/ja4", params)

    def update(self, data=None):
        """PUT /open/ja4 - Update JA4 fingerprint rules."""
        return self._http.put(f"/open/ja4", data)
