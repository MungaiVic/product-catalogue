package service

import "github.com/gofiber/fiber/v2"

type ProductService interface {
	GetProducts(c *fiber.Ctx) error
	GetProduct(c *fiber.Ctx) error
	CreateProduct(c *fiber.Ctx) error
	UpdateProduct(c *fiber.Ctx) error
	DeleteProduct(c *fiber.Ctx) error
}
