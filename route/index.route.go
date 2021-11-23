package route

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-crud/middleware"
	"go-fiber-crud/src/book"
	"go-fiber-crud/src/user"
	"go-fiber-crud/utils"
)

func RoutesInit(app *fiber.App) {
	// Init route static asset
	app.Static("/asset", "./public")
	app.Static("/images", "./public/images")

	// Route Prefix
	apiRoute := app.Group("/api")
	apiV1 := apiRoute.Group("/v1", func(c *fiber.Ctx) error {
		// Set Header Response
		c.Set("Version", "v1")
		c.Set("Callback-Token", "some-token-here")
		return c.Next()
	})

	// API V1 SCOPE
	// ::::::::::::

	// Grouping Route Using Prefix
	userRoute := apiV1.Group("/user")
	bookRoute := apiV1.Group("/book")

	// Route Main
	app.Get("/", helloWorld)

	// Route User V1
	userRoute.Get("/all", user.GetUsers)
	userRoute.Get("/:id", user.GetUser)
	userRoute.Post("/", middleware.AuthMiddleware, user.SaveUser)
	userRoute.Delete("/:id", user.DeleteUser)
	userRoute.Put("/:id", user.UpdateUser)

	// Route BOOK V1
	bookRoute.Get("/all", book.GetBooks)
	bookRoute.Get("/:id", book.GetBook)
	bookRoute.Post("/", middleware.AuthMiddleware, utils.HandleFileUpload, book.SaveBook)
	bookRoute.Delete("/:id", book.DeleteBook)
	bookRoute.Put("/:id", middleware.AuthMiddleware, utils.HandleFileUpload, book.UpdateBook)

	// ::::::::::::
}

func helloWorld(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"hello": "world",
	})
}
