import vn.hitechcloud.waf.WafClient;
import java.util.HashMap;
import java.util.Map;

public class BasicUsage {
    public static void main(String[] args) {
        WafClient client = new WafClient(null);
        client.login("admin", "********", null); // CSRF -> login -> JWT

        Map<String, String> query = new HashMap<>();
        query.put("page", "1");
        query.put("page_size", "20");
        System.out.println("Sites: " + client.sites().list(query));

        Map<String, Object> rule = new HashMap<>();
        rule.put("type", "ip");
        rule.put("value", "203.0.113.10");
        rule.put("action", "deny");
        client.acl().create("123", rule);

        System.out.println("QPS: " + client.statistics().qps(null));
        System.out.println("License: " + client.license().status(null));
    }
}
