package vn.hitechcloud.waf;

public class WafException extends RuntimeException {
    private final int statusCode;
    private final Object data;

    public WafException(String message, int statusCode, Object data) {
        super(message);
        this.statusCode = statusCode;
        this.data = data;
    }

    public int getStatusCode() { return statusCode; }
    public Object getData() { return data; }
}
