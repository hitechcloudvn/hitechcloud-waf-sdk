import { WafClient, WafError } from "@hitechcloudvn/waf-sdk";

async function main() {
  const client = new WafClient({ username: "admin", password: "********" });
  await client.login("admin", "********"); // CSRF -> login -> JWT

  try {
    const sites = await client.sites.list({ page: 1, page_size: 20 });
    console.log("Sites:", sites);

    await client.acl.create("123", { type: "ip", value: "203.0.113.10", action: "deny" });

    console.log("QPS:", await client.statistics.qps());
    console.log("License:", await client.license.status());
  } catch (e) {
    if (e instanceof WafError) {
      console.error(`API error ${e.statusCode}: ${e.message}`);
    } else {
      throw e;
    }
  }
}

main();
