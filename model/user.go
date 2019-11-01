package model

// named collection of attributes for User model
type User struct {
	ID uint64 `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
}
