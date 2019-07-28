package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"video-server/handler"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user", handler.CreateUser)
	router.POST("/user/:user_name", handler.Login)
	return router
}

func main() {
	router := RegisterHandlers()
	http.ListenAndServe(":8000", router)
}
