# HiTechCloud WAF SDK (Python)

> [English](README.md) | **Tiếng Việt**

---

SDK Python chính thức cho API quản trị HiTechCloud WAF.

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

Dùng môi trường thử nghiệm với `WafClient(sandbox=True, ...)`.
