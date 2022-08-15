package routers

import (
	"github.com/gin-gonic/gin"
	"go-cloud/conf"
	"go-cloud/internal/middleware"
	"go-cloud/routers/api"
	"net/http"
)

func NewRouters() *gin.Engine {
	gin.SetMode(conf.ServerSetting.RunMode)
	r := gin.New()
	//中间件
	//跨域
	r.Use(middleware.Cors())
	//翻译
	r.Use(middleware.Translations())
	r.POST("/register", api.Register)
	r.POST("/login", api.Login)
	//设置静态文件
	r.StaticFS("/static", http.Dir(conf.AppSetting.UploadSavePath))
	apiv1 := r.Group("/api/v1")
	apiv1.Use(middleware.JWT())
	{
		apiv1.GET("/user/:id", nil)
		//文件分享
		share := apiv1.Group("share")
		share.Use(middleware.ShareOwner())
		share.Use(middleware.ShareAvailable())
		{
		}
	}

	return r
}
