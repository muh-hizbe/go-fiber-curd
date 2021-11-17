package main

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-crud/book"
	"go-fiber-crud/database"
	"go-fiber-crud/database/migration"
	"go-fiber-crud/user"
)

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello world")
}

func Routers(app *fiber.App) {
	apiRoute := app.Group("/api")
	apiV1 := apiRoute.Group("/v1")

	userRoute := apiV1.Group("/user")
	bookRoute := apiV1.Group("/book")

	// Route User V1
	userRoute.Get("/all", user.GetUsers)
	userRoute.Get("/:id", user.GetUser)
	userRoute.Post("/", user.SaveUser)
	userRoute.Delete("/:id", user.DeleteUser)
	userRoute.Put("/:id", user.UpdateUser)

	// Route BOOK V1
	bookRoute.Get("/all", book.GetBooks)
	bookRoute.Get("/:id", book.GetBook)
	bookRoute.Post("/", book.SaveBook)
	bookRoute.Delete("/:id", book.DeleteBook)
	bookRoute.Put("/:id", book.UpdateBook)
}


func main() {
	database.InitDatabase()
	migration.RunMigration()

	app := fiber.New()
	app.Use(func(c *fiber.Ctx) error {
		c.Set("X-XSS-Protection", "1; mode=block")
		c.Set("X-Content-Type-Options", "nosniff")
		c.Set("X-Download-Options", "noopen")
		c.Set("Strict-Transport-Security", "max-age=5184000")
		c.Set("X-Frame-Options", "SAMEORIGIN")
		c.Set("X-DNS-Prefetch-Control", "off")

		return c.Next()
	})
	app.Get("/", helloWorld)
	Routers(app)

	app.Listen(":3000")
}
