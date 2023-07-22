package handlers

import (
	"catalogue/internal/service"

	"github.com/gofiber/fiber/v2"
)

func SetupProductRoutes(group *fiber.Group) {
	productRoutes := group.Group("/products")

	productRoutes.Get("/", func(c *fiber.Ctx) error {
		return service.GetProducts(c)
	})

	productRoutes.Get("/:id", func(c *fiber.Ctx) error {
		return service.GetProduct(c)
	})

	productRoutes.Post("/", func(c *fiber.Ctx) error {
		return service.CreateProduct(c)
	})

	productRoutes.Put("/:id", func(c *fiber.Ctx) error {
		return service.UpdateProduct(c)
	})

	productRoutes.Delete("/:id", func(c *fiber.Ctx) error {
		return service.DeleteProduct(c)
	})

}
