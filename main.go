package main

import (
	"net/http"

	"demo/api"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server.
// @host localhost:8080
// @basePath /api
func main() {
	http.HandleFunc("/", api.GetGeneric)
	http.HandleFunc("/multi", api.GetGenericMulti)
	http.ListenAndServe(":8080", nil)
}
