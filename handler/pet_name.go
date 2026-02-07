package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/BoomNooB/medium-go-di/validatorwrapper"
	"github.com/labstack/echo/v4"
)

type PetNameRequest struct {
	PetName string `json:"petName" validate:"required,min=2,max=50"`
	OwnerID string `json:"ownerId" validate:"required,uuid_rfc4122"`
}

type PetNameHandler struct {
	v Valiator
}

func NewPetNameHandler(validator Valiator) *PetNameHandler {
	return &PetNameHandler{
		v: validator,
	}
}

func (ph *PetNameHandler) ValidatePetName(c echo.Context) error {
	ctx := c.Request().Context()
	req := PetNameRequest{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			newBadRequestResponse(badRequestJSONSyntax),
		)
	}

	err = ph.v.StructValidation(ctx, &req)
	if err != nil {
		// check if it's a validation error or not
		if errors.Is(err, validatorwrapper.ErrValidationFailed) {
			return c.JSON(
				http.StatusBadRequest,
				newBadRequestResponse(badRequestNotValid),
			)
		}

		// else it's an internal error
		log.Printf("Pet Name: %s, Owner ID: %s\n", req.PetName, req.OwnerID)
		log.Printf("Error: %v\n", err)
		return c.JSON(
			http.StatusInternalServerError,
			newInternalErrorResponse(),
		)
	}

	log.Println("Yay! Pet name is valid!")
	return c.JSON(
		http.StatusOK,
		newOkResponse(),
	)
}
