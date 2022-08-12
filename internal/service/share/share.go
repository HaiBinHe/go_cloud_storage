package share

import (
	"github.com/gin-gonic/gin"
	"go-cloud/conf"
	"go-cloud/internal/model"
	"go-cloud/pkg/app"
	error2 "go-cloud/pkg/error"
	"go-cloud/pkg/logger"
	"go-cloud/pkg/response"
	"go-cloud/tools"
	"net/url"
	"strconv"
	"time"
)

type CreateShare struct {
	SourceID  uint64 `json:"source_id"`
	IsDir     bool   `json:"is_dir" binding:"oneof=0 1"`
	ShareCode string `json:"share_code" binding:"len=6"`
	ExpireAt  int64  `json:"expire_at"`
}

func (s *CreateShare) Create(c *gin.Context) {
	var err error
	userCtx, _ := c.Get("user")
	user := userCtx.(*model.User)
	cs := CreateShare{}
	valid, err := app.BindAndValid(c, &cs)
	if !valid {
		logger.StdLog().Errorf(c, "app.BindAndValid err :%v", err)
		response.RespError(c, error2.InvalidParams)
		return
	}
	var sourceName string
	//判断是否是 目录
	if cs.IsDir {
		folder, _ := model.GetFolderByID(s.SourceID)
		sourceName = folder.FileFolderName
	} else {
		file, _ := model.GetFileByFileIDAndUserID(s.SourceID, user.ID)
		sourceName = file.FileName
	}

	share := model.FileShare{
		SourceName: sourceName,
		IsDir:      s.IsDir,
		UserID:     user.ID,
		ShareCode:  s.ShareCode,
		ShareTime:  time.Now().Add(time.Duration(s.ExpireAt) * time.Second),
	}
	err = share.Create()
	if err != nil {
		response.RespError(c, "创建分享链接失败")
		return
	}
	//TODO 生成分享链接
	code := strconv.FormatUint(s.SourceID, 10)
	sid := tools.EncodeMD5(sourceName + code)
	sharePath, _ := url.Parse("/s/" + sid)
	response.RespData(c, map[string]interface{}{
		"data":     share,
		"shareUrl": conf.AppSetting.ShareUrl + sharePath.String(),
	})

}
