package vn.hitechcloud.waf.resources;

import vn.hitechcloud.waf.HttpClient;

public class BaseResource {
    protected final HttpClient http;

    public BaseResource(HttpClient http) {
        this.http = http;
    }
}
