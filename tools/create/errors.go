package create

import (
	"errors"
)

var (
	// ErrUnableToParseModFile defines an error for failing to parse the go.mod file.
	ErrUnableToParseModFile = errors.New("unable to parse go.mod file")

	// ErrLineNotFound defines an error for being unable to find a line number in a file.
	ErrLineNotFound = errors.New("line number not found")

	// ErrNameEmpty defines an error for an empty name.
	ErrNameEmpty = errors.New("name should not be empty")

	// ErrNameNoUnderscore defines an error for a name with underscores (_).
	ErrNameNoUnderscore = errors.New("no underscores (_) in name")

	// ErrNameNoUnderscore defines an error for a name with hyphens (-).
	ErrNameNoHyphens = errors.New("no hyphens (-) in name")

	// ErrNameNoUnderscore defines an error for a name that is not lowercased.
	ErrNameNotLowercase = errors.New("name should be all lowercase")
)
