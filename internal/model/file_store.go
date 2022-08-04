package model


type FileStore struct {
	BaseModel
	FolderName string `json:"folder_name" gorm:"type:varchar(20);comment:'文件仓库名称'"`
	MaxSize int `json:"max_size" gorm:"comment:'文件仓库最大容量'"`
	CurrentSize int `json:"current_size" gorm:"comment:'文件仓库当前容量'"`
	UserID uint64	`json:"user_ID" gorm:"column:user_id;comment:'所属用户'"`
}

func (f *FileStore) TableName() string {
	return "file_store"
}
