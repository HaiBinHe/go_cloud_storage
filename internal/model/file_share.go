package model

import "time"

type FileShare struct {
	BaseModel
	UserID uint64 	`json:"user_id" gorm:"comment:'用户'"`
	FileID uint64	`json:"file_id" gorm:"comment:'文件'"`
	ShareCode string	`json:"share_code" gorm:"type:varchar(255);comment:'分享码'"`
	ShareStatus uint `json:"share_status" gorm:"size:2;comment:'分享状态(0:未分享,1:已过期)'" `
	ShareTime time.Duration `json:"share_time" gorm:"type:varchar(255);comment:'有效时长'"`
}

func (s *FileShare) TableName() string {
	return "file_share"
}
