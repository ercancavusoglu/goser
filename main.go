package main

import (
	"fmt"
	"github.com/ercancavusoglu/goser/database"
	"github.com/gofiber/fiber/v2"
	"log"
)

type Todo struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

var todos = []Todo{
	{Id: 1, Name: "Walk", Completed: false},
	{Id: 2, Name: "Sleep", Completed: false},
}

func welcome(c *fiber.Ctx) error {
	fmt.Println("🥉 Son handler")
	return c.SendString("Merhaba, Dünya 👋!")
}

func main() {
	database.ConnectDb()

	app := fiber.New()

	app.Get("/api", welcome)

	log.Fatal(app.Listen(":3000"))

}
