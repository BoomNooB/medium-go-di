package handler

import (
	"bytes"
	"context"
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

// mockValidator implements Valiator interface for testing internal errors
type mockValidator struct {
	err error
}

func (m *mockValidator) StructValidation(ctx context.Context, req any) error {
	return m.err
}

func TestFavoriteNumHandler_Favorite_Success(t *testing.T) {
	// Setup
	e := echo.New()
	req := FavoriteNumRequest{
		UserID: "550e8400-e29b-41d4-a716-446655440000",
		FavNum: 42,
	}
	reqBody, _ := json.Marshal(req)
	httpReq := httptest.NewRequest(http.MethodPost, "/favorite", bytes.NewReader(reqBody))
	httpReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(httpReq, rec)

	// Test
	v := validator.New(validator.WithRequiredStructEnabled())
	vw := validatorwrapper.NewValidatorWrapper(v)
	h := NewFavoriteNumHandler(vw)
	err := h.Favorite(c)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var resp Response
	json.Unmarshal(rec.Body.Bytes(), &resp)
	assert.True(t, resp.IsOK)
	assert.Empty(t, resp.Msg)
}

func TestFavoriteNumHandler_Favorite_InvalidJSON(t *testing.T) {
	// Setup
	e := echo.New()
	invalidJSON := []byte(`{"userId": "not-closed`)
	httpReq := httptest.NewRequest(http.MethodPost, "/favorite", bytes.NewReader(invalidJSON))
	httpReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(httpReq, rec)

	// Test
	v := validator.New(validator.WithRequiredStructEnabled())
	vw := validatorwrapper.NewValidatorWrapper(v)
	h := NewFavoriteNumHandler(vw)
	err := h.Favorite(c)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var resp Response
	json.Unmarshal(rec.Body.Bytes(), &resp)
	assert.False(t, resp.IsOK)
	assert.Equal(t, badRequestJSONSyntax, resp.Msg)
}

func TestFavoriteNumHandler_Favorite_MissingUserID(t *testing.T) {
	// Setup
	e := echo.New()
	req := FavoriteNumRequest{
		FavNum: 42,
	}
	reqBody, _ := json.Marshal(req)
	httpReq := httptest.NewRequest(http.MethodPost, "/favorite", bytes.NewReader(reqBody))
	httpReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(httpReq, rec)

	// Test
	v := validator.New(validator.WithRequiredStructEnabled())
	vw := validatorwrapper.NewValidatorWrapper(v)
	h := NewFavoriteNumHandler(vw)
	err := h.Favorite(c)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var resp Response
	json.Unmarshal(rec.Body.Bytes(), &resp)
	assert.False(t, resp.IsOK)
	assert.Equal(t, badRequestNotValid, resp.Msg)
}

func TestFavoriteNumHandler_Favorite_InvalidUUID(t *testing.T) {
	// Setup
	e := echo.New()
	req := FavoriteNumRequest{
		UserID: "not-a-uuid",
		FavNum: 42,
	}
	reqBody, _ := json.Marshal(req)
	httpReq := httptest.NewRequest(http.MethodPost, "/favorite", bytes.NewReader(reqBody))
	httpReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(httpReq, rec)

	// Test
	v := validator.New(validator.WithRequiredStructEnabled())
	vw := validatorwrapper.NewValidatorWrapper(v)
	h := NewFavoriteNumHandler(vw)
	err := h.Favorite(c)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var resp Response
	json.Unmarshal(rec.Body.Bytes(), &resp)
	assert.False(t, resp.IsOK)
	assert.Equal(t, badRequestNotValid, resp.Msg)
}

func TestFavoriteNumHandler_Favorite_MissingFavNum(t *testing.T) {
	// Setup
	e := echo.New()
	req := FavoriteNumRequest{
		UserID: "550e8400-e29b-41d4-a716-446655440000",
	}
	reqBody, _ := json.Marshal(req)
	httpReq := httptest.NewRequest(http.MethodPost, "/favorite", bytes.NewReader(reqBody))
	httpReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(httpReq, rec)

	// Test
	v := validator.New(validator.WithRequiredStructEnabled())
	vw := validatorwrapper.NewValidatorWrapper(v)
	h := NewFavoriteNumHandler(vw)
	err := h.Favorite(c)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var resp Response
	json.Unmarshal(rec.Body.Bytes(), &resp)
	assert.False(t, resp.IsOK)
	assert.Equal(t, badRequestNotValid, resp.Msg)
}

func TestFavoriteNumHandler_Favorite_ZeroFavNum(t *testing.T) {
	// Setup
	e := echo.New()
	req := FavoriteNumRequest{
		UserID: "550e8400-e29b-41d4-a716-446655440000",
		FavNum: 0,
	}
	reqBody, _ := json.Marshal(req)
	httpReq := httptest.NewRequest(http.MethodPost, "/favorite", bytes.NewReader(reqBody))
	httpReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(httpReq, rec)

	// Test
	v := validator.New(validator.WithRequiredStructEnabled())
	vw := validatorwrapper.NewValidatorWrapper(v)
	h := NewFavoriteNumHandler(vw)
	err := h.Favorite(c)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var resp Response
	json.Unmarshal(rec.Body.Bytes(), &resp)
	assert.False(t, resp.IsOK)
	assert.Equal(t, badRequestNotValid, resp.Msg)
}

func TestFavoriteNumHandler_Favorite_NegativeFavNum(t *testing.T) {
	// Setup
	e := echo.New()
	req := FavoriteNumRequest{
		UserID: "550e8400-e29b-41d4-a716-446655440000",
		FavNum: -5,
	}
	reqBody, _ := json.Marshal(req)
	httpReq := httptest.NewRequest(http.MethodPost, "/favorite", bytes.NewReader(reqBody))
	httpReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(httpReq, rec)

	// Test
	v := validator.New(validator.WithRequiredStructEnabled())
	vw := validatorwrapper.NewValidatorWrapper(v)
	h := NewFavoriteNumHandler(vw)
	err := h.Favorite(c)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var resp Response
	json.Unmarshal(rec.Body.Bytes(), &resp)
	assert.False(t, resp.IsOK)
	assert.Equal(t, badRequestNotValid, resp.Msg)
}

func TestFavoriteNumHandler_Favorite_InternalError(t *testing.T) {
	// Setup
	e := echo.New()
	req := FavoriteNumRequest{
		UserID: "550e8400-e29b-41d4-a716-446655440000",
		FavNum: 42,
	}
	reqBody, _ := json.Marshal(req)
	httpReq := httptest.NewRequest(http.MethodPost, "/favorite", bytes.NewReader(reqBody))
	httpReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(httpReq, rec)

	// Test - Mock validator that returns a non-validation error
	mockV := &mockValidator{err: errors.New("unexpected internal error")}
	h := NewFavoriteNumHandler(mockV)
	err := h.Favorite(c)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)

	var resp Response
	json.Unmarshal(rec.Body.Bytes(), &resp)
	assert.False(t, resp.IsOK)
	assert.Equal(t, "internal server error", resp.Msg)
}

func TestNewFavoriteNumHandler(t *testing.T) {
	v := validator.New(validator.WithRequiredStructEnabled())
	vw := validatorwrapper.NewValidatorWrapper(v)
	h := NewFavoriteNumHandler(vw)
	assert.NotNil(t, h)
	assert.NotNil(t, h.v)
}
