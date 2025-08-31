package main

import (
	"log"

	loader "github.com/newOKC2010/go-fiber-public/src/config"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	loader.LoadEnv()

	app.Get("/", func(c *fiber.Ctx) error {
		greeting := loader.FormatGreeting("ผู้ใช้")
		return c.SendString(greeting)
	})

	app.Get("/age/:year", func(c *fiber.Ctx) error {
		year, _ := c.ParamsInt("year")
		age := loader.CalculateAge(year)
		return c.JSON(map[string]int{"age": age})
	})

	log.Printf("Database: %s", loader.GetDatabaseURL())
	log.Printf("JWT Secret: %s", loader.GetJWTSecret())
	log.Printf("Starting server on port: %s", loader.LoadPort())
	app.Listen(":" + loader.LoadPort())
}
