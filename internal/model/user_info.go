package model

type UserInfo struct {
	BaseModel
	Phone string `json:"phone" gorm:"size:11;comment:'电话号码'"`
	Email string `json:"email" gorm:"size:32;comment:'Email'"`
	Icon string `json:"icon" gorm:"size:255;comment:'头像地址'"`
	Introduction string `json:"introduction" gorm:"size:64;comment:'签名'"`
}

func (i *UserInfo) TableName() string {
	return "user_info"
}
