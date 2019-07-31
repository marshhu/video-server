package dbops

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

func testUserWorkFlow(t *testing.T) {
	t.Run("Add", testAddUser)
	t.Run("Get", testGetUser)
	t.Run("Del", testDeleteUser)
	t.Run("ReGet", testReGetUser)
}

func testAddUser(t *testing.T) {
	err := AddUserCredential("alice", "123")
	if err != nil {
		t.Errorf("Error of AddUser:%v", err)
	}
}

func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential("alice")
	if pwd != "123" || err != nil {
		t.Errorf("Error of GetUser")
	}
}

func testDeleteUser(t *testing.T) {
	err := DeleteUser("alice", "123")
	if err != nil {
		t.Errorf("Error of DeleteUser:%v", err)
	}
}

func testReGetUser(t *testing.T) {
	pwd, err := GetUserCredential("alice")
	if err != nil {
		t.Errorf("Error of ReGetUser:%v", err)
	}
	if pwd != "" {
		t.Errorf("Deleting user test failed")
	}
}
var tempVid string

func testVideoInfoWorkFlow(t *testing.T){
	clearTables()
	t.Run("PrepareUser",testAddUser)
	t.Run("AddVideo", testAddVideoInfo)
	t.Run("GetVideo", testGetVideoInfo)
	t.Run("DelVideo", testDeleteVideoInfo)
	t.Run("ReGetVideo", testReGetVideoInfo)
}

func testAddVideoInfo(t *testing.T){
	vid,err := AddNewVideo(1, "my-video")
	if err != nil || vid == nil {
		t.Errorf("Error of AddVideoInfo:%v", err)
	}
	tempVid = vid.Id
}

func testGetVideoInfo(t *testing.T){
	_,err := GetVideoInfo(tempVid)
	if err != nil {
		t.Errorf("Error of GetVideoInfo:%v", err)
	}
}

func testDeleteVideoInfo(t *testing.T){
	err := DeleteVideoInfo(tempVid)
	if err != nil {
		t.Errorf("Error of DeleteVideoInfo:%v", err)
	}
}

func testReGetVideoInfo(t *testing.T){
	vid,err := GetVideoInfo(tempVid)
	if err != nil {
		t.Errorf("Error of ReGetVideoInfo:%v", err)
	}
	if vid != nil {
		t.Errorf("Deleting VideoInfo test failed")
	}
}

func TestComments(t *testing.T){
	clearTables()
	t.Run("AddUser",testAddUser)
	t.Run("AddComments",testAddComments)
	t.Run("ListComments",testListComments)
}

func testAddComments(t *testing.T){
	vid :="1111"
	aid :=1
	content :="this is a test video"
	err := AddNewComments(vid,aid,content)
	if err != nil{
		t.Errorf("Error of AddComments: %v",err)
	}
}

func testListComments(t *testing.T){
	vid :="1111"
	from,_ := strconv.Atoi(strconv.FormatInt(time.Now().AddDate(0,0,-2).UnixNano() / 1e9,10))
	to,_ :=  strconv.Atoi(strconv.FormatInt(time.Now().UnixNano() / 1e9,10))
	res,err := ListComments(vid,from,to)
	if err != nil{
		t.Errorf("Error of ListComments: %v",err)
	}
	for i,ele := range res{
		fmt.Printf("comments: %d, %v \n",i,ele)
	}
}