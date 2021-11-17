package user

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-crud/database"
)

func GetUsers(c *fiber.Ctx) error {
	var users []User
	database.DB.Find(&users)
	return c.JSON(FormatUsers(users))
}

func GetUser(c *fiber.Ctx) error {
	id := c.Params("id")

	var user User
	database.DB.Find(&user, id)

	if user.Email == "" {
		return c.Status(404).SendString("User not found.")
	}

	return c.JSON(FormatUser(user))
}

func SaveUser(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	database.DB.Create(&user)
	return c.JSON(FormatUser(*user))
}

func UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")

	var user User
	database.DB.First(&user, id)

	if user.Email == "" {
		return c.Status(404).SendString("User Not Found.")
	}

	if err:= c.BodyParser(&user); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	database.DB.Save(&user)
	return c.JSON(FormatUser(user))
}

func DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	var user User
	database.DB.First(&user, id)

	if user.Email == "" {
		return c.Status(404).SendString("User not found")
	}

	database.DB.Delete(&user)
	return c.SendString("user was deleted")
}