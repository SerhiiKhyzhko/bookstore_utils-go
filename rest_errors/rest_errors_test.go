package rest_errors

import (
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerError("test error message", errors.New("database error"))
	assert.NotEqual(t, err, nil)
	assert.Equal(t, err.Message(), "test error message")
	assert.Equal(t, err.Status(), http.StatusInternalServerError)
	assert.Equal(t, err.Error(), "message: test error message - status: 500 - error: internal server error - causes: [ [database error] ]")

	causes := err.Causes()
	assert.NotEqual(t, causes, nil)
	assert.Equal(t, len(causes), 1)
	assert.Equal(t, causes[0], "database error")
}

func TestNewNotFoundError(t *testing.T) {
	err := NewNotFoundError("test email not found")
	assert.NotEqual(t, err, nil)
	assert.Equal(t, err.Message(), "test email not found")
	assert.Equal(t, err.Status(), http.StatusNotFound)
	assert.Equal(t, err.Error(), "message: test email not found - status: 404 - error: not found - causes: [ [] ]")
}

func TestNewBadRequestError(t *testing.T) {
	err := NewBadRequestError("invalid test email")
	assert.NotEqual(t, err, nil)
	assert.Equal(t, err.Message(), "invalid test email")
	assert.Equal(t, err.Status(), http.StatusBadRequest)
	assert.Equal(t, err.Error(), "message: invalid test email - status: 400 - error: bad request - causes: [ [] ]")
}

func TestNewError(t *testing.T) {
	err := NewRestError("test custom error", 400, "bad request", nil)
	assert.NotEqual(t, err, nil)
	assert.Equal(t, err.Message(), "test custom error")
	assert.Equal(t, err.Status(), 400)
	assert.Equal(t, err.Error(), "message: test custom error - status: 400 - error: bad request - causes: [ [] ]")
	assert.Equal(t, err.Causes(), nil)
}

func TestNewRestErrorFromBytes(t *testing.T) {
	tStruct := NewRestError("test custom error", 400, "test error", nil)
	tJSON, _ := json.Marshal(tStruct) 

	testRes, err := NewRestErrorFromBytes(tJSON)
	assert.Equal(t, err, nil)
	assert.NotEqual(t, testRes, nil)
	assert.Equal(t, testRes.Message(), "test custom error")
	assert.Equal(t, testRes.Status(), 400)
	assert.Equal(t, testRes.Error(), "message: test custom error - status: 400 - error: test error - causes: [ [] ]")
	assert.Equal(t, testRes.Causes(), nil)
}