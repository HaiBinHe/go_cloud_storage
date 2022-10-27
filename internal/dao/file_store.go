package dao

import (
	"go-cloud/internal/model"
)

//根据用户ID获取仓库列表
func GetUserFileStore(UID uint64) (fs []*model.FileStore) {
	model.Db.Where("user_id = ?", UID).Find(&fs)
	return
}

//判断用户存储仓库容量是否足够
func CapacityIsEnough(fileSize int64, fileStoreID uint64) bool {
	var fs model.FileStore
	model.Db.First(&fs, fileStoreID)
	return fs.CurrentSize - fileSize < 0
}
