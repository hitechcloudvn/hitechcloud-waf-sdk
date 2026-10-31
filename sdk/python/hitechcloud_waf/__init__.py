"""
HiTechCloud WAF SDK for Python

Official SDK for the HiTechCloud WAF management API (https://vcwaf.hitechcloud.vn/api).
227 endpoints across 25 resource groups.

    from hitechcloud_waf import WafClient

    client = WafClient(username="admin", password="********")
    sites = client.sites.list()
"""

__version__ = "1.0.0"
__author__ = "HiTechCloud Vietnam"
__email__ = "support@hitechcloud.vn"

from .client import WafClient
from .exceptions import (
    WafError,
    AuthenticationError,
    AuthorizationError,
    NotFoundError,
    RateLimitError,
    ValidationError,
    ServerError,
)

__all__ = [
    "WafClient",
    "WafError",
    "AuthenticationError",
    "AuthorizationError",
    "NotFoundError",
    "RateLimitError",
    "ValidationError",
    "ServerError",
]
