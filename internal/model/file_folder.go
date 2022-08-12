package model

type FileFolder struct {
	BaseModel
	ParentID       uint64 `json:"parent_id" gorm:"column:parent_id;comment:'父文件夹ID'"`
	FileFolderName string `json:"file_folder_name" gorm:"column:folder_name;size:20;comment:'文件夹名称'"`
	FileStoreID    uint64 `json:"file_store_id" gorm:"comment:'所属文件仓库'"`
}

func (f *FileFolder) TableName() string {
	return "file_folder"
}

func GetFolderByID(folderID uint64) (FileFolder, error) {
	ff := FileFolder{}
	err := Db.Where("id = ?", folderID).First(&ff).Error
	if err != nil {
		return FileFolder{}, err
	}
	return ff, nil
}
