package test

import (
	"go-cloud/conf"
	"go-cloud/internal/model"
	"log"
	"testing"
)

func Test_User(t *testing.T) {
	_ = conf.InitSettings()
	_ = model.InitMySQLConn()
	u := &model.User{
		UserName: "A",
		Password: "123123",
	}
	err := u.Create()
	if err != nil {
		log.Println(err)
		return
	}
	u1, _ := model.QueryUserByWhere("user_name = ?", u.UserName)
	u2 := model.GetUserByID(u1.ID)
	log.Println(u1)
	log.Println(u2)
	flag, _ := u1.CheckPassword("123123")
	log.Println(flag)
}
