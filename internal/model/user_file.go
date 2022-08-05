package model

import (
	"context"
	"go-cloud/pkg/logger"
	"go-cloud/tools"
)

type UserFile struct {
	BaseModel
	FileName     string `json:"file_name" gorm:"size:20;comment:'文件名'"`
	FileHash     string `json:"file_hash" gorm:"size:32;comment:'文件Hash'"`
	FileStoreID  uint64 `json:"file_store_id" gorm:"comment:'文件所属仓库'"`
	FileFolderID uint64 `json:"file_folder_id" gorm:"comment:'文件所属文件夹'"`
	FileSavePath string `json:"file_save_path" gorm:"size:32;comment:'文件保存路径'"`
	FileSize     int64  `json:"file_size" gorm:"comment:'文件大小(KB)'"`
	DownloadNum  int    `json:"download_num" gorm:"comment:'下载次数'"`
	FileType     string `json:"file_type" gorm:"size:10;comment:'文件类型'"`
	FileExt      string `json:"file_ext" gorm:"size:10;comment:'文件后缀'"`
}

func (uf *UserFile) TableName() string {
	return "user_file"
}

//根据文件夹id查询包含的文件
func GetFilesByFolderID(fileFolderID uint64, page, offset int) ([]*UserFile, error) {
	var userFiles []*UserFile
	err := Db.Limit(page).Offset(offset).Where("file_folder_id=?", fileFolderID).Find(&userFiles).Error
	if err != nil {
		logger.StdLog().Error(context.Background(), "sql failed:", err)
		return nil, err
	}
	return userFiles, nil
}

//获取文件数量
func GetUserFileCount(fileStoreID uint64) (count *int64) {
	var files []UserFile
	Db.Where("file_store_id = ?", fileStoreID).Find(&files).Count(count)
	return
}

//根据文件类型获取文件
func GetFilesByType(t tools.FileType, fileStoreID uint64) ([]UserFile, error) {
	var files []UserFile
	err := Db.Where("file_store_id = ? AND file_type = ?", fileStoreID, t).Find(&files).Error
	if err != nil {
		logger.StdLog().Error(context.Background(), "GetFilesByType work failed")
		return nil, err
	}
	return files, nil
}

//创建
func (uf *UserFile) Create() error {
	return Db.Create(&uf).Error
}

//更新
func (uf *UserFile) Update() error {
	return Db.Where("id = ?", uf.ID).Updates(&uf).Error
}

//删除
func (uf *UserFile) Delete() error {
	return Db.Delete(&uf).Error
}
