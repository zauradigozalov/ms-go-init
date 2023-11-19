package model

import (
	"fmt"
	"net/http"
)

const (
	Exception       = "error.exception"
	ContentType     = "Content-Type"
	ApplicationJSON = "application/json"
	ContextHeader   = "contextHeader"
)

type ErrorResponse struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

type RESTCompatible interface {
	GetException() *RestException
}

func (e *ErrorResponse) ToError(statusCode int) error {
	return &RestException{
		Message:    e.Message,
		Code:       e.Code,
		StatusCode: statusCode,
	}
}

type RestException struct {
	Message    string `json:"message"`
	Code       string `json:"code"`
	StatusCode int    `json:"-"`
}

func (e *RestException) Error() string {
	return e.Message
}

func (e *RestException) GetException() *RestException {
	return e
}

type BadRequestError struct {
	RestException
}

func NewBadRequestError(message string) error {
	return &BadRequestError{
		RestException{
			Message:    message,
			Code:       fmt.Sprintf("%s.invalid.request", Exception),
			StatusCode: http.StatusBadRequest,
		},
	}
}

type InternalServerError struct {
	RestException
}

func NewInternalServerError(message string) error {
	return &InternalServerError{
		RestException{
			Message:    message,
			Code:       fmt.Sprintf("%s.internal.server.error", Exception),
			StatusCode: http.StatusInternalServerError,
		},
	}
}
