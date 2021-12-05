//@Title		routers.go
//@Description	路由
//@Author		zy
//@Update		2021.12.5

package routers

import (
	"csa_5/controller"
	"csa_5/middleware"
	"github.com/gin-gonic/gin"
)

//CollectRoute
//@title		CollectRoute()
//@description	各种路由集合
//@author		zy
//@param		r *gin.Engine
//@return		*gin.Engine
func CollectRoute(r *gin.Engine) *gin.Engine {

	// 主页进行token鉴权
	r.GET("/home", middleware.JWTAuthMiddleware(), controller.Home)

	// 注册
	r.POST("/register", controller.Register)

	// login要获取token
	r.POST("/login", controller.Login)

	// 发布文章
	r.POST("/postArticle", middleware.JWTAuthMiddleware(), controller.PostArticle)

	// 删除文章
	r.DELETE("/deleteArticle", middleware.JWTAuthMiddleware(), controller.Delete)

	// 点赞文章
	r.PUT("/like", middleware.JWTAuthMiddleware(), controller.Like)

	// 给其他用户留言
	r.POST("/message", middleware.JWTAuthMiddleware(), controller.MessageToOther)

	// 删除自己的留言
	r.DELETE("/deleteMessage", middleware.JWTAuthMiddleware(), controller.DeleteMsg)
	return r
}