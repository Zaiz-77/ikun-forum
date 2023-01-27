package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"zaizwk/ginessential/common"
	"zaizwk/ginessential/model"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取认证头
		tokenString := c.GetHeader("Authorization")

		// 验证 token 格式
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "Token格式错误！权限不足！",
			})
			c.Abort()
			return
		}

		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "Token无效！权限不足！",
			})
			c.Abort()
			return
		}

		// 验证通过
		userID := claims.UserID
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userID)

		// 用户是否存在
		if user.ID == 0 {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "数据错误！权限不足！",
			})
			c.Abort()
			return
		}

		// 存在，写入上下文
		c.Set("user", user)
		c.Next()
	}
}
