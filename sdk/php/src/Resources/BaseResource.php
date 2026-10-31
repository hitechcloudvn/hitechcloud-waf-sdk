<?php

namespace HiTechCloud\WAF\Resources;

use HiTechCloud\WAF\HttpClient;

class BaseResource
{
    protected HttpClient $http;

    public function __construct(HttpClient $http)
    {
        $this->http = $http;
    }
}
