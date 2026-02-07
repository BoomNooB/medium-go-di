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

func TestPetNameHandler_ValidatePetName_Success(t *testing.T) {
	// Setup
	e := echo.New()
	req := PetNameRequest{
		PetName: "Buddy",
		OwnerID: "550e8400-e29b-41d4-a716-446655440000",
	}
	reqBody, _ := json.Marshal(req)
	httpReq := httptest.NewRequest(http.MethodPost, "/pet-name", bytes.NewReader(reqBody))
	httpReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(httpReq, rec)

	// Test
	v := validator.New(validator.WithRequiredStructEnabled())
	vw := validatorwrapper.NewValidatorWrapper(v)
	h := NewPetNameHandler(vw)
	err := h.ValidatePetName(c)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

	var resp Response
	json.Unmarshal(rec.Body.Bytes(), &resp)
	assert.True(t, resp.IsOK)
	assert.Empty(t, resp.Msg)
}

func TestPetNameHandler_ValidatePetName_InvalidJSON(t *testing.T) {
	// Setup
	e := echo.New()
	invalidJSON := []byte(`{"petName": "not-closed`)
	httpReq := httptest.NewRequest(http.MethodPost, "/pet-name", bytes.NewReader(invalidJSON))
	httpReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(httpReq, rec)

	// Test
	v := validator.New(validator.WithRequiredStructEnabled())
	vw := validatorwrapper.NewValidatorWrapper(v)
	h := NewPetNameHandler(vw)
	err := h.ValidatePetName(c)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var resp Response
	json.Unmarshal(rec.Body.Bytes(), &resp)
	assert.False(t, resp.IsOK)
	assert.Equal(t, badRequestJSONSyntax, resp.Msg)
}

func TestPetNameHandler_ValidatePetName_MissingPetName(t *testing.T) {
	// Setup
	e := echo.New()
	req := PetNameRequest{
		OwnerID: "550e8400-e29b-41d4-a716-446655440000",
	}
	reqBody, _ := json.Marshal(req)
	httpReq := httptest.NewRequest(http.MethodPost, "/pet-name", bytes.NewReader(reqBody))
	httpReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(httpReq, rec)

	// Test
	v := validator.New(validator.WithRequiredStructEnabled())
	vw := validatorwrapper.NewValidatorWrapper(v)
	h := NewPetNameHandler(vw)
	err := h.ValidatePetName(c)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var resp Response
	json.Unmarshal(rec.Body.Bytes(), &resp)
	assert.False(t, resp.IsOK)
	assert.Equal(t, badRequestNotValid, resp.Msg)
}

func TestPetNameHandler_ValidatePetName_NameTooShort(t *testing.T) {
	// Setup
	e := echo.New()
	req := PetNameRequest{
		PetName: "A",
		OwnerID: "550e8400-e29b-41d4-a716-446655440000",
	}
	reqBody, _ := json.Marshal(req)
	httpReq := httptest.NewRequest(http.MethodPost, "/pet-name", bytes.NewReader(reqBody))
	httpReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(httpReq, rec)

	// Test
	v := validator.New(validator.WithRequiredStructEnabled())
	vw := validatorwrapper.NewValidatorWrapper(v)
	h := NewPetNameHandler(vw)
	err := h.ValidatePetName(c)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var resp Response
	json.Unmarshal(rec.Body.Bytes(), &resp)
	assert.False(t, resp.IsOK)
	assert.Equal(t, badRequestNotValid, resp.Msg)
}

func TestPetNameHandler_ValidatePetName_NameTooLong(t *testing.T) {
	// Setup
	e := echo.New()
	req := PetNameRequest{
		PetName: "ThisIsAnExtremelyLongPetNameThatExceedsFiftyCharactersLimit",
		OwnerID: "550e8400-e29b-41d4-a716-446655440000",
	}
	reqBody, _ := json.Marshal(req)
	httpReq := httptest.NewRequest(http.MethodPost, "/pet-name", bytes.NewReader(reqBody))
	httpReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(httpReq, rec)

	// Test
	v := validator.New(validator.WithRequiredStructEnabled())
	vw := validatorwrapper.NewValidatorWrapper(v)
	h := NewPetNameHandler(vw)
	err := h.ValidatePetName(c)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var resp Response
	json.Unmarshal(rec.Body.Bytes(), &resp)
	assert.False(t, resp.IsOK)
	assert.Equal(t, badRequestNotValid, resp.Msg)
}

func TestPetNameHandler_ValidatePetName_InvalidOwnerID(t *testing.T) {
	// Setup
	e := echo.New()
	req := PetNameRequest{
		PetName: "Buddy",
		OwnerID: "not-a-uuid",
	}
	reqBody, _ := json.Marshal(req)
	httpReq := httptest.NewRequest(http.MethodPost, "/pet-name", bytes.NewReader(reqBody))
	httpReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(httpReq, rec)

	// Test
	v := validator.New(validator.WithRequiredStructEnabled())
	vw := validatorwrapper.NewValidatorWrapper(v)
	h := NewPetNameHandler(vw)
	err := h.ValidatePetName(c)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var resp Response
	json.Unmarshal(rec.Body.Bytes(), &resp)
	assert.False(t, resp.IsOK)
	assert.Equal(t, badRequestNotValid, resp.Msg)
}

func TestPetNameHandler_ValidatePetName_MissingOwnerID(t *testing.T) {
	// Setup
	e := echo.New()
	req := PetNameRequest{
		PetName: "Buddy",
	}
	reqBody, _ := json.Marshal(req)
	httpReq := httptest.NewRequest(http.MethodPost, "/pet-name", bytes.NewReader(reqBody))
	httpReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(httpReq, rec)

	// Test
	v := validator.New(validator.WithRequiredStructEnabled())
	vw := validatorwrapper.NewValidatorWrapper(v)
	h := NewPetNameHandler(vw)
	err := h.ValidatePetName(c)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusBadRequest, rec.Code)

	var resp Response
	json.Unmarshal(rec.Body.Bytes(), &resp)
	assert.False(t, resp.IsOK)
	assert.Equal(t, badRequestNotValid, resp.Msg)
}

func TestPetNameHandler_ValidatePetName_InternalError(t *testing.T) {
	// Setup
	e := echo.New()
	req := PetNameRequest{
		PetName: "Buddy",
		OwnerID: "550e8400-e29b-41d4-a716-446655440000",
	}
	reqBody, _ := json.Marshal(req)
	httpReq := httptest.NewRequest(http.MethodPost, "/pet-name", bytes.NewReader(reqBody))
	httpReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(httpReq, rec)

	// Test - Mock validator that returns a non-validation error
	mockV := &mockValidator{err: errors.New("unexpected internal error")}
	h := NewPetNameHandler(mockV)
	err := h.ValidatePetName(c)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusInternalServerError, rec.Code)

	var resp Response
	json.Unmarshal(rec.Body.Bytes(), &resp)
	assert.False(t, resp.IsOK)
	assert.Equal(t, "internal server error", resp.Msg)
}

func TestNewPetNameHandler(t *testing.T) {
	v := validator.New(validator.WithRequiredStructEnabled())
	vw := validatorwrapper.NewValidatorWrapper(v)
	h := NewPetNameHandler(vw)
	assert.NotNil(t, h)
	assert.NotNil(t, h.v)
}
