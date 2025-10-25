package rest_errors

import (
	"errors"
	"net/http"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerError("test error message", errors.New("database error"))
	assert.NotEqual(t, err, nil)
	assert.Equal(t, err.Message, "test error message")
	assert.Equal(t, err.Status, http.StatusInternalServerError)
	assert.Equal(t, err.Error, "internal server error")

	assert.NotEqual(t, err.Causes, nil)
	assert.Equal(t, len(err.Causes), 1)
	assert.Equal(t, err.Causes[0], "database error")
}

func TestNewNotFoundError(t *testing.T) {
	err := NewNotFoundError("test email not found")
	assert.NotEqual(t, err, nil)
	assert.Equal(t, err.Message, "test email not found")
	assert.Equal(t, err.Status, http.StatusNotFound)
	assert.Equal(t, err.Error, "not found")
}

func TestNewBadRequestError(t *testing.T) {
	err := NewBadRequestError("invalid test email")
	assert.NotEqual(t, err, nil)
	assert.Equal(t, err.Message, "invalid test email")
	assert.Equal(t, err.Status, http.StatusBadRequest)
	assert.Equal(t, err.Error, "bad request")
}

func TestNewError(t *testing.T) {
	err := NewError("test custom error")
	assert.NotEqual(t, err, nil)
	assert.Equal(t, err.Error(), "test custom error")
}