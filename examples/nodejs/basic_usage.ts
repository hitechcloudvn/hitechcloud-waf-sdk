// HiTechCloud WAF SDK — Node.js example / Ví dụ Node.js
//
// Install the SDK from sdk/nodejs first, then run: / Cài đặt SDK từ sdk/nodejs trước, sau đó chạy:
//   npx tsx basic_usage.ts
import { WafClient, WafError } from "@hitechcloudvn/waf-sdk";

async function main() {
  const client = new WafClient({ username: "admin", password: "********" });
  await client.login("admin", "********"); // CSRF -> login -> JWT / CSRF -> đăng nhập -> JWT

  try {
    // List sites / Liệt kê site
    const sites = await client.sites.list({ page: 1, page_size: 20 });
    console.log("Sites:", sites);

    // Create an ACL rule / Tạo quy tắc ACL
    await client.acl.create("123", { type: "ip", value: "203.0.113.10", action: "deny" });

    // Read statistics and license / Đọc thống kê và giấy phép
    console.log("QPS:", await client.statistics.qps());
    console.log("License:", await client.license.status());
  } catch (e) {
    // Error handling / Xử lý lỗi
    if (e instanceof WafError) {
      console.error(`API error ${e.statusCode}: ${e.message}`);
    } else {
      throw e;
    }
  }
}

main();
