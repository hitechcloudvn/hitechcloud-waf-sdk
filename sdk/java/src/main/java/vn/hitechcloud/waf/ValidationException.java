package vn.hitechcloud.waf;

public class ValidationException extends WafException {
    public ValidationException(String message, int statusCode, Object data) {
        super(message, statusCode, data);
    }
}
