package errors

import "net/http"

type ApiError interface {
	Status() int
	Message() string
	Error() string
}

type apiError struct {
	Astatus  int    `json:"status"`
	Amessage string `json:"message"`
//omitempty : dont show when "error" is empty
	Aerror   string `json:"error,omitempty"`
}

func (e *apiError) Status() int {
	return e.Astatus
}

func (e *apiError) Message() string {
	return e.Amessage
}

func (e *apiError) Error() string {
	return e.Aerror
}

func NewNotFoundApiError(message string) ApiError {
	return &apiError{
		Astatus: http.StatusNotFound,
		Amessage: message,
	}
}

func NewInternalServerError(message string) ApiError {
	return &apiError{
		Astatus: http.StatusNotFound,
		Amessage: message,
	}
}

func NewBadRequesError(message string) ApiError {
	return &apiError{
		Astatus: http.StatusBadRequest,
		Amessage: message,
	}
}

func NewApiError(statusCode int,message string) ApiError{
	return &apiError{
		Astatus: statusCode,
		Amessage: message,
	}
}