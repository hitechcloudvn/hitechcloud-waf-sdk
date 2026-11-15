# HiTechCloud WAF SDK (Java)

> [English](README.md) | **Tiếng Việt**

---

```xml
<dependency>
    <groupId>vn.hitechcloud</groupId>
    <artifactId>waf-sdk</artifactId>
    <version>1.0.0</version>
</dependency>
```

```java
import vn.hitechcloud.waf.WafClient;
import java.util.Map;

WafClient client = new WafClient(null);   // then login:
client.login("admin", "********", null);  // CSRF -> login -> JWT
// or: new WafClient("EXISTING_JWT");

Map<String, Object> sites = client.sites().list(null);
Map<String, Object> site = client.sites().get("123", null);
client.statistics().qps(null);
```
