# HiTechCloud WAF SDK (Node.js / TypeScript)

> [English](README.md) | **Tiếng Việt**

---

```bash
npm install @hitechcloudvn/waf-sdk
```

```ts
import { WafClient } from "@hitechcloudvn/waf-sdk";

const client = new WafClient({ username: "admin", password: "********" });
await client.login("admin", "********"); // CSRF -> login -> JWT

const sites = await client.sites.list();
const site = await client.sites.get("123");
await client.policies.create({ name: "block-sqli" });
await client.statistics.qps();
```

Truyền `{ sandbox: true }` để trỏ tới `https://vcwaf-sandbox.hitechcloud.vn`.
