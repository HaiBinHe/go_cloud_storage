package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-cloud/pkg/auth"
	"net/http"
	"time"
)

const JwtName = "Authorization"

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := getToken(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status": 401,
				"msg":    "token不存在",
			})
			return
		}
		claims, err := auth.ParseToken(token)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status": 401,
				"msg":    "无效token",
			})
			return
		}
		if time.Now().Unix() > claims.ExpiresAt {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status": 401,
				"msg":    "token已过期",
			})
			return
		}
		c.Next()
	}
}

// 各种方法获取 token
// 为了防范 CSRF 攻击,不获取 query 和 from 里的 token
func getToken(c *gin.Context) (string, error) {
	if token := c.GetHeader(JwtName); token != "" {
		return token, nil
	}

	if token, _ := c.Cookie(JwtName); token != "" {
		return token, nil
	}
	return "", errors.New("没有找到" + JwtName)
}
