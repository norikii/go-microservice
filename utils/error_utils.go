package utils

type AppError struct {
	Message string `json:"message"`
	Status string `json:"status"`
	StatusCode int `json:"status_code"`
}
