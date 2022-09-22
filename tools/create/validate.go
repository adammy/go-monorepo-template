package create

import (
	"strings"
)

// ValidateName validates the name supplied by the dev for a service/package.
func ValidateName(name string) error {
	switch {
	case name == "":
		return ErrNameEmpty
	case strings.Contains(name, "_"):
		return ErrNameNoUnderscore
	case strings.Contains(name, "-"):
		return ErrNameNoHyphens
	case strings.ToLower(name) != name:
		return ErrNameNotLowercase
	default:
		return nil
	}
}
