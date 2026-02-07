package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/BoomNooB/medium-go-di/validatorwrapper"
	"github.com/labstack/echo/v4"
)

type GuessCatNameRequest struct {
	GuessName string `json:"guessName" validate:"required,min=1,max=30"`
	UserID    string `json:"userId" validate:"required,uuid_rfc4122"`
	Attempts  int    `json:"attempts" validate:"required,gte=1,lte=3"`
}

type GuessCatNameHandler struct {
	v Valiator
}

func NewGuessCatNameHandler(validator Valiator) *GuessCatNameHandler {
	return &GuessCatNameHandler{
		v: validator,
	}
}

func (gh *GuessCatNameHandler) GuessTheCatName(c echo.Context) error {
	ctx := c.Request().Context()
	req := GuessCatNameRequest{}
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			newBadRequestResponse(badRequestJSONSyntax),
		)
	}

	err = gh.v.StructValidation(ctx, &req)
	if err != nil {
		// check if it's a validation error or not
		if errors.Is(err, validatorwrapper.ErrValidationFailed) {
			return c.JSON(
				http.StatusBadRequest,
				newBadRequestResponse(badRequestNotValid),
			)
		}

		// else it's an internal error
		log.Printf("Guess Name: %s, User ID: %s, Attempts: %d\n", req.GuessName, req.UserID, req.Attempts)
		log.Printf("Error: %v\n", err)
		return c.JSON(
			http.StatusInternalServerError,
			newInternalErrorResponse(),
		)
	}

	log.Println("Yay! Valid guess for cat name!")
	return c.JSON(
		http.StatusOK,
		newOkResponse(),
	)
}
