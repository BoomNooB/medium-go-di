package main

import (
	"log"

	"github.com/BoomNooB/medium-go-di/handler"
	"github.com/BoomNooB/medium-go-di/validatorwrapper"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	log.Println("Starting application...")

	// Initialize validator once (DI)
	v := validator.New(validator.WithRequiredStructEnabled())
	vWrapper := validatorwrapper.NewValidatorWrapper(v)

	// Initialize all handlers with the same validator instance (DI)
	favHandler := handler.NewFavoriteNumHandler(vWrapper)
	petNameHandler := handler.NewPetNameHandler(vWrapper)
	thaiCIDHandler := handler.NewThaiCIDHandler(vWrapper)
	guessCatHandler := handler.NewGuessCatNameHandler(vWrapper)

	// Setup Echo server
	e := echo.New()
	e.Use(middleware.Recover())

	// Register all routes
	e.POST("/api/v1/favorite", favHandler.Favorite)
	e.POST("/api/v1/pet-name", petNameHandler.ValidatePetName)
	e.POST("/api/v1/thai-cid", thaiCIDHandler.ValidateThaiCID)
	e.POST("/api/v1/guess-cat", guessCatHandler.GuessTheCatName)

	if err := e.Start(":1323"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
