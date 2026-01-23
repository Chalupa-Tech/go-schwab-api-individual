package schwab

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Error represents an API error response.
type Error struct {
	StatusCode int      `json:"-"`
	Message    string   `json:"message"`
	Errors     []string `json:"errors,omitempty"`
	CorrelID   string   `json:"-"`
}

func (e *Error) Error() string {
	msg := fmt.Sprintf("schwab api error (status %d): %s", e.StatusCode, e.Message)
	if len(e.Errors) > 0 {
		msg += fmt.Sprintf(" details: %v", e.Errors)
	}
	if e.CorrelID != "" {
		msg += fmt.Sprintf(" (CorrelID: %s)", e.CorrelID)
	}
	return msg
}

// CheckResponse checks the API response for errors and returns them.
func CheckResponse(resp *http.Response) error {
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return nil
	}

	apiErr := &Error{
		StatusCode: resp.StatusCode,
		CorrelID:   resp.Header.Get("Schwab-Client-CorrelId"),
	}

	data, err := io.ReadAll(resp.Body)
	if err == nil && len(data) > 0 {
		if err := json.Unmarshal(data, apiErr); err != nil {
			// If we can't unmarshal the error body, use the raw body as the message
			apiErr.Message = string(data)
		}
	} else {
		apiErr.Message = http.StatusText(resp.StatusCode)
	}

	if apiErr.Message == "" {
		apiErr.Message = "unknown error"
	}

	return apiErr
}
