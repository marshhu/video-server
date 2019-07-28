package handler

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)

func StreamHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	io.WriteString(w, "streamHandlerr Handler")
}

func UploadHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	io.WriteString(w, "uploadHandler Handler")
}