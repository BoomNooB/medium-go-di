package main

import (
	"log"

	"github.com/BoomNooB/medium-go-di/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	log.Println("Starting application...")

	// Initialize handler and Echo server
	h := handler.NewHandler()
	e := echo.New()
	e.Use(middleware.Recover())
	e.POST("/api/v1/favorite", h.Favorite)

	if err := e.Start(":1323"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
