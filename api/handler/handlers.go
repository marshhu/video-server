package handler

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"net/http"
	"video-server/api/dbops"
	"video-server/api/defs"
	"video-server/api/session"
)

func CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	res,_ := ioutil.ReadAll(r.Body)
	ubody := &defs.UserCredential{}
	if err := json.Unmarshal(res,ubody);err != nil{
		sendErrorResponse(w,defs.ErrorRequestBodyParseFailed)
		return
	}

	if err := dbops.AddUserCredential(ubody.Username,ubody.Pwd);err != nil{
		sendErrorResponse(w,defs.ErrorDBError)
		return
	}

	id := session.GenerateNewSessionId(ubody.Username)
	su := &defs.SignedUp{Success:true,SessionId:id}
	if resp,err := json.Marshal(su);err != nil{
		sendErrorResponse(w,defs.ErrorInternalFault)
		return
	}else{
		sendNormalResponse(w,string(resp),201)
		return
	}
}

func Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userName := ps.ByName("user_name")
	io.WriteString(w, userName)
}
