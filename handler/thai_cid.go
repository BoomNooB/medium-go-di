package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type ThaiCIDRequest struct {
	CitizenID string `json:"citizenId" validate:"required,len=13,numeric"`
	FullName  string `json:"fullName" validate:"required,min=3"`
}

type ThaiCIDHandler struct {
	v *validator.Validate
}

func NewThaiCIDHandler(validator *validator.Validate) *ThaiCIDHandler {
	return &ThaiCIDHandler{
		v: validator,
	}
}

func (h *ThaiCIDHandler) ValidateThaiCID(c echo.Context) error {
	ctx := c.Request().Context()
	req := ThaiCIDRequest{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			newBadRequestResponse(badRequestJSONSyntax),
		)
	}

	err = h.v.StructCtx(ctx, &req)
	if err != nil {
		// check if it's a validation error or not
		if errors.As(err, &validator.ValidationErrors{}) {
			return c.JSON(
				http.StatusBadRequest,
				newBadRequestResponse(badRequestNotValid),
			)
		}

		// else it's an internal error
		log.Printf("Citizen ID: %s, Full Name: %s\n", req.CitizenID, req.FullName)
		log.Printf("Error: %v\n", err)
		return c.JSON(
			http.StatusInternalServerError,
			newInternalErrorResponse(),
		)
	}

	log.Println("Yay! Thai Citizen ID is valid!")
	return c.JSON(
		http.StatusOK,
		newOkResponse(),
	)
}
