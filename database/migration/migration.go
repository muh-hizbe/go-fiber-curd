package migration

import (
	"fmt"
	"go-fiber-crud/database"
	"go-fiber-crud/src/book"
	"go-fiber-crud/src/user"
	"log"
)

func RunMigration() {
	err := database.DB.AutoMigrate(&user.User{})
	if err != nil {
		log.Println(err)
	}
	err1 := database.DB.AutoMigrate(&book.Book{})
	if err1 != nil {
		log.Println(err1)
	}

	fmt.Println("Database Migrated")
}
