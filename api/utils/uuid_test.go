package utils

import (
	"testing"
)

func Test_NewUUID(t *testing.T){
	uuid ,err:= NewUUID()
	if err != nil{
		t.Errorf("Error of NewUUID:%v",err)

	}
	t.Log(uuid)
}
