"""Attacks & Events resource (7 endpoints).

Attack events, matched rules, raw records and export
"""

from .base import BaseResource


class EventsResource(BaseResource):
    """Attacks & Events: Attack events, matched rules, raw records and export."""

    def list(self, params=None):
        """GET /open/events - List attack events."""
        return self._http.get(f"/open/events", params)

    def list_rules(self, params=None):
        """GET /open/events/rule - List events grouped by rule."""
        return self._http.get(f"/open/events/rule", params)

    def get_record(self, id, params=None):
        """GET /open/record/{id} - Get a single event record."""
        return self._http.get(f"/open/record/{id}", params)

    def get_record_rule(self, id, params=None):
        """GET /open/record/rule/{id} - Get the rule for a record."""
        return self._http.get(f"/open/record/rule/{id}", params)

    def records(self, params=None):
        """GET /open/records - List raw records."""
        return self._http.get(f"/open/records", params)

    def records_rule(self, params=None):
        """GET /open/records/rule - List raw records grouped by rule."""
        return self._http.get(f"/open/records/rule", params)

    def export(self, params=None):
        """GET /commercial/record/export - Export records (commercial)."""
        return self._http.get(f"/commercial/record/export", params)
