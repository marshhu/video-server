package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"video-server/streamserver/handler"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.GET("/videos/:vid-id",handler.StreamHandler)
	router.POST("/upload/:vid-id",handler.UploadHandler)
	return router
}
func main() {
	r := RegisterHandlers()
	http.ListenAndServe(":9000",r)
}
