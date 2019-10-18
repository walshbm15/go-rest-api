package router

import (
	"github.com/julienschmidt/httprouter"
	"github.com/walshbm15/go-rest-api/logger"
	"github.com/walshbm15/go-rest-api/routes"
)

//Reads from the routes slice to translate the values to httprouter.Handle
func NewRouter(routes routes.Routes) *httprouter.Router {

	router := httprouter.New()
	for _, route := range routes {
		var handle httprouter.Handle

		handle = route.HandlerFunc
		handle = logger.Logger(handle)

		router.Handle(route.Method, route.Path, handle)
	}

	return router
}