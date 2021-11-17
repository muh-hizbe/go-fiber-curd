package migration

import (
	"fmt"
	"go-fiber-crud/book"
	"go-fiber-crud/database"
	"go-fiber-crud/user"
)

func RunMigration() {
	database.DB.AutoMigrate(&user.User{})
	database.DB.AutoMigrate(&book.Book{})

	fmt.Println("Database Migrated")
}
