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
	message string
	status  int 
	error   string 
	causes []any 
}

func (e restErr) Message() string{
	return e.message
}

func (e restErr) Status() int{
	return e.status
}

func (e restErr) Causes() []any {
	return e.causes
}

func (e restErr) Error() string {
	return fmt.Sprintf("message: %s - status: %d - error: %s - causes: [ %v ]", 
	e.message, e.status, e.error, e.causes)
}

func NewRestError(message string, status int, err string, causes []any) RestErr {
	return restErr{
		message: message,
		status: status,
		error: err,
		causes: causes,
	}
}

func NewBadRequestError(message string) RestErr {
	return restErr{
		message: message,
		status: http.StatusBadRequest,
		error: "bad request",
	}
}

func NewNotFoundError(message string) RestErr {
	return restErr{
		message: message,
		status: http.StatusNotFound,
		error: "not found",
	}
}

func NewInternalServerError(message string, err error) RestErr {
	result := restErr{
		message: message,
		status: http.StatusInternalServerError,
		error: "internal server error",
	}
	if err != nil {
		result.causes = append(result.causes, err.Error())
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