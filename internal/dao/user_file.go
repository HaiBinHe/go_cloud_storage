package dao

import (
	"context"
	"go-cloud/internal/model"
	"go-cloud/pkg/logger"
	"go-cloud/tools"
)

//根据文件ID查询
func GetFileByFileIDAndUserID(fileID, UserID uint64) (*model.UserFile, error) {
	uf := &model.UserFile{}
	err := model.Db.Where("id = ? and user_id = ?", fileID, UserID).Find(uf).Error
	if err != nil {
		return nil, err
	}
	return uf, nil
}

//根据文件夹id查询包含的文件
func GetFilesByFolderID(fileFolderID uint64, page, offset int) ([]*model.UserFile, error) {
	var userFiles []*model.UserFile
	err := model.Db.Limit(page).Offset(offset).Where("file_folder_id=?", fileFolderID).Find(&userFiles).Error
	if err != nil {
		logger.StdLog().Error(context.Background(), "sql failed:", err)
		return nil, err
	}
	return userFiles, nil
}

//获取文件数量
func GetUserFileCount(fileStoreID uint64) (count *int64) {
	var files []model.UserFile
	model.Db.Where("file_store_id = ?", fileStoreID).Find(&files).Count(count)
	return
}

//根据文件类型获取文件
func GetFilesByType(t tools.FileType, fileStoreID uint64) ([]model.UserFile, error) {
	var files []model.UserFile
	err := model.Db.Where("file_store_id = ? AND file_type = ?", fileStoreID, t).Find(&files).Error
	if err != nil {
		logger.StdLog().Error(context.Background(), "GetFilesByType work failed")
		return nil, err
	}
	return files, nil
}
