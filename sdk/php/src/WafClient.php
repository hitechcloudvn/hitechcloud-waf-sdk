<?php

namespace HiTechCloud\WAF;

class WafClient
{
    const PRODUCTION_URL = 'https://vcwaf.hitechcloud.vn';
    const SANDBOX_URL = 'https://vcwaf-sandbox.hitechcloud.vn';

    private HttpClient $http;

    public Resources\Auth $auth;
    public Resources\Users $users;
    public Resources\AuthDefense $authDefense;
    public Resources\Sites $sites;
    public Resources\Policies $policies;
    public Resources\Skynet $skynet;
    public Resources\AntiTamper $antiTamper;
    public Resources\Acl $acl;
    public Resources\Challenge $challenge;
    public Resources\WaitingRoom $waitingRoom;
    public Resources\Ja4 $ja4;
    public Resources\GlobalSettings $globalSettings;
    public Resources\Cloud $cloud;
    public Resources\Ssl $ssl;
    public Resources\Events $events;
    public Resources\SecurityPosture $securityPosture;
    public Resources\Statistics $statistics;
    public Resources\Report $report;
    public Resources\Alarm $alarm;
    public Resources\Portal $portal;
    public Resources\AntiBot $antiBot;
    public Resources\Settings $settings;
    public Resources\System $system;
    public Resources\License $license;
    public Resources\Mcp $mcp;

    public function __construct(array $config = [])
    {
        $baseUrl = $config['base_url']
            ?? (($config['sandbox'] ?? false) ? self::SANDBOX_URL : self::PRODUCTION_URL);
        $this->http = new HttpClient(
            $baseUrl,
            $config['timeout'] ?? 30,
            $config['max_retries'] ?? 3
        );

        $this->auth = new Resources\Auth($this->http);
        $this->users = new Resources\Users($this->http);
        $this->authDefense = new Resources\AuthDefense($this->http);
        $this->sites = new Resources\Sites($this->http);
        $this->policies = new Resources\Policies($this->http);
        $this->skynet = new Resources\Skynet($this->http);
        $this->antiTamper = new Resources\AntiTamper($this->http);
        $this->acl = new Resources\Acl($this->http);
        $this->challenge = new Resources\Challenge($this->http);
        $this->waitingRoom = new Resources\WaitingRoom($this->http);
        $this->ja4 = new Resources\Ja4($this->http);
        $this->globalSettings = new Resources\GlobalSettings($this->http);
        $this->cloud = new Resources\Cloud($this->http);
        $this->ssl = new Resources\Ssl($this->http);
        $this->events = new Resources\Events($this->http);
        $this->securityPosture = new Resources\SecurityPosture($this->http);
        $this->statistics = new Resources\Statistics($this->http);
        $this->report = new Resources\Report($this->http);
        $this->alarm = new Resources\Alarm($this->http);
        $this->portal = new Resources\Portal($this->http);
        $this->antiBot = new Resources\AntiBot($this->http);
        $this->settings = new Resources\Settings($this->http);
        $this->system = new Resources\System($this->http);
        $this->license = new Resources\License($this->http);
        $this->mcp = new Resources\Mcp($this->http);

        if (!empty($config['token'])) {
            $this->http->setToken($config['token']);
        } elseif (!empty($config['username']) && !empty($config['password'])) {
            $this->login($config['username'], $config['password'], $config['tfa_code'] ?? null);
        }
    }

    /** Run the full CSRF -> login (-> TFA) flow and store the JWT. */
    public function login(string $username, string $password, ?string $tfaCode = null): ?array
    {
        $csrf = $this->auth->csrf() ?? [];
        $csrfToken = $this->dig($csrf, 'csrf_token') ?? $this->dig($csrf, 'token');
        $payload = ['username' => $username, 'password' => $password];
        if ($csrfToken) {
            $payload['csrf_token'] = $csrfToken;
        }
        $result = $this->auth->login($payload) ?? [];
        $token = $this->dig($result, 'jwt') ?? $this->dig($result, 'token');
        if ($tfaCode && !$token) {
            $result = $this->auth->tfa(['code' => $tfaCode, 'username' => $username]) ?? [];
            $token = $this->dig($result, 'jwt') ?? $this->dig($result, 'token');
        }
        if ($token) {
            $this->http->setToken($token);
        }
        return $result;
    }

    public function setToken(string $token): void
    {
        $this->http->setToken($token);
    }

    public function logout(): ?array
    {
        $result = $this->auth->logout();
        $this->http->clearToken();
        return $result;
    }

    private function dig(array $obj, string $key)
    {
        if (array_key_exists($key, $obj)) {
            return $obj[$key];
        }
        if (isset($obj['data']) && is_array($obj['data']) && array_key_exists($key, $obj['data'])) {
            return $obj['data'][$key];
        }
        return null;
    }
}
