package dao

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"go-cloud/internal/model"
	"strings"
)


//检查密码是否与数据库记录的相同
func CheckPassword(password string) (bool, error) {
	passwordStore := strings.Split(password, ":")
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

func QueryUserByWhere(where ...interface{}) (*model.User, error) {
	user := &model.User{}
	return user, model.Db.First(&user, where...).Error
}

// 根据OpenID获取用户信息
func GetUserByOpenID(OpenID string)  *model.User {
	user := &model.User{}
	model.Db.Where("open_id = ?", OpenID).First(&user)
	return user
}

//根据ID获取用户信息
func GetUserByID(userID uint64) *model.User {
	user := &model.User{}
	model.Db.Set("gorm:auto_preload", true).First(&user, userID)
	return user
}

//根据OpenID查看用户是否存在
func UserIsExist(OpenID string) bool {
	return model.Db.Where("open_id=?", OpenID).RowsAffected > 0
}