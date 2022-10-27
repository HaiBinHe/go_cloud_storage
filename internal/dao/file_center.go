package dao

import (
	"go-cloud/internal/model"
)

// 文件是否存在
func FileIsExist(fileHash string) bool {
	fc := &model.FileCenter{}
	n := model.Db.Where("file_hash = ?", fileHash).First(&fc).RowsAffected
	return n == 1
}
