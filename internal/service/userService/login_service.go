package userService

import (
	"github.com/gin-gonic/gin"
	"go-cloud/conf"
	"go-cloud/internal/model"
	"go-cloud/pkg/auth"
	"go-cloud/pkg/logger"
	"go-cloud/pkg/response"
	"time"
)

//DoLogin 登陆时根据用户信息生成token
func DoLogin(ctx *gin.Context, user model.User) {
	accessExpTime := time.Now().Add(conf.JWTSetting.Expire)
	refreshExpTime := time.Now().Add(conf.JWTSetting.Expire * 4)
	accessToken, err := auth.GenerateToken(user, accessExpTime)
	if err != nil {
		logger.StdLog().Fatal(ctx, err)
		return
	}
	refreshToken, err := auth.GenerateToken(user, refreshExpTime)
	if err != nil {
		logger.StdLog().Fatal(ctx, err)
		return
	}
	response.RespData(
		ctx,
		map[string]string{
			"msg":          "登陆成功",
			"accessToken":  accessToken,
			"refreshToken": refreshToken,
		})
}
