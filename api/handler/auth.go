package handler

import (
	"net/http"
	"video-server/api/defs"
	"video-server/api/session"
)
var HEADER_FIELD_SESSION ="X-Session-Id"
var HEADER_FILE_UNAME = "X-User-Name"

func ValidateUserSession(r *http.Request) bool{
	sid := r.Header.Get(HEADER_FIELD_SESSION)
	if len(sid) == 0 {
		return  false
	}

	name,ok := session.IsSessionExpired(sid)
	if ok {
		return  false
	}
	r.Header.Add(HEADER_FILE_UNAME,name)
	return  true
}

func ValidateUser(w http.ResponseWriter,r *http.Request) bool{
	name := r.Header.Get(HEADER_FILE_UNAME)
	if len(name) == 0 {
		sendErrorResponse(w,defs.ErrorRequestBodyParseFailed)
		return false
	}
	return true
}
