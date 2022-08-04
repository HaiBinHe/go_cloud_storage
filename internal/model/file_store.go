package model

type FileStore struct {
	BaseModel
	StoreName   string `json:"store_name" gorm:"type:varchar(20);comment:'文件仓库名称'"`
	MaxSize     int64  `json:"max_size" gorm:"type:bigint(20);default:104857600;comment:'文件仓库最大容量(KB)'"`
	CurrentSize int64  `json:"current_size" gorm:"type:bigint(20);default:104857600;comment:'文件仓库当前容量(KB)'"`
	UserID      uint64 `json:"user_ID" gorm:"comment:'所属用户'"`
}

func (f *FileStore) TableName() string {
	return "file_store"
}

//根据用户ID获取仓库列表
func GetUserFileStore(UID uint64) (fs []FileStore) {
	Db.Where("user_id = ?", UID).Find(&fs)
	return
}

//判断用户存储仓库容量是否足够
func CapacityIsEnough(fileSize int64, fileStoreID uint64) bool {
	var fs FileStore
	Db.First(&fs, fileStoreID)
	if fs.CurrentSize-fileSize < 0 {
		return false
	} else {
		return true
	}
}
