"""HiTechCloud WAF SDK - Python example."""

from hitechcloud_waf import WafClient, WafError

client = WafClient(username="admin", password="********")  # CSRF -> login -> JWT

try:
    sites = client.sites.list(params={"page": 1, "page_size": 20})
    print("Sites:", sites)

    # Create an ACL rule on site 123
    client.acl.create("123", data={"type": "ip", "value": "203.0.113.10", "action": "deny"})

    # Read statistics
    print("QPS:", client.statistics.qps())
    print("Attacks:", client.statistics.attack(params={"range": "24h"}))

    # License
    print("License:", client.license.status())
except WafError as e:
    print(f"API error {e.status_code}: {e.message}")
