package model

import (
	"fmt"
	"github.com/tatrasoft/go-microservice/utils"
	"net/http"
)

var users = map[int64]*User {
	123: {
		ID:        123,
		FirstName: "Bob",
		LastName:  "Wise",
		Email:     "b.wise@email.com",
	},
}

func GetUser(userID int64) (*User, *utils.AppError) {
	// here we should be connecting to the database and
	// getting the user from the persistent layer
	user := users[userID]
	if user == nil {
		return nil, &utils.AppError{
			Message:    fmt.Sprintf("user %v was not found", userID),
			Status:     "not_found",
			StatusCode: http.StatusNotFound,
		}
	}

	return user, nil
}
