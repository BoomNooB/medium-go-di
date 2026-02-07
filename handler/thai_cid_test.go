package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestThaiCIDHandler_ValidateThaiCID_Success(t *testing.T) {
	// Setup
	e := echo.New()
	req := ThaiCIDRequest{
		CitizenID: "1234567890123",
		FullName:  "John Doe",
	}
	reqBody, _ := json.Marshal(req)
	httpReq := httptest.NewRequest(http.MethodPost, "/thai-cid", bytes.NewReader(reqBody))
	httpReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(httpReq, rec)

	// Test
	v := validator.New(validator.WithRequiredStructEnabled())
	h := NewThaiCIDHandler(v)
	err := h.ValidateThaiCID(c)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var resp Response
	json.Unmarshal(rec.Body.Bytes(), &resp)
	assert.True(t, resp.IsOK)
	assert.Empty(t, resp.Msg)
}

func TestThaiCIDHandler_ValidateThaiCID_InvalidJSON(t *testing.T) {
	// Setup
	e := echo.New()
	invalidJSON := []byte(`{"citizenId": "not-closed`)
	httpReq := httptest.NewRequest(http.MethodPost, "/thai-cid", bytes.NewReader(invalidJSON))
	httpReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(httpReq, rec)

	// Test
	v := validator.New(validator.WithRequiredStructEnabled())
	h := NewThaiCIDHandler(v)
	err := h.ValidateThaiCID(c)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var resp Response
	json.Unmarshal(rec.Body.Bytes(), &resp)
	assert.False(t, resp.IsOK)
	assert.Equal(t, badRequestJSONSyntax, resp.Msg)
}

func TestThaiCIDHandler_ValidateThaiCID_MissingCitizenID(t *testing.T) {
	// Setup
	e := echo.New()
	req := ThaiCIDRequest{
		FullName: "John Doe",
	}
	reqBody, _ := json.Marshal(req)
	httpReq := httptest.NewRequest(http.MethodPost, "/thai-cid", bytes.NewReader(reqBody))
	httpReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(httpReq, rec)

	// Test
	v := validator.New(validator.WithRequiredStructEnabled())
	h := NewThaiCIDHandler(v)
	err := h.ValidateThaiCID(c)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var resp Response
	json.Unmarshal(rec.Body.Bytes(), &resp)
	assert.False(t, resp.IsOK)
	assert.Equal(t, badRequestNotValid, resp.Msg)
}

func TestThaiCIDHandler_ValidateThaiCID_CitizenIDTooShort(t *testing.T) {
	// Setup
	e := echo.New()
	req := ThaiCIDRequest{
		CitizenID: "12345",
		FullName:  "John Doe",
	}
	reqBody, _ := json.Marshal(req)
	httpReq := httptest.NewRequest(http.MethodPost, "/thai-cid", bytes.NewReader(reqBody))
	httpReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(httpReq, rec)

	// Test
	v := validator.New(validator.WithRequiredStructEnabled())
	h := NewThaiCIDHandler(v)
	err := h.ValidateThaiCID(c)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var resp Response
	json.Unmarshal(rec.Body.Bytes(), &resp)
	assert.False(t, resp.IsOK)
	assert.Equal(t, badRequestNotValid, resp.Msg)
}

func TestThaiCIDHandler_ValidateThaiCID_CitizenIDTooLong(t *testing.T) {
	// Setup
	e := echo.New()
	req := ThaiCIDRequest{
		CitizenID: "12345678901234",
		FullName:  "John Doe",
	}
	reqBody, _ := json.Marshal(req)
	httpReq := httptest.NewRequest(http.MethodPost, "/thai-cid", bytes.NewReader(reqBody))
	httpReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(httpReq, rec)

	// Test
	v := validator.New(validator.WithRequiredStructEnabled())
	h := NewThaiCIDHandler(v)
	err := h.ValidateThaiCID(c)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var resp Response
	json.Unmarshal(rec.Body.Bytes(), &resp)
	assert.False(t, resp.IsOK)
	assert.Equal(t, badRequestNotValid, resp.Msg)
}

func TestThaiCIDHandler_ValidateThaiCID_CitizenIDNotNumeric(t *testing.T) {
	// Setup
	e := echo.New()
	req := ThaiCIDRequest{
		CitizenID: "123456789012A",
		FullName:  "John Doe",
	}
	reqBody, _ := json.Marshal(req)
	httpReq := httptest.NewRequest(http.MethodPost, "/thai-cid", bytes.NewReader(reqBody))
	httpReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(httpReq, rec)

	// Test
	v := validator.New(validator.WithRequiredStructEnabled())
	h := NewThaiCIDHandler(v)
	err := h.ValidateThaiCID(c)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var resp Response
	json.Unmarshal(rec.Body.Bytes(), &resp)
	assert.False(t, resp.IsOK)
	assert.Equal(t, badRequestNotValid, resp.Msg)
}

func TestThaiCIDHandler_ValidateThaiCID_MissingFullName(t *testing.T) {
	// Setup
	e := echo.New()
	req := ThaiCIDRequest{
		CitizenID: "1234567890123",
	}
	reqBody, _ := json.Marshal(req)
	httpReq := httptest.NewRequest(http.MethodPost, "/thai-cid", bytes.NewReader(reqBody))
	httpReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(httpReq, rec)

	// Test
	v := validator.New(validator.WithRequiredStructEnabled())
	h := NewThaiCIDHandler(v)
	err := h.ValidateThaiCID(c)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var resp Response
	json.Unmarshal(rec.Body.Bytes(), &resp)
	assert.False(t, resp.IsOK)
	assert.Equal(t, badRequestNotValid, resp.Msg)
}

func TestThaiCIDHandler_ValidateThaiCID_FullNameTooShort(t *testing.T) {
	// Setup
	e := echo.New()
	req := ThaiCIDRequest{
		CitizenID: "1234567890123",
		FullName:  "AB",
	}
	reqBody, _ := json.Marshal(req)
	httpReq := httptest.NewRequest(http.MethodPost, "/thai-cid", bytes.NewReader(reqBody))
	httpReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(httpReq, rec)

	// Test
	v := validator.New(validator.WithRequiredStructEnabled())
	h := NewThaiCIDHandler(v)
	err := h.ValidateThaiCID(c)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var resp Response
	json.Unmarshal(rec.Body.Bytes(), &resp)
	assert.False(t, resp.IsOK)
	assert.Equal(t, badRequestNotValid, resp.Msg)
}

func TestNewThaiCIDHandler(t *testing.T) {
	v := validator.New(validator.WithRequiredStructEnabled())
	h := NewThaiCIDHandler(v)
	assert.NotNil(t, h)
	assert.NotNil(t, h.v)
}
