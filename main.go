//@Title		main.go
//@Description	运行界面
//@Author		zy
//@Update		2021.12.5

package main

import (
	"csa_5/common"
	"csa_5/routers"
	"github.com/gin-gonic/gin"
)

//功能
//1.注册登录
//2.发表文章
//3.删除文章
//4.点赞文章
//5.给他人留言
//6.主页
func main() {
	DB, _ := common.InitDB("user:123456@tcp(127.0.0.1:3306)/sql_test")  //连接数据库
	defer DB.Close()														 //defer关闭数据库
	r := gin.Default()														//默认路由
	r = routers.CollectRoute(r)

	panic(r.Run(":9090"))											//启动服务
}
