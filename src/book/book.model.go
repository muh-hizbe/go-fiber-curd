package book

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title     string
	Image     string
	Quantity  int
	Author    string
	Publisher string
}
