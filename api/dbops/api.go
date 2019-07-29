package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
	"video-server/api/defs"
	"video-server/api/utils"
)

func AddUserCredential(loginName string, pwd string) error {
	stmtIns, err := dbConn.Prepare("Insert into users(login_name,pwd) values(?,?);")
	if err != nil {
		return err
	}
	_,err = stmtIns.Exec(loginName, pwd)
	if err != nil{
		return  err
	}
	defer stmtIns.Close()
	return nil
}

func GetUserCredential(loginName string) (string, error) {
	stmtOut, err := dbConn.Prepare("select pwd from users where login_name =?;")
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}
	var pwd string
	err = stmtOut.QueryRow(loginName).Scan(&pwd)
	if err !=nil && err != sql.ErrNoRows{
		return "", err
	}

	defer stmtOut.Close()
	return pwd, nil
}

func DeleteUser(loginName string, pwd string) error {
	stmtDel, err := dbConn.Prepare("delete from users where login_name =? and pwd =?;")
	if err != nil {
		log.Printf("delete user error:%s", err)
		return err
	}
	_,err = stmtDel.Exec(loginName, pwd)
	if err != nil{
		return  err
	}
	defer stmtDel.Close()
	return nil
}

func AddNewVideo(authorId int,name string) (*defs.VideoInfo,error){
	vid,err := utils.NewUUID()
	if err != nil{
		return nil,err
	}

	t := time.Now()
	ctime := t.Format("Jan 02 2006,15:04:05")
	stmtIns,err:= dbConn.Prepare("insert into video_info(id,author_id,name,dispaly_ctime,create_time) values(?,?,?,?,?)")
	if err != nil {
		return nil,err
	}
	_,err = stmtIns.Exec(vid, authorId,name,ctime,t)
	if err != nil{
		return nil,err
	}
    res := &defs.VideoInfo{Id:vid,AuthorId:authorId,Name:name,DisplayCtime:ctime}
	defer stmtIns.Close()
    return res,nil
}
