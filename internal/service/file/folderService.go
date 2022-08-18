package file

import (
	"github.com/gin-gonic/gin"
	"go-cloud/internal/model"
	"go-cloud/pkg/response"
)

type CreateFolder struct {
	UserID   uint64
	ParentID uint64 `form:"parent_id"`
	StoreID  uint64 `form:"store_id"`
	Name     string `form:"name"`
}
type ListFolder struct {
	ParentID uint64 `form:"parent_id"`
	StoreID  uint64 `form:"store_id"`
	Page     int    `form:"page" binding:"required,min=1"`
	PageSize int    `form:"pageSize" binding:"required, min=10"`
	order    string `form:"order" binding:"required,eq=DESC|eq=ASC"`
}

func (f *CreateFolder) Create(c *gin.Context) {
	folder := &model.FileFolder{
		UserID:         f.UserID,
		ParentID:       f.ParentID,
		FileFolderName: f.Name,
		FileStoreID:    f.StoreID,
	}
	err := folder.Create()
	if err != nil {
		response.RespError(c, "创建目录失败")
		return
	}

}
func (f *ListFolder) ListFolders(c *gin.Context) {
	userCtx, _ := c.Get("user")
	user := userCtx.(*model.User)
	folders, totals := model.ListDirectory(user.ID, f.ParentID, f.StoreID, f.Page, f.PageSize, f.order)
	//返回所有目录
	response.RespList(c, folders, totals)

}
