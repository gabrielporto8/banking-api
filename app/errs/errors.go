package errs

import "net/http"

type AppError struct {
	Code int `json:",omitempty"`
	Err error `json:"err"`
}

func (e AppError) Error() string {
    return e.Err.Error()
}

func NewNotFoundError(err error) *AppError {
	return &AppError{
		Err: err,
		Code: http.StatusNotFound,
	}
}

func NewInternalError(err error) *AppError {
	return &AppError{
		Err: err,
		Code: http.StatusInternalServerError,
	}
}

func NewUnauthorizedError(err error) *AppError {
	return &AppError{
		Err: err,
		Code: http.StatusUnauthorized,
	}
}

func NewValidationError(err error) *AppError {
	return &AppError{
		Err: err,
		Code: http.StatusBadRequest,
	}
}

func NewConflictError(err error) *AppError {
	return &AppError{
		Err: err,
		Code: http.StatusConflict,
	}
}