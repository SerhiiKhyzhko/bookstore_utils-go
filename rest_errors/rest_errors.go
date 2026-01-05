package rest_errors

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RestErr interface {
	Message() string
	Status() int
	Error() string
	Causes() []any
}

type restErr struct {
	ErrorMessage string `json:"Message"`
	ErrorStatus  int `json:"Status"`
	ErrorText   string `json:"Error"`
	ErrorCauses []any `json:"Causes"`
}

func (e restErr) Message() string{
	return e.ErrorMessage
}

func (e restErr) Status() int{
	return e.ErrorStatus
}

func (e restErr) Causes() []any {
	return e.ErrorCauses
}

func (e restErr) Error() string {
	return fmt.Sprintf("message: %s - status: %d - error: %s - causes: [ %v ]", 
	e.ErrorMessage, e.ErrorStatus, e.ErrorText, e.ErrorCauses)
}

func NewRestError(message string, status int, err string, causes []any) RestErr {
	return restErr{
		ErrorMessage: message,
		ErrorStatus: status,
		ErrorText: err,
		ErrorCauses: causes,
	}
}

func NewBadRequestError(message string) RestErr {
	return restErr{
		ErrorMessage: message,
		ErrorStatus: http.StatusBadRequest,
		ErrorText: "bad request",
	}
}

func NewNotFoundError(message string) RestErr {
	return restErr{
		ErrorMessage: message,
		ErrorStatus: http.StatusNotFound,
		ErrorText: "not found",
	}
}

func NewInternalServerError(message string, err error) RestErr {
	result := restErr{
		ErrorMessage: message,
		ErrorStatus: http.StatusInternalServerError,
		ErrorText: "internal server error",
	}
	if err != nil {
		result.ErrorCauses = append(result.ErrorCauses, err.Error())
	}
	return result
}

func NewRestErrorFromBytes(bytes []byte) (RestErr, error) {
	var apiErr restErr
	if err := json.Unmarshal(bytes, &apiErr); err != nil {
		return nil, err
	}
	return apiErr, nil
}