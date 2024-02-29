package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/rest/fiber-go/database"
	"github.com/rest/fiber-go/routes"
)
func welcome ( c *fiber.Ctx ) error {
	return c.SendString("welcome to this api")
}

func setupRoutes(app*fiber.App) {
	app.Get("/api",welcome)
	app.Post("/api/users",routes.CreateUser)
	app.Get("app/users",routes.GetUsers)
}




func main () {
	database.ConnectDb()
	app := fiber.New()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}

