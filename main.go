package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

type Person struct {
	Name string `json:"name"`
	Pass string `json:"pass"`
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	app.Post("/user", func (c *fiber.Ctx) error {
		p := new(Person)

		if err := c.BodyParser(p); err != nil {
			return err
		}

		log.Println(p.Name)
		log.Println(p.Pass)

		return nil
	})

	app.Listen(":3000")
}
