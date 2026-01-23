package schwab

import "fmt"

// APIError represents an error returned by the Schwab API.
type APIError struct {
	// StatusCode is the HTTP status code of the response.
	StatusCode int `json:"-"`
	// Message is the error message.
	Message string `json:"message,omitempty"`
	// Errors is a list of error details.
	Errors []string `json:"errors,omitempty"`
	// SchwabClientCorrelId is the correlation ID for the request, useful for debugging with Schwab support.
	SchwabClientCorrelId string `json:"-"`
}

func (e *APIError) Error() string {
	if e.SchwabClientCorrelId != "" {
		return fmt.Sprintf("schwab api error (status %d, correl-id %s): %s %v", e.StatusCode, e.SchwabClientCorrelId, e.Message, e.Errors)
	}
	return fmt.Sprintf("schwab api error (status %d): %s %v", e.StatusCode, e.Message, e.Errors)
}
