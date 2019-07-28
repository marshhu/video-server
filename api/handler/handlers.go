package handler

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	io.WriteString(w, "Create User Handler")
}

func Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userName := ps.ByName("user_name")
	io.WriteString(w, userName)
}
