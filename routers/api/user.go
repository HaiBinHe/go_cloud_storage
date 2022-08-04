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
}
type LoginUser struct {
	LoginName string `form:"login_name" binding:"required,min=3,max=10"`
	Password  string `form:"password" binding:"required,min=8,max=20"`
}

//QQ扫码登陆
func QQScanLogin(c *gin.Context) {

}

//Register 注册
func Register(c *gin.Context) {
	var ru RegisterUser
	var err error
	//参数校验
	valid, err := app.BindAndValid(c, &ru)
	if !valid {
		logger.StdLog().Errorf(c, "app.BindAndValid err :%v", err)
		response.RespError(c, error2.InvalidParams)
		return
	}
	//根据用户名查找是否存在
	_, err = model.QueryUserByWhere("user_name = ?", ru.RegisterName)
	//用户名已存在
	if err == nil {
		logger.StdLog().Error(c, "User Exist :", ru.RegisterName)
		response.RespError(c, error2.UserExist)
		return
	}
	u := model.User{
		UserName: ru.RegisterName,
		Password: ru.Password,
	}
	//创建用户
	err = u.Create()
	if err != nil {
		logger.StdLog().Error(c, "创建用户出错")
		response.RespError(c, error2.ServerError)
		return
	}
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
		logger.StdLog().Error(c, "Can Not Find User :", user.LoginName)
		response.RespError(c, error2.UserNotExist)
		return
	}
	//检验密码
	flag, err := u.CheckPassword(user.Password)
	//解析密码错误
	if err != nil {
		logger.StdLog().Error(c, err)
		response.RespError(c, err)
		return
	}
	//密码错误
	if !flag {
		logger.StdLog().Error(c, "Invalid Password")
		response.RespError(c, error2.ErrorPassword)
		return
	}
	//登陆成功
	userService.DoLogin(c, u)
}
