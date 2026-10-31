# HiTechCloud WAF SDK (Node.js / TypeScript)

```bash
npm install @hitechcloud-vietnam/waf-sdk
```

```ts
import { WafClient } from "@hitechcloud-vietnam/waf-sdk";

const client = new WafClient({ username: "admin", password: "********" });
await client.login("admin", "********"); // CSRF -> login -> JWT

const sites = await client.sites.list();
const site = await client.sites.get("123");
await client.policies.create({ name: "block-sqli" });
await client.statistics.qps();
```

Pass `{ sandbox: true }` to target `https://vcwaf-sandbox.hitechcloud.vn`.
