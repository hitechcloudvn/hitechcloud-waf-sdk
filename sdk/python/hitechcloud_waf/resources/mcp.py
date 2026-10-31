"""MCP resource (2 endpoints).

Model Context Protocol bridge endpoints
"""

from .base import BaseResource


class McpResource(BaseResource):
    """MCP: Model Context Protocol bridge endpoints."""

    def get(self, params=None):
        """GET /open/mcp - Get MCP capability metadata."""
        return self._http.get(f"/open/mcp", params)

    def call(self, data=None):
        """POST /open/mcp - Invoke an MCP action."""
        return self._http.post(f"/open/mcp", data)
