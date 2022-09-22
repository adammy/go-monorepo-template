package httpapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/adammy/go-monorepo-template/pkg/pointer"
	"github.com/go-chi/chi/v5/middleware"
)

// SendJSON will set the appropriate headers and write the data argument to the response writer.
func SendJSON(w http.ResponseWriter, statusCode int, data interface{}) error {
	w.Header().Set(ContentTypeHeader, ApplicationJSON)
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		return fmt.Errorf("json encode failed: %w", err)
	}

	return nil
}

// SendErrorJSON will set the appropriate headers and write a structured Error to the response writer.
func SendErrorJSON(w http.ResponseWriter, statusCode int, err error) error {
	w.Header().Set(ContentTypeHeader, ApplicationJSON)
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(Error{
		Error:     err.Error(),
		Timestamp: time.Now(),
	}); err != nil {
		return fmt.Errorf("json encode failed: %w", err)
	}

	return nil
}

// SendErrorJSONWithRequest will set the appropriate headers and write a structured Error
// with a request ID to the response writer.
func SendErrorJSONWithRequest(w http.ResponseWriter, r *http.Request, statusCode int, err error) error {
	if r == nil {
		return SendErrorJSON(w, statusCode, err)
	}

	w.Header().Set(ContentTypeHeader, ApplicationJSON)
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(Error{
		Error:     err.Error(),
		Timestamp: time.Now(),
		RequestID: pointer.GetStringP(middleware.GetReqID(r.Context())),
	}); err != nil {
		return fmt.Errorf("json encode failed: %w", err)
	}

	return nil
}
