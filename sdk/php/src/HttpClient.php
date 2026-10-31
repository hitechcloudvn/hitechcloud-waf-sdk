<?php

namespace HiTechCloud\WAF;

use GuzzleHttp\Client;
use GuzzleHttp\Exception\RequestException;

class HttpClient
{
    private Client $client;
    private ?string $token = null;
    private int $maxRetries;

    public function __construct(string $baseUrl, int $timeout = 30, int $maxRetries = 3)
    {
        $this->maxRetries = $maxRetries;
        $this->client = new Client([
            'base_uri' => rtrim($baseUrl, '/') . '/api/',
            'timeout' => $timeout,
            'http_errors' => false,
            'headers' => [
                'Content-Type' => 'application/json',
                'User-Agent' => 'HiTechCloud-WAF-PHP/1.0.0',
            ],
        ]);
    }

    public function setToken(string $token): void
    {
        $this->token = $token;
    }

    public function clearToken(): void
    {
        $this->token = null;
    }

    private function headers(): array
    {
        $headers = ['Content-Type' => 'application/json'];
        if ($this->token) {
            $headers['Authorization'] = 'Bearer ' . $this->token;
        }
        return $headers;
    }

    public function request(string $method, string $path, ?array $data = null, array $params = []): ?array
    {
        $path = ltrim($path, '/');
        for ($attempt = 0; $attempt <= $this->maxRetries; $attempt++) {
            $options = ['headers' => $this->headers()];
            if ($data !== null) {
                $options['json'] = $data;
            }
            if (!empty($params)) {
                $options['query'] = $params;
            }

            $response = $this->client->request($method, $path, $options);
            $status = $response->getStatusCode();

            if ($status === 429 && $attempt < $this->maxRetries) {
                sleep((int) ($response->getHeader('Retry-After')[0] ?? 60));
                continue;
            }
            if ($status >= 500 && $attempt < $this->maxRetries) {
                sleep((int) pow(2, $attempt));
                continue;
            }

            $body = json_decode((string) $response->getBody(), true) ?: [];

            if ($status >= 400) {
                $this->raise($status, $body, $response);
            }
            return $body;
        }
        return null;
    }

    private function raise(int $status, array $body, $response): void
    {
        $message = $body['msg'] ?? $body['err'] ?? "HTTP {$status}";
        switch ($status) {
            case 401:
                throw new AuthenticationException($message, $body);
            case 403:
                throw new AuthorizationException($message, $body);
            case 404:
                throw new NotFoundException($message, $body);
            case 422:
                throw new ValidationException($message, $body);
            case 429:
                $retryAfter = (int) ($response->getHeader('Retry-After')[0] ?? 60);
                throw new RateLimitException($message, $retryAfter, $body);
            default:
                if ($status >= 500) {
                    throw new ServerException($message, $status, $body);
                }
                throw new WafException($message, $status, $body);
        }
    }

    public function get(string $path, array $params = []): ?array
    {
        return $this->request('GET', $path, null, $params);
    }

    public function post(string $path, ?array $data = null): ?array
    {
        return $this->request('POST', $path, $data ?? []);
    }

    public function put(string $path, ?array $data = null): ?array
    {
        return $this->request('PUT', $path, $data ?? []);
    }

    public function delete(string $path, ?array $data = null): ?array
    {
        return $this->request('DELETE', $path, $data);
    }
}
