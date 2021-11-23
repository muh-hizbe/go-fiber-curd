package book

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go-fiber-crud/database"
	"go-fiber-crud/utils"
	"log"
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

	// CONVERT FILENAME TO STRING
	filename := fmt.Sprintf("%v", c.Locals("filename"))
	// ADD IMAGE TO THE BOOK.IMAGE
	book.Image = filename

	err := database.DB.Create(&book).Error

	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"status":  "Error",
			"message": "Failed to create a new book",
		})
	}

	fmt.Println("Form value ::: ", c.FormValue("title"))
	fmt.Println("Form value ::: ", c.FormValue("author"))
	fmt.Println("Form value ::: ", c.FormValue("publisher"))
	fmt.Println("Form value ::: ", c.Locals("filename"))
	fmt.Println("Form value ::: ", *book)
	return c.JSON(FormatBook(*book))
}

func UpdateBook(c *fiber.Ctx) error {
	id := c.Params("id")

	var book Book
	database.DB.First(&book, id)

	if book.Title == "" {
		return c.Status(404).SendString("Book Not Found.")
	}

	if err := c.BodyParser(&book); err != nil {
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

	file := utils.File{
		Name: book.Image,
	}

	result, err := utils.HandleRemoveFile(file)
	if err != nil {
		log.Println("Fail remove file. hehe :)")
	} else {
		log.Println("File was deleted: ", result.Name, " In Path", result.Path)
	}
	database.DB.Delete(&book)
	return c.SendString("book was deleted")
}
