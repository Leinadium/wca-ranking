package handler

type serverError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Err     error  `json:"-"`
}

var (
	ErrDefault = serverError{Code: "000", Message: "Unknown Error", Err: nil}
)
