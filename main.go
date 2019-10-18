package main

import (
	"log"
	"net/http"

	"github.com/walshbm15/go-rest-api/router"
	"github.com/walshbm15/go-rest-api/routes"
)

func main() {
	router := router.NewRouter(routes.AllRoutes())
	log.Fatal(http.ListenAndServe(":8080", router))
}
