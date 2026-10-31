<?php

namespace HiTechCloud\WAF\Resources;

/**
 * MCP: Model Context Protocol bridge endpoints.
 */
class Mcp extends BaseResource
{
    /** GET /open/mcp - Get MCP capability metadata */
    public function get(array $params = []): ?array
    {
        return $this->http->get('/open/mcp', $params);
    }

    /** POST /open/mcp - Invoke an MCP action */
    public function call(?array $data = null): ?array
    {
        return $this->http->post('/open/mcp', $data);
    }

}