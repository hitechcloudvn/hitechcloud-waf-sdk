package vn.hitechcloud.waf;

public class NotFoundException extends WafException {
    public NotFoundException(String message, int statusCode, Object data) {
        super(message, statusCode, data);
    }
}
