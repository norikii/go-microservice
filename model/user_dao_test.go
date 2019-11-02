package model

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetUser_NoUserFound(t *testing.T) {
	// Initialization:

	// Execution:
	user, err := GetUser(0)

	// Validation:
	assert.Equal(t, http.StatusNotFound, err.StatusCode)
	assert.Equal(t, "not_found", err.Status)
	assert.Equal(t, "user 0 was not found", err.Message)

	assert.Nil(t, user, "user with user_id 0 does not exists")
	assert.NotNil(t, err, "user with user_id 0 does not exists, therefore error is required")
}

func TestGetUser_NoError(t *testing.T) {
	user, err := GetUser(123)

	assert.Nil(t, err)
	assert.NotNil(t, user)

	assert.EqualValues(t, 123, user.ID)
	assert.EqualValues(t, "Bob", user.FirstName)
	assert.EqualValues(t, "Wise", user.LastName)
	assert.EqualValues(t, "b.wise@email.com", user.Email)
}