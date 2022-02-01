package errs

import (
	"net/http"

	//"golang.org/x/text/message"
)

type AppError struct {
	Code int
	Message string
}

func NewNotFoundError(message string) *AppError {
	return &AppError{
		Message: message,
		Code: http.StatusNotFound,
	}
}

func NewUnexpectedError(message string) *AppError {
	return &AppError{
		Message: message,
		Code: http.StatusInternalServerError,
	}
}