package dao

import (
	"go-cloud/conf"
	"go-cloud/internal/model"
	"testing"
)

func initConn(t *testing.T){
	err := conf.InitSettings()
	if err != nil {
		t.Error(err)
	}
	err = model.InitMySQLConn()
	if err != nil {
		t.Error(err)
	}
}

func TestFileCenter(t *testing.T){
	initConn(t)
	println("isExist:", FileIsExist("aaa"))
}
