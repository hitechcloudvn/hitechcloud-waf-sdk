# EN: Smoke tests - the SDK package imports and exposes its public API.
# VI: Smoke test - gói SDK import được và khai báo API công khai.
from hitechcloud_waf import WafClient
from hitechcloud_waf.exceptions import WafError


def test_main_client_is_loadable():
    """EN: WafClient can be imported. / VI: Import được WafClient."""
    assert callable(WafClient)


def test_exception_hierarchy_is_loadable():
    """EN: WafError is the base exception. / VI: WafError là ngoại lệ gốc."""
    assert issubclass(WafError, Exception)
