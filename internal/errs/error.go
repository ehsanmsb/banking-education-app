package errs

import "net/http"

type AppError struct {
	Message string
	Code    int
}

func (err AppError) NewNotFoundError(message string) *AppError {
	return &AppError{Message: message, Code: http.StatusNotFound}
}

func (err AppError) NewInternalServerError(message string) *AppError {
	return &AppError{Message: message, Code: http.StatusInternalServerError}
}

func (err AppError) NewBadRequestError(message string) *AppError {
	return &AppError{Message: message, Code: http.StatusBadRequest}
}
