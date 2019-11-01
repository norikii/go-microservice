package services

import (
	"github.com/tatrasoft/go-microservice/model"
	"github.com/tatrasoft/go-microservice/utils"
)

func GetUser(userID int64) (*model.User, *utils.AppError) {
	return model.GetUser(userID, )
}

