package app

import (
	"github.com/tatrasoft/go-microservice/controller"
	"net/http"
)

func StartApp() {
	http.HandleFunc("/users", controller.GetUser)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
