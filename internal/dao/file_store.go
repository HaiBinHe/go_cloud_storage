package dao

import (
	"go-cloud/internal/model"
)

type FileStore struct {
	StoreName   string `json:"store_name" gorm:"type:varchar(20);comment:'文件仓库名称'"`
	MaxSize     int64  `json:"max_size" gorm:"type:bigint(20);default:104857600;comment:'文件仓库最大容量(KB)'"`
	UserID      uint64 `json:"user_id" gorm:"comment:'所属用户'"`
}
// 根据用户 id 删除 fileStore
func DeleteFileStoreByUserID(userId uint64) error {
	return model.Db.Where("user_id = ?", userId).Delete(&model.FileStore{}).Error
}

// 根据用户 id 修改 fileStore
func UpdateFileStoreByUserID(store FileStore) error{
	fs := &model.FileStore{
		UserID: store.UserID,
		MaxSize: store.MaxSize,
		StoreName: store.StoreName,
	}
	return fs.Update()
}

// 根据用户ID获取仓库列表
func GetUserFileStore(userId  uint64) (fs []*model.FileStore) {
	model.Db.Where("user_id = ?", userId ).Find(&fs)
	return
}

// 判断用户存储仓库容量是否足够
func CapacityIsEnough(fileSize int64, fileStoreID uint64) bool {
	var fs model.FileStore
	model.Db.First(&fs, fileStoreID)
	return fs.CurrentSize - fileSize < 0
}
