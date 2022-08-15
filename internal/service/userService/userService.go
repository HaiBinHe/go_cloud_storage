package userService

import "github.com/gin-gonic/gin"

type LoginUser struct {
	ID       uint64 `form:"user_id"`
	userName string `form:"user_name"`
}

//获取用户详细信息
func GetUserInfo(c *gin.Context) {

}
