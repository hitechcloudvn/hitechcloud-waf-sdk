"""Authentication & Token resource (7 endpoints).

CSRF, login, TFA, logout and API token lifecycle
"""

from .base import BaseResource


class AuthResource(BaseResource):
    """Authentication & Token: CSRF, login, TFA, logout and API token lifecycle."""

    def csrf(self, params=None):
        """GET /open/auth/csrf - Get a CSRF token for the login step."""
        return self._http.get(f"/open/auth/csrf", params)

    def login(self, data=None):
        """POST /open/auth/login - Authenticate with username/password (+csrf_token) and obtain a JWT."""
        return self._http.post(f"/open/auth/login", data)

    def logout(self, data=None):
        """POST /open/auth/logout - End the current session."""
        return self._http.post(f"/open/auth/logout", data)

    def tfa(self, data=None):
        """POST /open/auth/tfa - Verify a two-factor (OTP) code during login."""
        return self._http.post(f"/open/auth/tfa", data)

    def get_token(self, params=None):
        """GET /open/auth/token - Get the current API token."""
        return self._http.get(f"/open/auth/token", params)

    def refresh_token(self, data=None):
        """PUT /open/auth/token - Rotate / refresh the API token."""
        return self._http.put(f"/open/auth/token", data)

    def revoke_token(self, data=None):
        """DELETE /open/auth/token - Revoke the API token."""
        return self._http.delete(f"/open/auth/token", data)
