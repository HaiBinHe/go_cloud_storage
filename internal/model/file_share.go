package model

import (
	"go-cloud/pkg/app"
	"time"
)

type FileShare struct {
	BaseModel
	UserID      uint64    `json:"user_id" gorm:"comment:'用户'"`
	IsDir       bool      `json:"is_dir" gorm:"type:tinyint(1);comment:'是否为目录:0-文件，1-目录'"`
	SourceID    uint64    `json:"source_id" gorm:"comment:'资源ID'"`
	SourceName  string    `json:"source_name" gorm:"comment:'资源名称'"`
	ShareCode   string    `json:"share_code" gorm:"type:varchar(255);comment:'分享码'"`
	ShareStatus uint      `json:"share_status" gorm:"size:2;comment:'分享状态(0:未分享,1:已过期)'" `
	ShareTime   time.Time `json:"share_time" gorm:"comment:'有效时长'"`

	//非数据库字段
	User   User       `gorm:"-;PRELOAD:false;association_autoupdate:false"`
	Folder FileFolder `gorm:"-;PRELOAD:false;association_autoupdate:false"`
	File   UserFile   `gorm:"-;PRELOAD:false;association_autoupdate:false"`
}

func (s *FileShare) TableName() string {
	return "file_share"
}

//创建文件分享记录
func (s *FileShare) Create() error {
	return Db.Create(&s).Error
}

//此分享是否过期
func (s *FileShare) IsAvailable() bool {

	if time.Now().After(s.ShareTime) {
		return false
	}
	//检查源文件是否存在
	var sourceID uint64
	if s.IsDir {
		folder := s.SourceFolder()
		sourceID = folder.ID
	} else {
		file := s.SourceFile()
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
func (s *FileShare) Source() interface{} {
	if s.IsDir {
		return s.SourceFolder()
	}
	return s.SourceFile()
}

//获取目录
func (s *FileShare) SourceFolder() *FileFolder {
	//资源ID获取目录ID
	if s.Folder.ID == 0 {
		ff, _ := GetFolderByID(s.Folder.ID)
		s.Folder = ff
	}
	return &s.Folder
}

//获取文件
func (s *FileShare) SourceFile() *UserFile {
	if s.File.ID == 0 {
		f, _ := GetFileByFileIDAndUserID(s.File.ID, s.UserID)
		s.File = f
	}
	return &s.File
}

//获取分享的创建者
func (s *FileShare) Creator() *User {
	if s.User.ID == 0 {
		u := GetUserByID(s.User.ID)
		s.User = u
	}
	return &s.User
}

//删除分享
func (s *FileShare) Delete() error {
	return Db.Delete(&s).Error
}

//列出UID下的分享列表
func ListShares(uid uint64, page, pageSize int, order string) ([]FileShare, int64) {
	var (
		shares []FileShare
		total  int64
	)
	Dbchain := Db.Model(&FileShare{}).Where("user_id = ?", uid)
	//计算总数
	Dbchain.Count(&total)

	//分页
	offset := app.GetPageOffset(page, pageSize)
	Dbchain.Limit(pageSize).Offset(offset).Order(order).Find(&shares)
	return shares, total
}

//根据ShareID获取记录
func GetShareByID(sharID uint64) (FileShare, error) {
	fs := FileShare{}
	err := Db.Where("id = ?", sharID).First(&fs).Error
	if err != nil {
		return FileShare{}, err
	}
	return fs, nil
}
