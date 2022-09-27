package openapi

import (
	"errors"
)

var (
	// ErrValidation denotes the request did not meet the OpenAPI specification.
	ErrValidation = errors.New("oapi validation failed")
)
