//@Title		AuthMiddleware.go
//@Description	jwt鉴权中间件实现
//@Author		zy
//@Update		2021.12.5

package middleware

import (
	"csa_5/common"
	"csa_5/userinformation"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)


//JWTAuthMiddleware
//@title		JWTAuthMiddleware()
//@description	JWT鉴权中间件
//@author		zy
//@param
//@return		gin.HandlerFunc
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")		//获取Header的Authorization
		if authHeader == "" {
			c.JSON(http.StatusOK, gin.H{
				"code": 2003,
				"msg": "请求头中auth为空",
			})
			c.Abort()
			return
		}

		// 按空格分割
		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusOK, gin.H{
				"code": 2004,
				"msg": "请求头中auth格式有误",
			})
			c.Abort()
			return
		}
		// parts[1]是获取到的tokenString, 使用之前解析JWT函数来解析
		mc, err := common.ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": 2005,
				"msg": "无效的Token",
			})
			c.Abort()
			return
		}


		// 在数据库中验证mc中的Username
		// 若不存在则返回
		var u userinformation.UserInfo
		u.Username = mc.Username
		if !common.QueryUserInfo(u) {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"status": 401,
				"message": "权限不足",
			})
			c.Abort()
			return
		}

		//用户存在 将user信息写入请求的上下文c中
		c.Set("username", mc.Username)
		c.Next()	// 后续的处理函数可以用过c.Get("username")来获取当前请求的用户信息
	}
}