"""Reports resource (4 endpoints).

Scheduled / on-demand security reports
"""

from .base import BaseResource


class ReportResource(BaseResource):
    """Reports: Scheduled / on-demand security reports."""

    def list(self, params=None):
        """GET /business/report - List reports."""
        return self._http.get(f"/business/report", params)

    def create(self, data=None):
        """POST /business/report - Create a report."""
        return self._http.post(f"/business/report", data)

    def get(self, id, params=None):
        """GET /business/report/{id} - Get a report."""
        return self._http.get(f"/business/report/{id}", params)

    def delete(self, id, data=None):
        """DELETE /business/report/{id} - Delete a report."""
        return self._http.delete(f"/business/report/{id}", data)
