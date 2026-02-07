package handler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewOkResponse(t *testing.T) {
	resp := newOkResponse()
	assert.True(t, resp.IsOK)
	assert.Empty(t, resp.Msg)
}

func TestNewBadRequestResponse(t *testing.T) {
	msg := "test error message"
	resp := newBadRequestResponse(msg)
	assert.False(t, resp.IsOK)
	assert.Equal(t, msg, resp.Msg)
}

func TestNewInternalErrorResponse(t *testing.T) {
	resp := newInternalErrorResponse()
	assert.False(t, resp.IsOK)
	assert.Equal(t, "internal server error", resp.Msg)
}
