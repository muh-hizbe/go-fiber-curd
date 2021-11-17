package book

import "time"

type BookResponse struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Quantity int `json:"quantity"`
	Author string `json:"author"`
	Publisher string `json:"publisher"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FormatBook(book Book) BookResponse {
	formatter := BookResponse{}
	formatter.ID = int(book.ID)
	formatter.Title = book.Title
	formatter.Quantity = book.Quantity
	formatter.Author = book.Author
	formatter.Publisher = book.Publisher
	formatter.CreatedAt = book.CreatedAt
	formatter.UpdatedAt = book.UpdatedAt

	return formatter
}

func FormatBooks(books []Book) []BookResponse {
	if len(books) == 0 {
		return []BookResponse{}
	}

	var booksFormatter []BookResponse

	for _, book := range books {
		formatter := FormatBook(book)
		booksFormatter = append(booksFormatter, formatter)
	}

	return booksFormatter
}
