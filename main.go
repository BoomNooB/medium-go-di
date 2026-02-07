package main

import (
	"log"

	"github.com/BoomNooB/medium-go-di/handler"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	log.Println("Starting application...")

	// Initialize validator once (DI)
	v := validator.New(validator.WithRequiredStructEnabled())

	// Initialize all handlers with the same validator instance (DI)
	favHandler := handler.NewFavoriteNumHandler(v)
	petNameHandler := handler.NewPetNameHandler(v)
	thaiCIDHandler := handler.NewThaiCIDHandler(v)
	guessCatHandler := handler.NewGuessCatNameHandler(v)

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
