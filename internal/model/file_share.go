package model

import (
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
	User   *User       `gorm:"-;PRELOAD:false;association_autoupdate:false"`
	Folder *FileFolder `gorm:"-;PRELOAD:false;association_autoupdate:false"`
	File   *UserFile   `gorm:"-;PRELOAD:false;association_autoupdate:false"`
}

func (s *FileShare) TableName() string {
	return "file_share"
}

//创建文件分享记录
func (s *FileShare) Create() error {
	return Db.Create(&s).Error
}

//删除分享
func (s *FileShare) Delete() error {
	return Db.Delete(&s).Error
}
