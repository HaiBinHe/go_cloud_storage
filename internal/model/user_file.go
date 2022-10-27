package model

type UserFile struct {
	BaseModel
	UserID       uint64 `json:"user_id" gorm:"comment:'用户ID'"`
	FileName     string `json:"file_name" gorm:"size:20;comment:'文件名'"`
	FileHash     string `json:"file_hash" gorm:"size:32;comment:'文件Hash'"`
	FileStoreID  uint64 `json:"file_store_id" gorm:"comment:'文件所属仓库'"`
	FileFolderID uint64 `json:"file_folder_id" gorm:"comment:'文件所属文件夹'"`
	FileSavePath string `json:"file_save_path" gorm:"size:32;comment:'文件保存路径'"`
	FileSize     int64  `json:"file_size" gorm:"comment:'文件大小(KB)'"`
	DownloadNum  int    `json:"download_num" gorm:"comment:'下载次数'"`
	FileType     string `json:"file_type" gorm:"size:10;comment:'文件类型'"`
	FileExt      string `json:"file_ext" gorm:"size:10;comment:'文件后缀'"`
}

func (uf *UserFile) TableName() string {
	return "user_file"
}

//创建
func (uf *UserFile) Create() error {
	return Db.Create(&uf).Error
}

//更新
func (uf *UserFile) Update() error {
	return Db.Where("id = ?", uf.ID).Updates(&uf).Error
}

//删除
func (uf *UserFile) Delete() error {
	return Db.Delete(&uf).Error
}
