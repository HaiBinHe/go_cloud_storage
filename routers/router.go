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

	//设置静态文件
	r.StaticFS("/static", http.Dir(conf.AppSetting.UploadSavePath))
	apiv1 := r.Group("/api/v1")
	{
		apiv1.POST("/register", api.Register)
		apiv1.POST("/login", api.Login)
	}
	apiv1.Use(middleware.JWT())
	{
		// 更新用户信息
		apiv1.GET("/user/:id", nil)
		apiv1.POST("/user/:id", nil)
		// 文件
		apiv1.POST("/upload/file", api.Upload)
		//目录
		directory := apiv1.Group("/directory")
		{
			directory.POST("", api.CreateDir)
			directory.GET("", api.ListDirectories)
		}
		//文件分享
		share := apiv1.Group("share")

		{
			share.POST("", api.CreateShare)
			share.GET("", api.ListrShare)
			share.GET(":id", middleware.ShareAvailable(), api.GetShare)
			share.DELETE(":id", middleware.ShareAvailable(), api.DeleteShare)
		}
	}
	return r
}
