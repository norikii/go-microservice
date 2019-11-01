package main

import "net/http"

func main() {
	http.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Hello world from go server"))
	})

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		panic(err)
	}
}
