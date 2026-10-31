<?php

namespace HiTechCloud\WAF;

class WafException extends \Exception
{
    protected $data;

    public function __construct(string $message = "", int $code = 0, $data = null)
    {
        parent::__construct($message, $code);
        $this->data = $data;
    }

    public function getData()
    {
        return $this->data;
    }
}

class AuthenticationException extends WafException
{
    public function __construct(string $message = "Authentication failed", $data = null)
    {
        parent::__construct($message, 401, $data);
    }
}

class AuthorizationException extends WafException
{
    public function __construct(string $message = "Access denied", $data = null)
    {
        parent::__construct($message, 403, $data);
    }
}

class NotFoundException extends WafException
{
    public function __construct(string $message = "Resource not found", $data = null)
    {
        parent::__construct($message, 404, $data);
    }
}

class ValidationException extends WafException
{
    public function __construct(string $message = "Validation error", $data = null)
    {
        parent::__construct($message, 422, $data);
    }
}

class RateLimitException extends WafException
{
    private int $retryAfter;

    public function __construct(string $message = "Rate limit exceeded", int $retryAfter = 60, $data = null)
    {
        parent::__construct($message, 429, $data);
        $this->retryAfter = $retryAfter;
    }

    public function getRetryAfter(): int
    {
        return $this->retryAfter;
    }
}

class ServerException extends WafException
{
    public function __construct(string $message = "Server error", int $code = 500, $data = null)
    {
        parent::__construct($message, $code, $data);
    }
}
