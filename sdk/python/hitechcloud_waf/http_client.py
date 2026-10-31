"""HiTechCloud WAF SDK HTTP client."""

import json
import time
from urllib.parse import urljoin

import requests

from .exceptions import WafError, RateLimitError, ERROR_MAP


class HTTPClient:
    """Low-level HTTP client. The API base path (/api) is appended here."""

    def __init__(self, base_url="https://vcwaf.hitechcloud.vn", timeout=30, max_retries=3):
        self.base_url = base_url.rstrip("/") + "/api"
        self.timeout = timeout
        self.max_retries = max_retries
        self.session = requests.Session()
        self.session.headers.update({
            "Content-Type": "application/json",
            "Accept": "application/json",
            "User-Agent": "HiTechCloud-WAF-Python/1.0.0",
        })

    def set_token(self, token):
        self.session.headers["Authorization"] = f"Bearer {token}"

    def clear_token(self):
        self.session.headers.pop("Authorization", None)

    def request(self, method, path, params=None, json_data=None):
        url = self.base_url + "/" + path.lstrip("/")
        last_error = None
        for attempt in range(self.max_retries + 1):
            try:
                resp = self.session.request(
                    method=method, url=url, params=params,
                    json=json_data, timeout=self.timeout,
                )
                if resp.status_code == 429 and attempt < self.max_retries:
                    time.sleep(int(resp.headers.get("Retry-After", 60)))
                    continue
                if resp.status_code >= 500 and attempt < self.max_retries:
                    time.sleep(2 ** attempt)
                    continue
                if resp.status_code >= 400:
                    self._raise(resp)
                if resp.status_code == 204 or not resp.content:
                    return None
                return resp.json()
            except (requests.ConnectionError, requests.Timeout) as e:
                last_error = e
                if attempt < self.max_retries:
                    time.sleep(2 ** attempt)
                    continue
                raise WafError(f"Request failed: {e}")
        if last_error:
            raise WafError(f"Request failed: {last_error}")

    def _raise(self, resp):
        try:
            body = resp.json()
            message = body.get("msg") or body.get("err") or f"HTTP {resp.status_code}"
        except (json.JSONDecodeError, ValueError):
            body = {}
            message = f"HTTP {resp.status_code}: {resp.text[:200]}"
        cls = ERROR_MAP.get(resp.status_code, WafError)
        kwargs = {"message": message, "status_code": resp.status_code, "data": body}
        if resp.status_code == 429:
            kwargs["retry_after"] = int(resp.headers.get("Retry-After", 60))
        raise cls(**kwargs)

    def get(self, path, params=None):
        return self.request("GET", path, params=params)

    def post(self, path, data=None):
        return self.request("POST", path, json_data=data)

    def put(self, path, data=None):
        return self.request("PUT", path, json_data=data)

    def delete(self, path, data=None):
        return self.request("DELETE", path, json_data=data)

    def close(self):
        self.session.close()
