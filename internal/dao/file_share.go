package dao

import (
	"go-cloud/internal/model"
	"go-cloud/pkg/app"
	"time"
)

//根据ShareID获取记录
func GetShareByID(sharID uint64) (*model.FileShare, error) {
	fs :=  &model.FileShare{}
	err := model.Db.Where("id = ?", sharID).First(&fs).Error
	if err != nil {
		return nil, err
	}
	return fs, nil
}

//列出UID下的分享列表
func ListShares(uid uint64, page, pageSize int, order string) ([]*model.FileShare, int64) {
	var (
		shares []*model.FileShare
		total  int64
	)
	Dbchain := model.Db.Model(&model.FileShare{}).Where("user_id = ?", uid)
	//计算总数
	Dbchain.Count(&total)

	//分页
	offset := app.GetPageOffset(page, pageSize)
	Dbchain.Limit(pageSize).Offset(offset).Order(order).Find(&shares)
	return shares, total
}

//此分享是否过期
func  IsAvailable(s *model.FileShare) bool {

	if time.Now().After(s.ShareTime) {
		return false
	}
	//检查源文件是否存在
	var sourceID uint64
	if s.IsDir {
		folder := sourceFolder(s)
		sourceID = folder.ID
	} else {
		file := sourceFile(s)
		sourceID = file.ID
	}
	if sourceID == 0 {
		// TODO 是否要在这里删除这个无效分享？
		_ = s.Delete()

		return false
	}
	return true
}

//获取源对象
func Source(s *model.FileShare) interface{} {
	if s.IsDir {
		return sourceFolder(s)
	}
	return sourceFile(s)
}

//获取目录
func sourceFolder(s *model.FileShare) *model.FileFolder {
	//资源ID获取目录ID
	if s.Folder.ID == 0 {
		ff, _ := GetFolderByID(s.Folder.ID)
		s.Folder = ff
	}
	return s.Folder
}

//获取文件
func sourceFile(s *model.FileShare) *model.UserFile {
	if s.File.ID == 0 {
		f, _ := GetFileByFileIDAndUserID(s.File.ID, s.UserID)
		s.File = f
	}
	return s.File
}

//获取分享的创建者
func Creator(s *model.FileShare)  *model.User {
	if s.User.ID == 0 {
		u := GetUserByID(s.User.ID)
		s.User = u
	}
	return s.User
}