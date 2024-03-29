package dbops

import (
	"database/sql"
	"log"
	"strconv"
	"sync"
	"video-server/api/defs"
)

func InsertSession(sid string,ttl int64,uname string) error{
	ttlStr := strconv.FormatInt(ttl,10)
	stmtIns,err := dbConn.Prepare("insert into sessions(session_id,TTL,login_name) values(?,?,?);")
	if err != nil{
		return  err
	}
	_,err = stmtIns.Exec(sid,ttlStr,uname)
	if err!= nil{
		return  err
	}
	defer stmtIns.Close()
	return  nil
}

func RetrieveSession(sid string) (*defs.SimpleSession,error){
	ss := &defs.SimpleSession{}
	stmtOut,err := dbConn.Prepare("select TTL,login_name from sessions where session_id =?;")
	if err != nil{
		return  nil,err
	}
	var ttl string
	var uname string
	err = stmtOut.QueryRow(sid).Scan(&ttl,&uname)
	if err != nil && err != sql.ErrNoRows{
		return nil, err
	}
	if err == sql.ErrNoRows{
		return nil,nil
	}
    if res, err := strconv.ParseInt(ttl,10,64);err == nil {
		ss.TTL = res
		ss.Username = uname
	}else {
		return  nil,err
	}
    defer stmtOut.Close()
	return ss,nil
}

func RetrieveAllSessions()(*sync.Map,error){
	m := &sync.Map{}
	stmtOut,err := dbConn.Prepare("select TTL,login_name from sessions;")
	if err != nil{
		return  nil,err
	}
	rows,err := stmtOut.Query()
	if err != nil && err != sql.ErrNoRows{
		return nil, err
	}
	if err == sql.ErrNoRows{
		return nil,nil
	}
	for rows.Next(){
       var id string
       var ttlStr string
       var loginName string
       if er := rows.Scan(&id,&ttlStr,&loginName);er != nil{
       	  break
	   }
       if ttl,err1 := strconv.ParseInt(ttlStr,10,64);err1 == nil{
       	   ss := &defs.SimpleSession{Username: loginName,TTL:ttl}
       	   m.Store(id,ss)
	   }
	}
	defer stmtOut.Close()
	return m,nil
}

func DeleteSession(sid string) error{
	stmtDel, err := dbConn.Prepare("delete from sessions where session_id =?;")
	if err != nil {
		log.Printf("delete session error:%s", err)
		return err
	}
	_,err = stmtDel.Exec(sid)
	if err != nil{
		return  err
	}
	defer stmtDel.Close()
	return nil
}
