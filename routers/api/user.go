package api

import (
	"github.com/gin-gonic/gin"
	"go-cloud/internal/model"
	"go-cloud/internal/service/userService"
	"go-cloud/pkg/app"
	error2 "go-cloud/pkg/error"
	"go-cloud/pkg/logger"
	"go-cloud/pkg/response"
)

type RegisterUser struct {
	RegisterName string `form:"register_name" binding:"required,min=3,max=10"`
	Password     string `form:"password" binding:"required,min=8,max=20"`
	RePwd        string `form:"re_password" binding:"required,min=8,max=20"`
}
type LoginUser struct {
	LoginName string `form:"login_name" binding:"required,min=3,max=10"`
	Password  string `form:"password" binding:"required,min=8,max=20"`
}

//QQ扫码登陆
//Register 注册
func Register(c *gin.Context) {

}

//Login 登陆
func Login(c *gin.Context) {
	var user LoginUser
	var err error
	//参数校验
	valid, err := app.BindAndValid(c, &user)
	if !valid {
		logger.StdLog().Errorf(c, "app.BindAndValid err :%v", err)
		response.RespError(c, error2.InvalidParams)
		return
	}
	//根据用户名查找是否存在
	u, err := model.QueryUserByWhere("user_name = ?", user.LoginName)
	if err != nil {
		logger.StdLog().Fatal(c, "Can Not Find User :", user.LoginName)
		response.RespError(c, error2.UserNotExist)
		return
	}
	//检验密码
	//登陆成功
	userService.DoLogin(c, u)

}
