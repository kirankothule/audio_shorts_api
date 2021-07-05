package rest_utils

import (
	"errors"
	"net/http"
)

type RestErr struct {
	Meassage string `json:"message"`
	Status   int    `json:"status"`
	Error    string `json:"error"`
}

func NewError(msg string) error {
	return errors.New(msg)
}

func NewBadRequestError(message string) *RestErr {
	return &RestErr{
		Meassage: message,
		Status:   http.StatusBadRequest,
		Error:    "bad_request",
	}
}

func NewNotFoundError(message string) *RestErr {
	return &RestErr{
		Meassage: message,
		Status:   http.StatusNotFound,
		Error:    "not_found",
	}
}

func NewInternalServerError(message string) *RestErr {
	return &RestErr{
		Meassage: message,
		Status:   http.StatusInternalServerError,
		Error:    "internal_server_error",
	}
}
