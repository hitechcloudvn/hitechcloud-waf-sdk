package vn.hitechcloud.waf;

public class AuthorizationException extends WafException {
    public AuthorizationException(String message, int statusCode, Object data) {
        super(message, statusCode, data);
    }
}
