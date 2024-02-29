package routes

import (
	
	"github.com/gofiber/fiber/v2"
	"github.com/rest/fiber-go/database"
	"github.com/rest/fiber-go/models"
)

type User struct {
	//this is not a model see it as a serializer
	ID        uint    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func CreateResponseUser(user models.User) User {
	return User{ ID: user.ID, FirstName: user.FirstName, LastName: user.LastName}

}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	if err := c.BodyParser(&user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&user)
	responseUser := CreateResponseUser(user)
	return c.Status(200).JSON(responseUser)
}


func GetUsers(c *fiber.Ctx) error {
	users := []models.User{}
	database.Database.Db.Find(&users)
	responseUsers := []User{}
	for _, user := range users {
		responseUser := CreateResponseUser(user)
		responseUsers = append(responseUsers, responseUser)
	}

	return c.Status(200).JSON(responseUsers)
}
