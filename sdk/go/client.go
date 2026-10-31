package waf

import "time"

// Environment URLs and SDK version.
const (
	ProductionURL = "https://vcwaf.hitechcloud.vn"
	SandboxURL    = "https://vcwaf-sandbox.hitechcloud.vn"
	Version       = "1.0.0"
)

// Client is the main WAF management client.
type Client struct {
	httpClient *HTTPClient
	Auth *AuthResource
	Users *UsersResource
	AuthDefense *AuthDefenseResource
	Sites *SitesResource
	Policies *PoliciesResource
	Skynet *SkynetResource
	AntiTamper *AntiTamperResource
	Acl *AclResource
	Challenge *ChallengeResource
	WaitingRoom *WaitingRoomResource
	Ja4 *Ja4Resource
	GlobalSettings *GlobalSettingsResource
	Cloud *CloudResource
	Ssl *SslResource
	Events *EventsResource
	SecurityPosture *SecurityPostureResource
	Statistics *StatisticsResource
	Report *ReportResource
	Alarm *AlarmResource
	Portal *PortalResource
	AntiBot *AntiBotResource
	Settings *SettingsResource
	System *SystemResource
	License *LicenseResource
	Mcp *McpResource
}

// Option configures the client.
type Option func(*config)

type config struct {
	baseURL    string
	token      string
	username   string
	password   string
	tfaCode    string
	timeout    time.Duration
	maxRetries int
}

// WithToken sets an existing JWT.
func WithToken(token string) Option { return func(c *config) { c.token = token } }

// WithCredentials sets username/password for the login flow.
func WithCredentials(username, password string) Option {
	return func(c *config) { c.username = username; c.password = password }
}

// WithTFACode sets the two-factor OTP code.
func WithTFACode(code string) Option { return func(c *config) { c.tfaCode = code } }

// WithBaseURL sets a custom base URL.
func WithBaseURL(u string) Option { return func(c *config) { c.baseURL = u } }

// WithSandbox uses the sandbox environment.
func WithSandbox() Option { return func(c *config) { c.baseURL = SandboxURL } }

// WithTimeout sets the request timeout.
func WithTimeout(t time.Duration) Option { return func(c *config) { c.timeout = t } }

// WithMaxRetries sets the retry count.
func WithMaxRetries(n int) Option { return func(c *config) { c.maxRetries = n } }

// NewClient creates a new WAF client.
func NewClient(opts ...Option) (*Client, error) {
	cfg := &config{baseURL: ProductionURL, timeout: 30 * time.Second, maxRetries: 3}
	for _, opt := range opts {
		opt(cfg)
	}

	hc := &Client{}
	hc.httpClient = newHTTPClient(cfg.baseURL, cfg.timeout, cfg.maxRetries)

	hc.Auth = &AuthResource{client: hc.httpClient}
	hc.Users = &UsersResource{client: hc.httpClient}
	hc.AuthDefense = &AuthDefenseResource{client: hc.httpClient}
	hc.Sites = &SitesResource{client: hc.httpClient}
	hc.Policies = &PoliciesResource{client: hc.httpClient}
	hc.Skynet = &SkynetResource{client: hc.httpClient}
	hc.AntiTamper = &AntiTamperResource{client: hc.httpClient}
	hc.Acl = &AclResource{client: hc.httpClient}
	hc.Challenge = &ChallengeResource{client: hc.httpClient}
	hc.WaitingRoom = &WaitingRoomResource{client: hc.httpClient}
	hc.Ja4 = &Ja4Resource{client: hc.httpClient}
	hc.GlobalSettings = &GlobalSettingsResource{client: hc.httpClient}
	hc.Cloud = &CloudResource{client: hc.httpClient}
	hc.Ssl = &SslResource{client: hc.httpClient}
	hc.Events = &EventsResource{client: hc.httpClient}
	hc.SecurityPosture = &SecurityPostureResource{client: hc.httpClient}
	hc.Statistics = &StatisticsResource{client: hc.httpClient}
	hc.Report = &ReportResource{client: hc.httpClient}
	hc.Alarm = &AlarmResource{client: hc.httpClient}
	hc.Portal = &PortalResource{client: hc.httpClient}
	hc.AntiBot = &AntiBotResource{client: hc.httpClient}
	hc.Settings = &SettingsResource{client: hc.httpClient}
	hc.System = &SystemResource{client: hc.httpClient}
	hc.License = &LicenseResource{client: hc.httpClient}
	hc.Mcp = &McpResource{client: hc.httpClient}

	if cfg.token != "" {
		hc.httpClient.SetToken(cfg.token)
	} else if cfg.username != "" && cfg.password != "" {
		if _, err := hc.Login(cfg.username, cfg.password, cfg.tfaCode); err != nil {
			return nil, err
		}
	}
	return hc, nil
}

// Login runs the full CSRF -> login (-> TFA) flow and stores the JWT.
func (c *Client) Login(username, password, tfaCode string) (map[string]interface{}, error) {
	csrf, err := c.Auth.Csrf(nil)
	if err != nil {
		return nil, err
	}
	payload := map[string]interface{}{"username": username, "password": password}
	if t := dig(csrf, "csrf_token"); t != "" {
		payload["csrf_token"] = t
	} else if t := dig(csrf, "token"); t != "" {
		payload["csrf_token"] = t
	}
	result, err := c.Auth.Login(payload)
	if err != nil {
		return nil, err
	}
	token := dig(result, "jwt")
	if token == "" {
		token = dig(result, "token")
	}
	if tfaCode != "" && token == "" {
		result, err = c.Auth.Tfa(map[string]interface{}{"code": tfaCode, "username": username})
		if err != nil {
			return nil, err
		}
		token = dig(result, "jwt")
		if token == "" {
			token = dig(result, "token")
		}
	}
	if token != "" {
		c.httpClient.SetToken(token)
	}
	return result, nil
}

// SetToken sets the bearer token.
func (c *Client) SetToken(token string) { c.httpClient.SetToken(token) }

// Logout ends the session and clears the token.
func (c *Client) Logout() (map[string]interface{}, error) {
	result, err := c.Auth.Logout(nil)
	c.httpClient.ClearToken()
	return result, err
}

func dig(obj map[string]interface{}, key string) string {
	if obj == nil {
		return ""
	}
	if v, ok := obj[key].(string); ok {
		return v
	}
	if data, ok := obj["data"].(map[string]interface{}); ok {
		if v, ok := data[key].(string); ok {
			return v
		}
	}
	return ""
}
