package main

import (
	"fmt"
	"log"
	"os"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/thesilentline/shorten-url-project/routes"
)

func setupRoutes( app *fiber.App){
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/v1", routes.ShortenURL)
}

func main() {
	err := godotenv.Load()

	if err!=nil {
		fmt.Println(err)
	}

	app := fiber.New()

	app.Use(logger.New())	//integrate logging into your Go web application using middleware

	setupRoutes(app)

	log.Fatal(app.Listen(os.Getenv("APP_PORT")))

}
