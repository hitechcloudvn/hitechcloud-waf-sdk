"""HiTechCloud WAF SDK — Python example. / Ví dụ Python.

Run: / Cách chạy:
    pip install -e ../../sdk/python   (or pip install hitechcloud-waf / hoặc pip install hitechcloud-waf)
    python basic_usage.py
"""

from hitechcloud_waf import WafClient, WafError

client = WafClient(username="admin", password="********")  # CSRF -> login -> JWT / CSRF -> đăng nhập -> JWT

try:
    # List sites / Liệt kê site
    sites = client.sites.list(params={"page": 1, "page_size": 20})
    print("Sites:", sites)

    # Create an ACL rule on site 123 / Tạo quy tắc ACL cho site 123
    client.acl.create("123", data={"type": "ip", "value": "203.0.113.10", "action": "deny"})

    # Read statistics / Đọc thống kê
    print("QPS:", client.statistics.qps())
    print("Attacks:", client.statistics.attack(params={"range": "24h"}))

    # License / Giấy phép
    print("License:", client.license.status())
except WafError as e:
    # Error handling / Xử lý lỗi
    print(f"API error {e.status_code}: {e.message}")
