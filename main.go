package main

import (
	"auth0Authentication/pkg/handlers"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading the .env file: %v", err)
	}
	e := echo.New()
	e.POST("/api/v1/signup", handlers.Signup)
	e.POST("/api/v1/requestOTP", handlers.RequestOTP)
	e.POST("/api/v1/requestToken", handlers.RequestToken)
	e.Logger.Fatal(e.Start(":8080"))
}
