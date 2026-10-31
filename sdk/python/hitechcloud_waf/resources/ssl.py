"""SSL Certificates resource (4 endpoints).

TLS/SSL certificate management
"""

from .base import BaseResource


class SslResource(BaseResource):
    """SSL Certificates: TLS/SSL certificate management."""

    def list(self, params=None):
        """GET /open/cert - List certificates."""
        return self._http.get(f"/open/cert", params)

    def create(self, data=None):
        """POST /open/cert - Upload / create a certificate."""
        return self._http.post(f"/open/cert", data)

    def get(self, id, params=None):
        """GET /open/cert/{id} - Get a certificate."""
        return self._http.get(f"/open/cert/{id}", params)

    def delete(self, id, data=None):
        """DELETE /open/cert/{id} - Delete a certificate."""
        return self._http.delete(f"/open/cert/{id}", data)
