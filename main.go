package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go-fiber-crud/database"
	"go-fiber-crud/database/migration"
	"go-fiber-crud/route"
)

func main() {
	// Initial database connection
	database.InitDatabase()
	// Running migration db mysql
	migration.RunMigration()

	// Init Go Fiber
	app := fiber.New()
	// Set Cors Config
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET",
	}))
	// Set headers response
	app.Use(func(c *fiber.Ctx) error {
		c.Set("X-XSS-Protection", "1; mode=block")
		c.Set("X-Content-Type-Options", "nosniff")
		c.Set("X-Download-Options", "noopen")
		c.Set("Strict-Transport-Security", "max-age=5184000")
		c.Set("X-Frame-Options", "SAMEORIGIN")
		c.Set("X-DNS-Prefetch-Control", "off")

		return c.Next()
	})

	// Routing List
	route.RoutesInit(app)

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
