package waf

import "fmt"

// WafError is the base error type.
type WafError struct {
	StatusCode int
	Message    string
	Data       map[string]interface{}
}

func (e *WafError) Error() string {
	return fmt.Sprintf("HiTechCloud WAF error %d: %s", e.StatusCode, e.Message)
}

// AuthenticationError represents 401 errors.
type AuthenticationError struct{ WafError }

// AuthorizationError represents 403 errors.
type AuthorizationError struct{ WafError }

// NotFoundError represents 404 errors.
type NotFoundError struct{ WafError }

// ValidationError represents 422 errors.
type ValidationError struct{ WafError }

// RateLimitError represents 429 errors.
type RateLimitError struct {
	WafError
	RetryAfter int
}

// ServerError represents 5xx errors.
type ServerError struct{ WafError }

func mapError(statusCode int, data map[string]interface{}, retryAfter int) error {
	message := "API error"
	if v, ok := data["msg"].(string); ok && v != "" {
		message = v
	} else if v, ok := data["err"].(string); ok && v != "" {
		message = v
	}
	base := WafError{StatusCode: statusCode, Message: message, Data: data}
	switch statusCode {
	case 401:
		return &AuthenticationError{base}
	case 403:
		return &AuthorizationError{base}
	case 404:
		return &NotFoundError{base}
	case 422:
		return &ValidationError{base}
	case 429:
		return &RateLimitError{WafError: base, RetryAfter: retryAfter}
	default:
		if statusCode >= 500 {
			return &ServerError{base}
		}
		return &base
	}
}
