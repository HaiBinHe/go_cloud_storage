package api

import (
	"github.com/gin-gonic/gin"
	"go-cloud/internal/model"
	"go-cloud/internal/service/file"
	"go-cloud/pkg/app"
	error2 "go-cloud/pkg/error"
	"go-cloud/pkg/logger"
	"go-cloud/pkg/response"
)

func CreateDir(c *gin.Context) {
	var err error
	userCtx, _ := c.Get("user")
	user := userCtx.(*model.User)
	var cf *file.CreateFolder
	valid, err := app.BindAndValid(c, &cf)
	if !valid {
		logger.StdLog().Errorf(c, "app.BindAndValid err :%v", err)
		response.RespError(c, error2.InvalidParams)
		return
	}
	cf.UserID = user.ID
	cf.Create(c)
	if err != nil {
		response.RespError(c, error2.ServerError)
	}
	response.RespData(c, "目录创建成功")
}

func ListDirectories(c *gin.Context) {
	var listFolders *file.ListFolder
	var err error
	valid, err := app.BindAndValid(c, &listFolders)
	if !valid {
		logger.StdLog().Errorf(c, "app.BindAndValid err :%v", err)
		response.RespError(c, error2.InvalidParams)
		return
	}
	listFolders.ListFolders(c)
}
