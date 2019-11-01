package controller

import (
	"encoding/json"
	"github.com/tatrasoft/go-microservice/services"
	"github.com/tatrasoft/go-microservice/utils"
	"net/http"
	"strconv"
)

const(
	invalidUserIDErrMessage = "user_id must be a number"
	badRequestStatus = "bad_request"
)

func GetUser(resp http.ResponseWriter, req *http.Request){
	userID, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		jsonUserErr, _ := handleError(invalidUserIDErrMessage, badRequestStatus, http.StatusBadRequest)
		resp.WriteHeader(http.StatusBadRequest)
		resp.Write(jsonUserErr)

		return
	}

	user, userErr := services.GetUser(userID)
	if userErr != nil {
		jsonUserErr, _ := handleError(userErr.Message, userErr.Status, userErr.StatusCode)
		resp.WriteHeader(userErr.StatusCode)
		resp.Write(jsonUserErr)

		return
	}

	// return a user to the client
	jsonUser, _ := json.Marshal(user)
	resp.Write(jsonUser)
}

func handleError(message string, status string, statusCode int) ([]byte, error){
	apiErr := &utils.AppError{
		Message:    message,
		Status:     status,
		StatusCode: statusCode,
	}

	return json.Marshal(apiErr)
}
