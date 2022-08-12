package share

import (
	"github.com/gin-gonic/gin"
	"go-cloud/conf"
	"go-cloud/internal/model"
	"go-cloud/pkg/response"
	"go-cloud/tools"
	"net/url"
	"strconv"
	"time"
)

type ShareCreateService struct {
	SourceID  uint64 `form:"source_id"`
	IsDir     bool   `form:"is_dir" binding:"oneof=0 1"`
	ShareCode string `form:"share_code" binding:"len=6"`
	ExpireAt  int64  `form:"expire_at"`
}
type ShareListService struct {
	Page     int    `form:"page" binding:"required,min=1"`
	PageSize int    `form:"pageSize" binding:"required, min=10"`
	order    string `form:"order" binding:"required,eq=DESC|eq=ASC"`
}
type ShareGetService struct {
	ShareCode string `form:"share_code" binding:"required,len=6"`
}

//创建分享记录
func (s *ShareCreateService) Create(c *gin.Context) {
	var err error
	userCtx, _ := c.Get("user")
	user := userCtx.(*model.User)

	var sourceName string
	//判断是否是 目录
	if s.IsDir {
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

//删除分享记录
func Delete(c *gin.Context) {
	shareID, _ := strconv.ParseUint(c.Param("share_id"), 10, 64)
	share, err := model.GetShareByID(shareID)
	if err != nil {
		response.RespError(c, "查询不到数据")
		return
	}
	//获取当前用户
	userCtx, _ := c.Get("user")
	user := userCtx.(*model.User)
	if share.UserID != user.ID {
		response.RespError(c, "该分享记录对应的用户不正确")
		return
	}
	err = share.Delete()
	if err != nil {
		response.RespError(c, "删除分享记录失败")
		return
	}
	response.RespSuccess(c, "删除该分享记录成功")
}

//用户的分享列表
func (s *ShareListService) ListShare(c *gin.Context) {
	//获取当前用户
	userCtx, _ := c.Get("user")
	user := userCtx.(*model.User)
	shares, total := model.ListShares(user.ID, s.Page, s.PageSize, s.order)
	//列出分享对应的文件
	for i := 0; i < len(shares); i++ {
		shares[i].Source()
	}
	response.RespList(c, shares, total)
}

//获取分享
func (s *ShareGetService) GetShare(c *gin.Context) {
	//在中间件里设置 c.set(share)
	shareCtx, _ := c.Get("share")
	share := shareCtx.(*model.FileShare)
	if share.ShareCode != "" {
		//提取码正确
		if share.ShareCode == s.ShareCode {
			response.RespData(c, share)
			return
		} else {
			response.RespError(c, "提取码错误")
			return
		}
	}
}
