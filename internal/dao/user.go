package dao

import (
	"go-cloud/internal/cache"
	"go-cloud/internal/model"
)

func QueryUserByWhere(where ...interface{}) (*model.User, error) {
	user := &model.User{}
	return user, cache.Db.First(&user, where...).Error
}

// 根据OpenID获取用户信息
func GetUserByOpenID(OpenID string)  *model.User {
	user := &model.User{}
	cache.Db.Where("open_id = ?", OpenID).First(&user)
	return user
}

//根据ID获取用户信息
func GetUserByID(userID uint64) *model.User {
	user := &model.User{}
	cache.Db.Set("gorm:auto_preload", true).First(&user, userID)
	return user
}

//根据OpenID查看用户是否存在
func UserIsExist(OpenID string) bool {
	return cache.Db.Where("open_id=?", OpenID).RowsAffected > 0
}