# HiTechCloud WAF SDK (Python)

> **English** | [Tiếng Việt](README.vi.md)

---

Official Python SDK for the HiTechCloud WAF management API.

```bash
pip install hitechcloud-waf
```

```python
from hitechcloud_waf import WafClient

client = WafClient(username="admin", password="********")   # CSRF -> login -> JWT

# Sites
sites = client.sites.list()
site = client.sites.get("123")

# Policies / ACL / Skynet
client.policies.create(data={"name": "block-sqli"})
client.acl.create("123", data={"type": "ip", "value": "1.2.3.4"})
client.skynet.set_rule_switch(data={"enabled": True})

# Statistics & events
client.statistics.qps()
client.events.list(params={"page": 1, "page_size": 50})
```

Use the sandbox with `WafClient(sandbox=True, ...)`.
