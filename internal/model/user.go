package model

type User struct {
	BaseModel
	UserName string `json:"user_name" gorm:"type:varchar(20);comment:'用户名'"`
	Password string `json:"-" gorm:"size:32;comment:'密码'"`
	OpenID string	`json:"open_id" gorm:"column:open_id;type:varchar(255);comment:'扫码登陆'"`
	UserInfoID uint64 `json:"user_info_id" gorm:"column:user_info_id'"`
}

