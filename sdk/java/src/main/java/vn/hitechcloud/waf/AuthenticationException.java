package vn.hitechcloud.waf;

public class AuthenticationException extends WafException {
    public AuthenticationException(String message, int statusCode, Object data) {
        super(message, statusCode, data);
    }
}
