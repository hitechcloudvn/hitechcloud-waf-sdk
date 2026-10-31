package vn.hitechcloud.waf;

import vn.hitechcloud.waf.resources.*;

import java.util.HashMap;
import java.util.Map;

public class WafClient {
    public static final String PRODUCTION_URL = "https://vcwaf.hitechcloud.vn";
    public static final String SANDBOX_URL = "https://vcwaf-sandbox.hitechcloud.vn";
    public static final String VERSION = "1.0.0";

    private final HttpClient httpClient;
    private final Auth auth;
    private final Users users;
    private final AuthDefense authDefense;
    private final Sites sites;
    private final Policies policies;
    private final Skynet skynet;
    private final AntiTamper antiTamper;
    private final Acl acl;
    private final Challenge challenge;
    private final WaitingRoom waitingRoom;
    private final Ja4 ja4;
    private final GlobalSettings globalSettings;
    private final Cloud cloud;
    private final Ssl ssl;
    private final Events events;
    private final SecurityPosture securityPosture;
    private final Statistics statistics;
    private final Report report;
    private final Alarm alarm;
    private final Portal portal;
    private final AntiBot antiBot;
    private final Settings settings;
    private final System system;
    private final License license;
    private final Mcp mcp;

    public WafClient(String token) {
        this(token, PRODUCTION_URL, 30, 3);
    }

    public WafClient(String token, String baseUrl, int timeout, int maxRetries) {
        this.httpClient = new HttpClient(baseUrl, timeout, maxRetries);
        if (token != null) {
            this.httpClient.setToken(token);
        }
        this.auth = new Auth(httpClient);
        this.users = new Users(httpClient);
        this.authDefense = new AuthDefense(httpClient);
        this.sites = new Sites(httpClient);
        this.policies = new Policies(httpClient);
        this.skynet = new Skynet(httpClient);
        this.antiTamper = new AntiTamper(httpClient);
        this.acl = new Acl(httpClient);
        this.challenge = new Challenge(httpClient);
        this.waitingRoom = new WaitingRoom(httpClient);
        this.ja4 = new Ja4(httpClient);
        this.globalSettings = new GlobalSettings(httpClient);
        this.cloud = new Cloud(httpClient);
        this.ssl = new Ssl(httpClient);
        this.events = new Events(httpClient);
        this.securityPosture = new SecurityPosture(httpClient);
        this.statistics = new Statistics(httpClient);
        this.report = new Report(httpClient);
        this.alarm = new Alarm(httpClient);
        this.portal = new Portal(httpClient);
        this.antiBot = new AntiBot(httpClient);
        this.settings = new Settings(httpClient);
        this.system = new System(httpClient);
        this.license = new License(httpClient);
        this.mcp = new Mcp(httpClient);
    }

    /** Run the full CSRF -> login (-> TFA) flow and store the JWT. */
    public Map<String, Object> login(String username, String password, String tfaCode) {
        Map<String, Object> csrf = auth.csrf(null);
        Map<String, Object> payload = new HashMap<>();
        payload.put("username", username);
        payload.put("password", password);
        String csrfToken = dig(csrf, "csrf_token");
        if (csrfToken == null) csrfToken = dig(csrf, "token");
        if (csrfToken != null) payload.put("csrf_token", csrfToken);

        Map<String, Object> result = auth.login(payload);
        String token = dig(result, "jwt");
        if (token == null) token = dig(result, "token");

        if (tfaCode != null && token == null) {
            Map<String, Object> tfaBody = new HashMap<>();
            tfaBody.put("code", tfaCode);
            tfaBody.put("username", username);
            result = auth.tfa(tfaBody);
            token = dig(result, "jwt");
            if (token == null) token = dig(result, "token");
        }
        if (token != null) {
            httpClient.setToken(token);
        }
        return result;
    }

    public void setToken(String token) {
        httpClient.setToken(token);
    }

    public Map<String, Object> logout() {
        Map<String, Object> result = auth.logout(null);
        httpClient.clearToken();
        return result;
    }

    @SuppressWarnings("unchecked")
    private String dig(Map<String, Object> obj, String key) {
        if (obj == null) return null;
        if (obj.get(key) instanceof String) return (String) obj.get(key);
        if (obj.get("data") instanceof Map) {
            Object v = ((Map<String, Object>) obj.get("data")).get(key);
            if (v instanceof String) return (String) v;
        }
        return null;
    }

    public Auth auth() { return auth; }
    public Users users() { return users; }
    public AuthDefense authDefense() { return authDefense; }
    public Sites sites() { return sites; }
    public Policies policies() { return policies; }
    public Skynet skynet() { return skynet; }
    public AntiTamper antiTamper() { return antiTamper; }
    public Acl acl() { return acl; }
    public Challenge challenge() { return challenge; }
    public WaitingRoom waitingRoom() { return waitingRoom; }
    public Ja4 ja4() { return ja4; }
    public GlobalSettings globalSettings() { return globalSettings; }
    public Cloud cloud() { return cloud; }
    public Ssl ssl() { return ssl; }
    public Events events() { return events; }
    public SecurityPosture securityPosture() { return securityPosture; }
    public Statistics statistics() { return statistics; }
    public Report report() { return report; }
    public Alarm alarm() { return alarm; }
    public Portal portal() { return portal; }
    public AntiBot antiBot() { return antiBot; }
    public Settings settings() { return settings; }
    public System system() { return system; }
    public License license() { return license; }
    public Mcp mcp() { return mcp; }
}
