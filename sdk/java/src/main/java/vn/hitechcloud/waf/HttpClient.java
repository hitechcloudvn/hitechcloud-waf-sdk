package vn.hitechcloud.waf;

import com.google.gson.Gson;
import com.google.gson.reflect.TypeToken;

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStream;
import java.io.InputStreamReader;
import java.io.OutputStream;
import java.lang.reflect.Type;
import java.net.HttpURLConnection;
import java.net.URL;
import java.net.URLEncoder;
import java.nio.charset.StandardCharsets;
import java.util.Map;

public class HttpClient {
    private final String baseUrl;
    private final int timeout;
    private final int maxRetries;
    private String token;
    private final Gson gson = new Gson();

    public HttpClient(String baseUrl, int timeout, int maxRetries) {
        this.baseUrl = baseUrl.replaceAll("/+$", "") + "/api";
        this.timeout = timeout;
        this.maxRetries = maxRetries;
    }

    public void setToken(String token) { this.token = token; }
    public void clearToken() { this.token = null; }

    public Map<String, Object> request(String method, String path, Object body, Map<String, String> params) throws WafException {
        Exception lastError = null;
        for (int attempt = 0; attempt <= maxRetries; attempt++) {
            try {
                StringBuilder urlStr = new StringBuilder(baseUrl).append(path);
                if (params != null && !params.isEmpty()) {
                    boolean first = true;
                    for (Map.Entry<String, String> e : params.entrySet()) {
                        urlStr.append(first ? "?" : "&");
                        urlStr.append(URLEncoder.encode(e.getKey(), StandardCharsets.UTF_8))
                              .append("=")
                              .append(URLEncoder.encode(e.getValue(), StandardCharsets.UTF_8));
                        first = false;
                    }
                }

                HttpURLConnection conn = (HttpURLConnection) new URL(urlStr.toString()).openConnection();
                conn.setRequestMethod(method);
                conn.setConnectTimeout(timeout * 1000);
                conn.setReadTimeout(timeout * 1000);
                conn.setRequestProperty("Content-Type", "application/json");
                conn.setRequestProperty("User-Agent", "HiTechCloud-WAF-Java/" + WafClient.VERSION);
                if (token != null) {
                    conn.setRequestProperty("Authorization", "Bearer " + token);
                }

                if (body != null && ("POST".equals(method) || "PUT".equals(method) || "DELETE".equals(method))) {
                    conn.setDoOutput(true);
                    try (OutputStream os = conn.getOutputStream()) {
                        os.write(gson.toJson(body).getBytes(StandardCharsets.UTF_8));
                    }
                }

                int statusCode = conn.getResponseCode();

                if (statusCode == 429 && attempt < maxRetries) {
                    String ra = conn.getHeaderField("Retry-After");
                    int seconds = ra != null ? Integer.parseInt(ra) : 60;
                    Thread.sleep(seconds * 1000L);
                    continue;
                }
                if (statusCode >= 500 && attempt < maxRetries) {
                    Thread.sleep((long) Math.pow(2, attempt) * 1000);
                    continue;
                }

                InputStream is = statusCode >= 400 ? conn.getErrorStream() : conn.getInputStream();
                String responseBody = readStream(is);
                Type mapType = new TypeToken<Map<String, Object>>() {}.getType();
                Map<String, Object> result = responseBody.isEmpty() ? null : gson.fromJson(responseBody, mapType);

                if (statusCode >= 400) {
                    int retryAfter = 60;
                    String ra = conn.getHeaderField("Retry-After");
                    if (ra != null) {
                        try { retryAfter = Integer.parseInt(ra); } catch (NumberFormatException ignored) {}
                    }
                    throw mapError(statusCode, result, retryAfter);
                }
                return result;
            } catch (WafException e) {
                throw e;
            } catch (Exception e) {
                lastError = e;
                if (attempt < maxRetries) {
                    try { Thread.sleep((long) Math.pow(2, attempt) * 1000); }
                    catch (InterruptedException ie) { Thread.currentThread().interrupt(); }
                }
            }
        }
        throw new WafException("Request failed: " + (lastError != null ? lastError.getMessage() : "unknown"), 0, null);
    }

    public Map<String, Object> get(String path, Map<String, String> params) throws WafException {
        return request("GET", path, null, params);
    }
    public Map<String, Object> post(String path, Object body) throws WafException {
        return request("POST", path, body, null);
    }
    public Map<String, Object> put(String path, Object body) throws WafException {
        return request("PUT", path, body, null);
    }
    public Map<String, Object> delete(String path, Object body) throws WafException {
        return request("DELETE", path, body, null);
    }

    private String readStream(InputStream is) throws IOException {
        if (is == null) return "";
        StringBuilder sb = new StringBuilder();
        try (BufferedReader reader = new BufferedReader(new InputStreamReader(is, StandardCharsets.UTF_8))) {
            String line;
            while ((line = reader.readLine()) != null) {
                sb.append(line);
            }
        }
        return sb.toString();
    }

    private WafException mapError(int statusCode, Map<String, Object> data, int retryAfter) {
        String message = "API error";
        if (data != null) {
            if (data.get("msg") != null) message = data.get("msg").toString();
            else if (data.get("err") != null) message = data.get("err").toString();
        }
        switch (statusCode) {
            case 401: return new AuthenticationException(message, statusCode, data);
            case 403: return new AuthorizationException(message, statusCode, data);
            case 404: return new NotFoundException(message, statusCode, data);
            case 422: return new ValidationException(message, statusCode, data);
            case 429: return new RateLimitException(message, statusCode, data, retryAfter);
            default:
                if (statusCode >= 500) return new ServerException(message, statusCode, data);
                return new WafException(message, statusCode, data);
        }
    }
}
