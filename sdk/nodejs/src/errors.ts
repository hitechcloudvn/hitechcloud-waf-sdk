export class WafError extends Error {
  statusCode?: number;
  data?: unknown;
  constructor(message: string, statusCode?: number, data?: unknown) {
    super(message);
    this.name = "WafError";
    this.statusCode = statusCode;
    this.data = data;
  }
}

export class AuthenticationError extends WafError {
  constructor(message = "Authentication failed", data?: unknown) {
    super(message, 401, data);
    this.name = "AuthenticationError";
  }
}

export class AuthorizationError extends WafError {
  constructor(message = "Access denied", data?: unknown) {
    super(message, 403, data);
    this.name = "AuthorizationError";
  }
}

export class NotFoundError extends WafError {
  constructor(message = "Resource not found", data?: unknown) {
    super(message, 404, data);
    this.name = "NotFoundError";
  }
}

export class ValidationError extends WafError {
  constructor(message = "Validation error", data?: unknown) {
    super(message, 422, data);
    this.name = "ValidationError";
  }
}

export class RateLimitError extends WafError {
  retryAfter: number;
  constructor(message = "Rate limit exceeded", retryAfter = 60, data?: unknown) {
    super(message, 429, data);
    this.name = "RateLimitError";
    this.retryAfter = retryAfter;
  }
}

export class ServerError extends WafError {
  constructor(message = "Server error", statusCode = 500, data?: unknown) {
    super(message, statusCode, data);
    this.name = "ServerError";
  }
}
