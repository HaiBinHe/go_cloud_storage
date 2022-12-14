package model

type FileFolder struct {
	BaseModel
	UserID         uint64 `json:"user_id" gorm:"column:user_id;comment:'用户ID'"`
	ParentID       uint64 `json:"parent_id" gorm:"column:parent_id;comment:'父文件夹ID'"`
	FileFolderName string `json:"file_folder_name" gorm:"column:folder_name;size:20;comment:'文件夹名称'"`
	FileStoreID    uint64 `json:"file_store_id" gorm:"comment:'所属文件仓库'"`
}

func (f *FileFolder) TableName() string {
	return "file_folder"
}

func (f *FileFolder) Create() error {
	return Db.Create(&f).Error
}
