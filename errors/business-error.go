package error

import (
	"runtime/debug"
)

type BusinessError struct {
	Message    string
	StackTrace string
	Err        error
	ErrorCode  uint
}

func (e BusinessError) Error() string {
	return e.Err.Error()
}

func NewBusinessError(message string, err error, errorCode uint) *BusinessError {
	return &BusinessError{
		Message:    message,
		StackTrace: string(debug.Stack()),
		Err:        err,
		ErrorCode:  errorCode,
	}
}

type Kind uint8

const (
	Other         Kind = iota // Unclassified error. This value is not printed in the error message.
	Invalid                   // Invalid operation for this type of item.
	Permission                // Permission denied.
	IO                        // External I/O error such as network failure.
	Exist                     // Item already exists.
	NotExist                  // Item does not exist.
	IsDir                     // Item is a directory.
	NotDir                    // Item is not a directory.
	NotEmpty                  // Directory not empty.
	Private                   // Information withheld.
	Internal                  // Internal error or inconsistency.
	CannotDecrypt             // No wrapped key for user with read access.
	Transient                 // A transient error.
	BrokenLink                // Link target does not exist.
)
