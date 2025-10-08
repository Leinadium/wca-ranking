package handler

var (
	ErrCodeDefault        = "000"
	ErrCodeWCADefault     = "001"
	ErrCodeWCAUserInvalid = "002"
	ErrCodeWCAUserWait    = "003"

	ErrMessageDefault        = "Unknown Error"
	ErrMessageWCADefault     = "Could not reach WCA API. Try again later"
	ErrMessageWCAUserInvalid = "Attempting to modify another user"
	ErrMessageWCAUserWait    = "Wait %d hours to update again"
)

type BaseError struct {
	Code    *string
	Message *string
}

var (
	ErrDefault BaseError = BaseError{
		Code:    &ErrCodeDefault,
		Message: &ErrMessageDefault,
	}
	ErrWCADefault BaseError = BaseError{
		Code:    &ErrCodeWCADefault,
		Message: &ErrMessageWCADefault,
	}
	ErrUserInvalid BaseError = BaseError{
		Code:    &ErrCodeWCAUserInvalid,
		Message: &ErrMessageWCAUserInvalid,
	}
)
