package model

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"go-cloud/tools"
	"strings"
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

func QueryUserByWhere(where ...interface{}) (user User, err error) {
	err = Db.First(&user, where...).Error
	return
}

// 根据OpenID获取用户信息
func GetUserByOpenID(OpenID string) (user User) {
	Db.Where("open_id = ?", OpenID).First(&user)
	return
}

//根据OpenID查看用户是否存在
func UserIsExist(OpenID string) bool {
	affected := Db.Where("open_id=?", OpenID).RowsAffected
	if affected > 0 {
		return true
	} else {
		return false
	}
}

//创建新用户
func (u *User) Create() error {
	password, err := setPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = password
	return Db.Create(&u).Error
}

// 根据 `struct` 更新属性，只会更新非零值的字段
func (u *User) Update() error {
	return Db.Where("id = ?", u.ID).Updates(&u).Error
}

//删除用户
func (u *User) Delete() error {
	return Db.Delete(&u).Error
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

//检查密码是否与数据库记录的相同
func (u *User) CheckPassword(password string) (bool, error) {
	passwordStore := strings.Split(u.Password, ":")
	if len(passwordStore) != 2 {
		return false, errors.New("UnKnown Password Type")
	}

	//计算 salt和密码组合的SHA1摘要
	hash := sha1.New()
	_, err := hash.Write([]byte(passwordStore[0] + password))
	bs := hex.EncodeToString(hash.Sum(nil))
	if err != nil {
		return false, err
	}
	return bs == passwordStore[1], nil
}
