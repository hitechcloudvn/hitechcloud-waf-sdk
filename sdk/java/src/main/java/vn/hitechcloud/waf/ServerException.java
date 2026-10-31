package vn.hitechcloud.waf;

public class ServerException extends WafException {
    public ServerException(String message, int statusCode, Object data) {
        super(message, statusCode, data);
    }
}
