package model

type FileFolder struct {
	BaseModel
	ParentID uint64 `json:"parent_id" gorm:"comment:'父文件夹ID'"`
	FileFolderName string	`json:"file_folder_name" gorm:"size:20;comment:'文件夹名称'"`
	FileStoreID uint64 `json:"file_store_id" gorm:"comment:'所属文件仓库'"`

}

func (f *FileFolder) TableName() string {
	return "file_folder"
}
