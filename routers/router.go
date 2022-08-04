package routers

import (
	"github.com/gin-gonic/gin"
	"go-cloud/conf"
)

func NewRouters() *gin.Engine {
	gin.SetMode(conf.ServerSetting.RunMode)
	r := gin.New()
	return r
}
