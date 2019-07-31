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
	if err == sql.ErrNoRows{
		return "",nil
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
	stmtIns,err:= dbConn.Prepare("INSERT INTO video_info(id, author_id, name, display_ctime, create_time) VALUES (?, ?, ?, ?, ?);")
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

func GetVideoInfo(vid string) (*defs.VideoInfo,error){
	stmtOut, err := dbConn.Prepare("select author_id,name,display_ctime from video_info where id = ?;")
	if err != nil {
		log.Printf("%s", err)
		return nil, err
	}
	var aid int
	var name string
	var ctime string
	err = stmtOut.QueryRow(vid).Scan(&aid,&name,&ctime)
	if err != nil && err != sql.ErrNoRows{
		return nil, err
	}
	if err == sql.ErrNoRows{
		return nil,nil
	}

	defer stmtOut.Close()
	res := &defs.VideoInfo{Id:vid,AuthorId:aid,Name:name,DisplayCtime:ctime}
	return res, nil
}

func DeleteVideoInfo(vid string) error{
	stmtDel, err := dbConn.Prepare("delete from video_info where id =?;")
	if err != nil {
		log.Printf("delete user error:%s", err)
		return err
	}
	_,err = stmtDel.Exec(vid)
	if err != nil{
		return  err
	}
	defer stmtDel.Close()
	return nil
}

func AddNewComments(vid string ,aid int ,content string) error{
	id,err := utils.NewUUID()
	if err != nil{
		return  err
	}
	stmtIns,err := dbConn.Prepare("INSERT INTO comments (id, video_id, author_id, content) VALUES (?, ?, ?, ?);")
	if err != nil{
		return  err
	}
	_,err = stmtIns.Exec(id,vid,aid,content)
	if err != nil{
		return  err
	}
	defer stmtIns.Close()
	return nil
}

func ListComments(vid string,from,to int) ([]*defs.Comment,error){
	stmtOut,err := dbConn.Prepare(`select comments.id,users.login_name,comments.content from comments 
	inner join users on comments.author_id = users.id where comments.video_id = ? and comments.time > FROM_UNIXTIME(?) and comments.time <= FROM_UNIXTIME(?);`)
	if err != nil{
		return nil,err
	}
	var res []*defs.Comment
	rows,err :=stmtOut.Query(vid,from,to)
	for rows.Next(){
		var id,name,content string
		if err = rows.Scan(&id,&name,&content);err != nil{
			return res,err
		}
		c := &defs.Comment{Id:id,VideoId:vid,Author:name,Content:content}
		res = append(res,c)
	}
	defer stmtOut.Close()
	return res,nil
}