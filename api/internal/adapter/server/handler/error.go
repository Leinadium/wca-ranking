package handler

var (
	ErrDefaultCode    = "000"
	ErrDefaultMessage = "Unknown Error"
)

type BaseError struct {
	Code    *string
	Message *string
}

var ErrDefault BaseError = BaseError{
	Code:    &ErrDefaultCode,
	Message: &ErrDefaultMessage,
}
