package app

import (
	"github.com/gin-gonic/gin"
	"go-cloud/conf"
)

//获取页码
func GetPage(c *gin.Context) int {
	page := StrTo(c.Query("page")).MustInt()
	if page <= 0 {
		return 1
	}
	return page
}

//每页显示的数量
func GetPageSize(c *gin.Context) int {
	pageSize := StrTo(c.Query("page_size")).MustInt()
	if pageSize <= 0 {
		return conf.AppSetting.DefaultPageSize
	}
	if pageSize > conf.AppSetting.MaxPageSize {
		return conf.AppSetting.MaxPageSize
	}
	return pageSize
}

//偏移量
func GetPageOffset(page, pageSize int) int {
	offset := 0
	if page > 0 {
		offset = (page - 1) * pageSize
	}
	return offset
}
