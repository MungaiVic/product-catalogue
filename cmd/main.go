package main

import (
	"catalogue/internal/handlers"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	envvar := godotenv.Load(".env")
	if envvar != nil {
		log.Fatal("APP_PORT is not set")
	}

	app := fiber.New(fiber.Config{
		Prefork: false,
	})
	// Initialize default config
	app.Use(logger.New(logger.Config{
		Format:     "[${time}] ${status} - ${method} ${path}\n",
		TimeZone:   "Africa/Nairobi",
		TimeFormat: "2006-01-02 15:04:05",
	}))

	v1 := app.Group("/v1").(*fiber.Group)

	handlers.SetupProductRoutes(v1)

	log.Fatal(app.Listen(fmt.Sprintf(":%v", os.Getenv("APP_PORT"))))
}
