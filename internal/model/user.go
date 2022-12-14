package model

import (
	"crypto/sha1"
	"encoding/hex"
	"go-cloud/tools"
)

type User struct {
	BaseModel
	UserName     string `json:"user_name" gorm:"type:varchar(20);comment:'用户名'"`
	Password     string `json:"-" gorm:"size:255;comment:'密码'"`
	Phone        string `json:"phone" gorm:"column:phone;size:11;comment:'电话号码'"`
	Email        string `json:"email" gorm:"size:32;comment:'Email'"`
	Icon         string `json:"icon" gorm:"size:255;comment:'头像地址'"`
	Introduction string `json:"introduction" gorm:"size:64;comment:'签名'"`
	OpenID       string `json:"open_id" gorm:"column:open_id;type:varchar(255);comment:'扫码登陆'"`
}

//创建新用户
func (u *User) Create() error {
	password, err := setPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = password
	err = Db.Create(&u).Error
	if err != nil {
		return err
	}
	//创建默认的文件仓库
	store := FileStore{
		UserID:    u.ID,
		StoreName: "default",
		//默认1G
		MaxSize:     1048576,
		CurrentSize: 0,
	}
	err = store.Create()
	if err != nil {
		return err
	}

	return nil
}

// 根据 `struct` 更新属性，只会更新非零值的字段
func (u *User) Update() error {
	return Db.Model(&User{}).Where("id = ?", u.ID).Updates(&u).Error
}

//加密密码
func setPassword(password string) (string, error) {
	//生成8位salt
	salt := tools.GetRandomBoth(8)
	//计算 salt和密码组合的SHA1摘要
	hash := sha1.New()
	_, err := hash.Write([]byte(salt + password))
	bs := hex.EncodeToString(hash.Sum(nil))
	if err != nil {
		return "", err
	}
	//存储 Salt 值和摘要， ":"分割
	pwd := salt + ":" + bs
	return pwd, nil
}
