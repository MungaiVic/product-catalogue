package main

import (
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

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Market 👋!")
	})

	app.Use(logger.New(logger.Config{
		Format:     "${cyan}[${time}] ${status} - ${method} ${path}\n",
		TimeZone:   "Africa/Nairobi",
		TimeFormat: "2006-01-02 15:04:05",
	}))

	log.Fatal(app.Listen(fmt.Sprintf(":%v", os.Getenv("APP_PORT"))))
}
