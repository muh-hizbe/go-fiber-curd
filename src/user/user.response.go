package user

import (
	"time"
)

type UserResponse struct {
	ID        int       `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FormatUser(user User) UserResponse {
	formatter := UserResponse{}
	formatter.ID = int(user.ID)
	formatter.FirstName = user.FirstName
	formatter.LastName = user.LastName
	formatter.Email = user.Email
	formatter.CreatedAt = user.CreatedAt
	formatter.UpdatedAt = user.UpdatedAt

	return formatter
}

func FormatUsers(users []User) []UserResponse {
	if len(users) == 0 {
		return []UserResponse{}
	}

	var usersFormatter []UserResponse

	for _, user := range users {
		formatter := FormatUser(user)
		usersFormatter = append(usersFormatter, formatter)
	}

	return usersFormatter
}
