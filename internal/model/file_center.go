package model

import (
	"context"
	"go-cloud/pkg/logger"
)

//FileCenter 文件中心存储，用于实现秒传
type FileCenter struct {
	BaseModel
	FileName     string `json:"file_name" gorm:"size:20;comment:'文件名'"`
	FileHash     string `json:"file_hash" gorm:"size:32;comment:'文件Hash'"`
	FileSavePath string `json:"file_save_path" gorm:"size:32;comment:'文件保存路径'"`
	FileSize     int    `json:"file_size" gorm:"comment:'文件大小(KB)'"`
	FileType     string `json:"file_type" gorm:"size:10;comment:'文件类型'"`
	FileExt      string `json:"file_ext" gorm:"size:10;comment:'文件后缀'"`
}

func (f *FileCenter) TableName() string {
	return "file_center"
}
func NewFileCenter() FileCenter {
	return FileCenter{}
}

//FileIsExist 根据给出的文件Hash判断是否存在于中心存储表中
func (f *FileCenter) FileIsExist(fileHash string) bool {
	var fc FileCenter
	err := Db.Where("file_hash = ?", fileHash).First(&fc).Error
	if err != nil {
		logger.StdLog().Error(context.Background(), "The file does not exist in the fileCenterTable")
		return false
	} else {
		return true
	}
}
