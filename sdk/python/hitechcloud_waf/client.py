"""HiTechCloud WAF SDK - main client."""

from .http_client import HTTPClient
from .resources.auth import AuthResource
from .resources.users import UsersResource
from .resources.auth_defense import AuthDefenseResource
from .resources.sites import SitesResource
from .resources.policies import PoliciesResource
from .resources.skynet import SkynetResource
from .resources.anti_tamper import AntiTamperResource
from .resources.acl import AclResource
from .resources.challenge import ChallengeResource
from .resources.waiting_room import WaitingRoomResource
from .resources.ja4 import Ja4Resource
from .resources.global_settings import GlobalSettingsResource
from .resources.cloud import CloudResource
from .resources.ssl import SslResource
from .resources.events import EventsResource
from .resources.security_posture import SecurityPostureResource
from .resources.statistics import StatisticsResource
from .resources.report import ReportResource
from .resources.alarm import AlarmResource
from .resources.portal import PortalResource
from .resources.anti_bot import AntiBotResource
from .resources.settings import SettingsResource
from .resources.system import SystemResource
from .resources.license import LicenseResource
from .resources.mcp import McpResource


class WafClient:
    """HiTechCloud WAF management client.

    Auth flow: GET /open/auth/csrf -> POST /open/auth/login (-> jwt) -> Bearer.

        client = WafClient(username="admin", password="********")
        client = WafClient(token="EXISTING_JWT")
        client = WafClient(sandbox=True, username="admin", password="********")
    """

    PRODUCTION_URL = "https://vcwaf.hitechcloud.vn"
    SANDBOX_URL = "https://vcwaf-sandbox.hitechcloud.vn"

    def __init__(self, base_url=None, username=None, password=None, token=None,
                 tfa_code=None, sandbox=False, timeout=30, max_retries=3):
        if base_url is None:
            base_url = self.SANDBOX_URL if sandbox else self.PRODUCTION_URL
        self._http = HTTPClient(base_url=base_url, timeout=timeout, max_retries=max_retries)

        self.auth = AuthResource(self._http)
        self.users = UsersResource(self._http)
        self.auth_defense = AuthDefenseResource(self._http)
        self.sites = SitesResource(self._http)
        self.policies = PoliciesResource(self._http)
        self.skynet = SkynetResource(self._http)
        self.anti_tamper = AntiTamperResource(self._http)
        self.acl = AclResource(self._http)
        self.challenge = ChallengeResource(self._http)
        self.waiting_room = WaitingRoomResource(self._http)
        self.ja_4 = Ja4Resource(self._http)
        self.global_settings = GlobalSettingsResource(self._http)
        self.cloud = CloudResource(self._http)
        self.ssl = SslResource(self._http)
        self.events = EventsResource(self._http)
        self.security_posture = SecurityPostureResource(self._http)
        self.statistics = StatisticsResource(self._http)
        self.report = ReportResource(self._http)
        self.alarm = AlarmResource(self._http)
        self.portal = PortalResource(self._http)
        self.anti_bot = AntiBotResource(self._http)
        self.settings = SettingsResource(self._http)
        self.system = SystemResource(self._http)
        self.license = LicenseResource(self._http)
        self.mcp = McpResource(self._http)

        if token:
            self._http.set_token(token)
        elif username and password:
            self.login(username, password, tfa_code=tfa_code)

    def login(self, username, password, tfa_code=None):
        """Run the full CSRF -> login (-> TFA) flow and store the JWT."""
        csrf = self.auth.csrf() or {}
        csrf_token = _dig(csrf, "csrf_token") or _dig(csrf, "token")
        payload = {"username": username, "password": password}
        if csrf_token:
            payload["csrf_token"] = csrf_token
        result = self.auth.login(data=payload) or {}
        token = _dig(result, "jwt") or _dig(result, "token")
        if tfa_code and not token:
            tfa = self.auth.tfa(data={"code": tfa_code, "username": username}) or {}
            token = _dig(tfa, "jwt") or _dig(tfa, "token")
            result = tfa
        if token:
            self._http.set_token(token)
        return result

    def set_token(self, token):
        self._http.set_token(token)

    def logout(self):
        result = self.auth.logout()
        self._http.clear_token()
        return result

    def __enter__(self):
        return self

    def __exit__(self, *args):
        self._http.close()


def _dig(obj, key):
    if not isinstance(obj, dict):
        return None
    if key in obj:
        return obj[key]
    data = obj.get("data")
    if isinstance(data, dict):
        return data.get(key)
    return None
