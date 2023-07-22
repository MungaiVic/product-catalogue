package service

import "github.com/gofiber/fiber/v2"

func GetProducts(c *fiber.Ctx) error {
	return c.SendString("Hello, Market 👋!")
}

func GetProduct(c *fiber.Ctx) error {
	return c.SendString("Get a product here!")
}

func CreateProduct(c *fiber.Ctx) error {
	return c.SendString("Create a product here 👋!")
}

func UpdateProduct(c *fiber.Ctx) error {
	return c.SendString("Update a product here 👋!")
}

func DeleteProduct(c *fiber.Ctx) error {
	return c.SendString("Product will be deleted here 👋!")
}
