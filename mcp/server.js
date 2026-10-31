#!/usr/bin/env node
/**
 * HiTechCloud WAF MCP Server
 * Model Context Protocol server exposing core WAF management actions.
 *
 * Env:
 *   WAF_BASE_URL   (default https://vcwaf.hitechcloud.vn)
 *   WAF_TOKEN      JWT bearer token (obtain via the login flow)
 */
const { Server } = require("@modelcontextprotocol/sdk/server/index.js");
const { StdioServerTransport } = require("@modelcontextprotocol/sdk/server/stdio.js");

const BASE_URL = (process.env.WAF_BASE_URL || "https://vcwaf.hitechcloud.vn").replace(/\/+$/, "") + "/api";
const TOKEN = process.env.WAF_TOKEN;

async function apiCall(method, path, body) {
  const headers = {
    "Content-Type": "application/json",
    "User-Agent": "HiTechCloud-WAF-MCP/1.0.0",
  };
  if (TOKEN) headers["Authorization"] = `Bearer ${TOKEN}`;
  const opts = { method, headers };
  if (body) opts.body = JSON.stringify(body);
  const res = await fetch(`${BASE_URL}${path}`, opts);
  return res.json();
}

const TOOLS = [
  {
    "name": "waf_list_sites",
    "description": "List protected sites",
    "inputSchema": {
      "type": "object",
      "properties": {}
    }
  },
  {
    "name": "waf_get_site",
    "description": "Get a protected site by id",
    "inputSchema": {
      "type": "object",
      "properties": {
        "id": {
          "type": "string",
          "description": "id identifier"
        }
      },
      "required": [
        "id"
      ]
    }
  },
  {
    "name": "waf_list_policies",
    "description": "List protection policies",
    "inputSchema": {
      "type": "object",
      "properties": {}
    }
  },
  {
    "name": "waf_list_events",
    "description": "List recent attack events",
    "inputSchema": {
      "type": "object",
      "properties": {}
    }
  },
  {
    "name": "waf_qps",
    "description": "Get current QPS",
    "inputSchema": {
      "type": "object",
      "properties": {}
    }
  },
  {
    "name": "waf_attack_stats",
    "description": "Get attack statistics",
    "inputSchema": {
      "type": "object",
      "properties": {}
    }
  },
  {
    "name": "waf_security_posture",
    "description": "Get realtime security posture",
    "inputSchema": {
      "type": "object",
      "properties": {}
    }
  },
  {
    "name": "waf_list_certs",
    "description": "List SSL certificates",
    "inputSchema": {
      "type": "object",
      "properties": {}
    }
  },
  {
    "name": "waf_license_status",
    "description": "Get license status",
    "inputSchema": {
      "type": "object",
      "properties": {}
    }
  },
  {
    "name": "waf_acl_relieve",
    "description": "Unblock an ACL-blocked client",
    "inputSchema": {
      "type": "object",
      "properties": {
        "body": {
          "type": "object",
          "description": "Request body"
        }
      }
    }
  }
];

const HANDLERS = {
  waf_list_sites: (args) => apiCall('GET', `/open/site`, null),
  waf_get_site: (args) => apiCall('GET', `/open/site/${args.id}`, null),
  waf_list_policies: (args) => apiCall('GET', `/open/policy`, null),
  waf_list_events: (args) => apiCall('GET', `/open/events`, null),
  waf_qps: (args) => apiCall('GET', `/stat/qps`, null),
  waf_attack_stats: (args) => apiCall('GET', `/stat/advance/attack`, null),
  waf_security_posture: (args) => apiCall('GET', `/open/security_posture/realtime`, null),
  waf_list_certs: (args) => apiCall('GET', `/open/cert`, null),
  waf_license_status: (args) => apiCall('GET', `/open/system/license/status`, null),
  waf_acl_relieve: (args) => apiCall('PUT', `/acl/relieve`, args.body),
};

async function main() {
  const server = new Server(
    { name: "hitechcloud-waf", version: "1.0.0" },
    { capabilities: { tools: {} } }
  );

  server.setRequestHandler("tools/list", async () => ({ tools: TOOLS }));

  server.setRequestHandler("tools/call", async (request) => {
    const { name, arguments: args } = request.params;
    const handler = HANDLERS[name];
    if (!handler) throw new Error(`Unknown tool: ${name}`);
    const result = await handler(args || {});
    return { content: [{ type: "text", text: JSON.stringify(result, null, 2) }] };
  });

  const transport = new StdioServerTransport();
  await server.connect(transport);
  console.error("HiTechCloud WAF MCP server running on stdio");
}

main().catch((e) => {
  console.error(e);
  process.exit(1);
});
