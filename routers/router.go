package routers

import (
	"github.com/gin-gonic/gin"
	"go-cloud/conf"
	"go-cloud/routers/api"
)

func NewRouters() *gin.Engine {
	gin.SetMode(conf.ServerSetting.RunMode)
	r := gin.New()
	//
	r.POST("/register", api.Register)
	r.POST("/login", api.Login)
	return r
}
