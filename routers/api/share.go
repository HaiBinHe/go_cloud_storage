package api

import (
	"github.com/gin-gonic/gin"
	"go-cloud/internal/service/share"
	"go-cloud/pkg/app"
	error2 "go-cloud/pkg/error"
	"go-cloud/pkg/logger"
	"go-cloud/pkg/response"
)

//创建分享
func CreateShare(c *gin.Context) {
	var cs *share.CreateShare
	valid, err := app.BindAndValid(c, &cs)
	if !valid {
		logger.StdLog().Errorf(c, "app.BindAndValid err :%v", err)
		response.RespError(c, error2.InvalidParams)
		return
	}
	cs.Create(c)
}

//查看分享
//列出分享
func ListrShare(c *gin.Context) {
	var listShare *share.ShareListService
	var err error
	valid, err := app.BindAndValid(c, &listShare)
	if !valid {
		logger.StdLog().Errorf(c, "app.BindAndValid err :%v", err)
		response.RespError(c, error2.InvalidParams)
		return
	}
	listShare.ListShare(c)

}

//删除分享
func DeleteShare(c *gin.Context) {
	share.Delete(c)
}
