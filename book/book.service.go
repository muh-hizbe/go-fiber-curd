package book

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-crud/database"
)

func GetBooks(c *fiber.Ctx) error {
	var books []Book
	database.DB.Find(&books)
	return c.JSON(FormatBooks(books))
}

func GetBook(c *fiber.Ctx) error {
	id := c.Params("id")

	var book Book
	database.DB.Find(&book, id)

	if book.Title == "" {
		return c.Status(404).SendString("Book not found.")
	}

	return c.JSON(FormatBook(book))
}

func SaveBook(c *fiber.Ctx) error {
	book := new(Book)
	if err := c.BodyParser(book); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	database.DB.Create(&book)
	return c.JSON(FormatBook(*book))
}

func UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")

	var book Book
	database.DB.First(&book, id)

	if book.Title == "" {
		return c.Status(404).SendString("Book Not Found.")
	}

	if err:= c.BodyParser(&book); err != nil {
		return c.Status(500).SendString(err.Error())
	}

	database.DB.Save(&book)
	return c.JSON(FormatBook(book))
}

func DeleteBook(c *fiber.Ctx) error {
	id := c.Params("id")

	var book Book
	database.DB.First(&book, id)


	if book.Title == "" {
		return c.Status(404).SendString("Book not found")
	}
	database.DB.Delete(&book)
	return c.SendString("book was deleted")
}