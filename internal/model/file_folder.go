package model

import "go-cloud/pkg/app"

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

func GetFolderByID(folderID uint64) (FileFolder, error) {
	ff := FileFolder{}
	err := Db.Where("id = ?", folderID).First(&ff).Error
	if err != nil {
		return FileFolder{}, err
	}
	return ff, nil
}
func (f *FileFolder) Create() error {
	return Db.Create(&f).Error
}

func ListDirectory(uid, pid, sid uint64, page, pageSize int, order string) ([]FileFolder, int64) {
	var (
		folders []FileFolder
		total   int64
	)
	Dbchain := Db.Model(&FileFolder{}).Where("user_id = ? AND parent_id = ? AND store_id = ?", uid, pid, sid).Find(&folders)
	Dbchain.Count(&total)

	//分页
	offset := app.GetPageOffset(page, pageSize)
	Dbchain.Limit(pageSize).Offset(offset).Order(order).Find(&folders)
	return folders, total

}
