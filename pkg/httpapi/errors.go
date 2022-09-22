package httpapi

import (
	"errors"
)

var (
	// ErrBadRequest denotes the client attempted a bad request.
	ErrBadRequest = errors.New("bad request")
)
