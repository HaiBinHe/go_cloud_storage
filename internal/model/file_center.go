package model

//FileCenter 文件中心存储，用于实现秒传
type FileCenter struct {
	BaseModel
	FileName     string `json:"file_name" gorm:"size:20;comment:'文件名'"`
	FileHash     string `json:"file_hash" gorm:"size:32;comment:'文件Hash'"`
	FileSavePath string `json:"file_save_path" gorm:"size:32;comment:'文件保存路径'"`
	FileSize     int64    `json:"file_size" gorm:"comment:'文件大小(KB)'"`
	FileType     string `json:"file_type" gorm:"size:10;comment:'文件类型'"`
	FileExt      string `json:"file_ext" gorm:"size:10;comment:'文件后缀'"`
}

func (f *FileCenter) TableName() string {
	return "file_center"
}

func (f *FileCenter) Create() error {
	return Db.Create(&f).Error
}
