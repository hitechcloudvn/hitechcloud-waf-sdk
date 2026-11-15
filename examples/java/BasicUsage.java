// HiTechCloud WAF SDK — Java example / Ví dụ Java
//
// Build with the vn.hitechcloud:waf-sdk dependency, then: / Biên dịch với dependency vn.hitechcloud:waf-sdk, sau đó:
//   java BasicUsage.java   (Java 11+ single-file)

import vn.hitechcloud.waf.WafClient;
import java.util.HashMap;
import java.util.Map;

public class BasicUsage {
    public static void main(String[] args) {
        WafClient client = new WafClient(null);
        client.login("admin", "********", null); // CSRF -> login -> JWT / CSRF -> đăng nhập -> JWT

        // List sites / Liệt kê site
        Map<String, String> query = new HashMap<>();
        query.put("page", "1");
        query.put("page_size", "20");
        System.out.println("Sites: " + client.sites().list(query));

        // Create an ACL rule / Tạo quy tắc ACL
        Map<String, Object> rule = new HashMap<>();
        rule.put("type", "ip");
        rule.put("value", "203.0.113.10");
        rule.put("action", "deny");
        client.acl().create("123", rule);

        // Read statistics and license / Đọc thống kê và giấy phép
        System.out.println("QPS: " + client.statistics().qps(null));
        System.out.println("License: " + client.license().status(null));
    }
}
