package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type FavoriteNumRequest struct {
	UserID string `json:"userId" validate:"required,uuid_rfc4122"`
	FavNum int    `json:"favNum" validate:"required,gt=0"`
}

type FavoriteNumHandler struct {
	v *validator.Validate
}

func NewFavoriteNumHandler(validator *validator.Validate) *FavoriteNumHandler {
	return &FavoriteNumHandler{
		v: validator,
	}
}

func (h *FavoriteNumHandler) Favorite(c echo.Context) error {
	ctx := c.Request().Context()
	req := FavoriteNumRequest{}
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
		// print both user, fav and error
		log.Printf("User ID: %s, Favorite Number: %d\n", req.UserID, req.FavNum)
		log.Printf("Error: %v\n", err)
		return c.JSON(
			http.StatusInternalServerError,
			newInternalErrorResponse(),
		)
	}

	log.Println("Yay!")
	return c.JSON(
		http.StatusOK,
		newOkResponse(),
	)
}
