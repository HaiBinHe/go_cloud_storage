package middleware

import (
	"github.com/gin-gonic/gin"
	"go-cloud/internal/model"
	"go-cloud/pkg/response"
	"strconv"
)

// ShareOwner 检查当前登录用户是否为分享所有者
func ShareOwner() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user *model.User
		if userCtx, ok := c.Get("user"); ok {
			user = userCtx.(*model.User)
		} else {
			response.RespError(c, "请先登录")
			c.Abort()
			return
		}

		if share, ok := c.Get("share"); ok {
			if share.(*model.FileShare).Creator().ID != user.ID {
				response.RespError(c, "分享不存在")
				c.Abort()
				return
			}
		}
		c.Next()
	}
}

// ShareAvailable 检查分享是否可用
func ShareAvailable() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user *model.User
		if userCtx, ok := c.Get("user"); ok {
			user = userCtx.(*model.User)
		}
		shareID, _ := strconv.ParseUint(c.Param("share_id"), 10, 64)
		share, err := model.GetShareByID(shareID)
		if err != nil {
			response.RespError(c, "分享不存在或者已失效")
			c.Abort()
			return

		}
		c.Set("user", user)
		c.Set("share", share)
		c.Next()
	}
}
