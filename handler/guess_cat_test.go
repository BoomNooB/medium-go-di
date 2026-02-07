package handler

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/BoomNooB/medium-go-di/validatorwrapper"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGuessCatNameHandler_GuessTheCatName_Success(t *testing.T) {
	// Setup
	e := echo.New()
	req := GuessCatNameRequest{
		GuessName: "Fluffy",
		UserID:    "550e8400-e29b-41d4-a716-446655440000",
		Attempts:  2,
	}
	reqBody, _ := json.Marshal(req)
	httpReq := httptest.NewRequest(http.MethodPost, "/guess-cat", bytes.NewReader(reqBody))
	httpReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(httpReq, rec)

	// Test
	v := validator.New(validator.WithRequiredStructEnabled())
	vw := validatorwrapper.NewValidatorWrapper(v)
	h := NewGuessCatNameHandler(vw)
	err := h.GuessTheCatName(c)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var resp Response
	json.Unmarshal(rec.Body.Bytes(), &resp)
	assert.True(t, resp.IsOK)
	assert.Empty(t, resp.Msg)
}

func TestGuessCatNameHandler_GuessTheCatName_InvalidJSON(t *testing.T) {
	// Setup
	e := echo.New()
	invalidJSON := []byte(`{"guessName": "not-closed`)
	httpReq := httptest.NewRequest(http.MethodPost, "/guess-cat", bytes.NewReader(invalidJSON))
	httpReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(httpReq, rec)

	// Test
	v := validator.New(validator.WithRequiredStructEnabled())
	vw := validatorwrapper.NewValidatorWrapper(v)
	h := NewGuessCatNameHandler(vw)
	err := h.GuessTheCatName(c)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var resp Response
	json.Unmarshal(rec.Body.Bytes(), &resp)
	assert.False(t, resp.IsOK)
	assert.Equal(t, badRequestJSONSyntax, resp.Msg)
}

func TestGuessCatNameHandler_GuessTheCatName_MissingGuessName(t *testing.T) {
	// Setup
	e := echo.New()
	req := GuessCatNameRequest{
		UserID:   "550e8400-e29b-41d4-a716-446655440000",
		Attempts: 2,
	}
	reqBody, _ := json.Marshal(req)
	httpReq := httptest.NewRequest(http.MethodPost, "/guess-cat", bytes.NewReader(reqBody))
	httpReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(httpReq, rec)

	// Test
	v := validator.New(validator.WithRequiredStructEnabled())
	vw := validatorwrapper.NewValidatorWrapper(v)
	h := NewGuessCatNameHandler(vw)
	err := h.GuessTheCatName(c)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var resp Response
	json.Unmarshal(rec.Body.Bytes(), &resp)
	assert.False(t, resp.IsOK)
	assert.Equal(t, badRequestNotValid, resp.Msg)
}

func TestGuessCatNameHandler_GuessTheCatName_NameTooLong(t *testing.T) {
	// Setup
	e := echo.New()
	req := GuessCatNameRequest{
		GuessName: "ThisIsAReallyLongCatNameThatExceedsTheMaximumLength",
		UserID:    "550e8400-e29b-41d4-a716-446655440000",
		Attempts:  2,
	}
	reqBody, _ := json.Marshal(req)
	httpReq := httptest.NewRequest(http.MethodPost, "/guess-cat", bytes.NewReader(reqBody))
	httpReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(httpReq, rec)

	// Test
	v := validator.New(validator.WithRequiredStructEnabled())
	vw := validatorwrapper.NewValidatorWrapper(v)
	h := NewGuessCatNameHandler(vw)
	err := h.GuessTheCatName(c)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var resp Response
	json.Unmarshal(rec.Body.Bytes(), &resp)
	assert.False(t, resp.IsOK)
	assert.Equal(t, badRequestNotValid, resp.Msg)
}

func TestGuessCatNameHandler_GuessTheCatName_InvalidUserID(t *testing.T) {
	// Setup
	e := echo.New()
	req := GuessCatNameRequest{
		GuessName: "Fluffy",
		UserID:    "not-a-uuid",
		Attempts:  2,
	}
	reqBody, _ := json.Marshal(req)
	httpReq := httptest.NewRequest(http.MethodPost, "/guess-cat", bytes.NewReader(reqBody))
	httpReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(httpReq, rec)

	// Test
	v := validator.New(validator.WithRequiredStructEnabled())
	vw := validatorwrapper.NewValidatorWrapper(v)
	h := NewGuessCatNameHandler(vw)
	err := h.GuessTheCatName(c)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var resp Response
	json.Unmarshal(rec.Body.Bytes(), &resp)
	assert.False(t, resp.IsOK)
	assert.Equal(t, badRequestNotValid, resp.Msg)
}

func TestGuessCatNameHandler_GuessTheCatName_AttemptsZero(t *testing.T) {
	// Setup
	e := echo.New()
	req := GuessCatNameRequest{
		GuessName: "Fluffy",
		UserID:    "550e8400-e29b-41d4-a716-446655440000",
		Attempts:  0,
	}
	reqBody, _ := json.Marshal(req)
	httpReq := httptest.NewRequest(http.MethodPost, "/guess-cat", bytes.NewReader(reqBody))
	httpReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(httpReq, rec)

	// Test
	v := validator.New(validator.WithRequiredStructEnabled())
	vw := validatorwrapper.NewValidatorWrapper(v)
	h := NewGuessCatNameHandler(vw)
	err := h.GuessTheCatName(c)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var resp Response
	json.Unmarshal(rec.Body.Bytes(), &resp)
	assert.False(t, resp.IsOK)
	assert.Equal(t, badRequestNotValid, resp.Msg)
}

func TestGuessCatNameHandler_GuessTheCatName_AttemptsExceedsMax(t *testing.T) {
	// Setup
	e := echo.New()
	req := GuessCatNameRequest{
		GuessName: "Fluffy",
		UserID:    "550e8400-e29b-41d4-a716-446655440000",
		Attempts:  5,
	}
	reqBody, _ := json.Marshal(req)
	httpReq := httptest.NewRequest(http.MethodPost, "/guess-cat", bytes.NewReader(reqBody))
	httpReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(httpReq, rec)

	// Test
	v := validator.New(validator.WithRequiredStructEnabled())
	vw := validatorwrapper.NewValidatorWrapper(v)
	h := NewGuessCatNameHandler(vw)
	err := h.GuessTheCatName(c)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var resp Response
	json.Unmarshal(rec.Body.Bytes(), &resp)
	assert.False(t, resp.IsOK)
	assert.Equal(t, badRequestNotValid, resp.Msg)
}

func TestGuessCatNameHandler_GuessTheCatName_InternalError(t *testing.T) {
	// Setup
	e := echo.New()
	req := GuessCatNameRequest{
		GuessName: "Fluffy",
		UserID:    "550e8400-e29b-41d4-a716-446655440000",
		Attempts:  2,
	}
	reqBody, _ := json.Marshal(req)
	httpReq := httptest.NewRequest(http.MethodPost, "/guess-cat", bytes.NewReader(reqBody))
	httpReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(httpReq, rec)

	// Test - Mock validator that returns a non-validation error
	mockV := &mockValidator{err: errors.New("unexpected internal error")}
	h := NewGuessCatNameHandler(mockV)
	err := h.GuessTheCatName(c)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)

	var resp Response
	json.Unmarshal(rec.Body.Bytes(), &resp)
	assert.False(t, resp.IsOK)
	assert.Equal(t, "internal server error", resp.Msg)
}

func TestNewGuessCatNameHandler(t *testing.T) {
	v := validator.New(validator.WithRequiredStructEnabled())
	vw := validatorwrapper.NewValidatorWrapper(v)
	h := NewGuessCatNameHandler(vw)
	assert.NotNil(t, h)
	assert.NotNil(t, h.v)
}
