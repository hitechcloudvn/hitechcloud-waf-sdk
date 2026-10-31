"""HiTechCloud WAF SDK exceptions."""


class WafError(Exception):
    """Base exception for the WAF SDK."""

    def __init__(self, message, status_code=None, data=None):
        self.message = message
        self.status_code = status_code
        self.data = data or {}
        super().__init__(self.message)


class AuthenticationError(WafError):
    """401 - authentication failed."""


class AuthorizationError(WafError):
    """403 - access forbidden."""


class NotFoundError(WafError):
    """404 - resource not found."""


class ValidationError(WafError):
    """422 - request validation failed."""


class RateLimitError(WafError):
    """429 - rate limit exceeded."""

    def __init__(self, message, retry_after=None, **kwargs):
        super().__init__(message, **kwargs)
        self.retry_after = retry_after


class ServerError(WafError):
    """5xx - server error."""


ERROR_MAP = {
    401: AuthenticationError,
    403: AuthorizationError,
    404: NotFoundError,
    422: ValidationError,
    429: RateLimitError,
    500: ServerError,
    502: ServerError,
    503: ServerError,
}
