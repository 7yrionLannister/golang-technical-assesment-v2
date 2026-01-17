package util

import (
	"fmt"
	"strings"
)

// Logs the error and returns a new error with the message and the original error.
// The message is converted to lowercase, it is best practice for error strings.
func HandleError(err error, message string) error {
	e := fmt.Errorf("%s: %w", strings.ToLower(message), err)
	// TODO log error if logging middleware doesn't
	return e
}
