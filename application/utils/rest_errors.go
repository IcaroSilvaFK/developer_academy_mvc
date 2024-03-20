package utils

import (
	"net/http"
)

type RestErr struct {
	Message string   `json:"message"`
	Err     string   `json:"error"`
	Code    int      `json:"code"`
	Causes  []Causes `json:"causes,omitempty"`
}

type Causes struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (err *RestErr) Error() string {
	return err.Message
}

func NewRestErr(message, err string, code int, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     err,
		Code:    code,
		Causes:  causes,
	}
}

func NewBadRequestException(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     BAD_REQUEST,
		Code:    http.StatusBadRequest,
	}
}

func NewBadRequestValidationException(message string, causes []Causes) *RestErr {
	return &RestErr{
		Message: message,
		Err:     BAD_REQUEST,
		Code:    http.StatusBadRequest,
		Causes:  causes,
	}
}

func NewInternalServerError() *RestErr {
	return &RestErr{
		Code: http.StatusInternalServerError,
		Err:  INTERNAL_SERVER_ERROR,
	}
}

func NewNotFoundException(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     NOT_FOUND_ERROR,
		Code:    http.StatusNotFound,
	}
}

func NewForbiddenException(message string) *RestErr {
	return &RestErr{
		Message: message,
		Err:     FORBIDDEN_ERROR,
		Code:    http.StatusForbidden,
	}
}
