//@Title		UserController.go
//@Description	客户端的各个请求
//@Author		zy
//@Update		2021.12.5

package controller

import (
	"csa_5/common"
	"csa_5/userinformation"
	"github.com/gin-gonic/gin"
	"net/http"
)

//Register
//@title		Register()
//@description	注册请求
//@author		zy
//@param		c *gin.Context
//@return
func Register(c *gin.Context) {
	var u userinformation.UserInfo
	err := c.ShouldBind(&u)				//参数绑定
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 2001,
			"message": "无效的参数",
		})
		return
	}

	//判断用户名是否已经存在
	if common.QueryUserInfo(u) {
		c.JSON(http.StatusForbidden, gin.H{
			"status": http.StatusForbidden,
			"message": "已存在该用户名",
		})
		return
	} else {
		if common.InsertUserInfo(u){
			c.JSON(http.StatusOK, gin.H{
				"status": http.StatusOK,
				"message": "注册成功",
			})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status": http.StatusInternalServerError,
				"message": "some errors in sql",
			})
		}
	}
}


//Login
//@title		Login()
//@description	登录请求
//@author		zy
//@param		c *gin.Context
//@return
func Login(c *gin.Context) {
	var u userinformation.UserInfo
	err := c.ShouldBind(&u)					//参数绑定
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 2001,
			"message": "无效的参数",
		})
		return
	}
	// 判断用户名密码是否正确
	if !common.QueryUserInfoExist(u) {
		c.JSON(http.StatusForbidden, gin.H{
			"status": http.StatusForbidden,
			"message": "用户名或密码有误!",
		})
		return
	}

	// 生成username对应的tokenString
	tokenString, err := common.GenToken(u.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"message": "系统异常",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 2000,
		"message": "login success",
		"data": gin.H{"token": tokenString},
	})
}

//Home
//@title		Home()
//@description	个人主页请求
//@author		zy
//@param		c *gin.Context
//@return
func Home(c *gin.Context) {
	username, _ := c.Get("username")		//获取当前登录的username
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"message": username.(string) + "的主页",
	})
}

//PostArticle
//@title		PostArticle()
//@description	发布文章请求
//@author		zy
//@param		c *gin.Context
//@return
func PostArticle(c *gin.Context) {
	var ArticleInfo userinformation.Article
	username, _ := c.Get("username")		//获取当前登录的username
	ArticleInfo.Username = username.(string)
	err := c.ShouldBind(&ArticleInfo)			//参数绑定

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 2001,
			"message": "无效的参数",
		})
		return
	}

	//不能发布空的文章
	if ArticleInfo.Content == "" || ArticleInfo.Title == ""{				//内容或标题为空      提示错误
		c.JSON(http.StatusForbidden, gin.H{
			"status": http.StatusForbidden,
			"message": "发表失败",
			"reason": "title or content 为空",
		})
		return
	}

	// 是否成功发表文章
	if common.InsertArticle(ArticleInfo) {
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
			"message": "发表成功",
			"data": gin.H{"article": ArticleInfo},
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"message": "发表失败！",
		})
	}
}

//Delete
//@title		Delete()
//@description	删除文章请求
//@author		zy
//@param		c *gin.Context
//@return
func Delete(c *gin.Context) {
	var ArticleInfo userinformation.Article
	username, _ := c.Get("username")		//获取当前登录的username
	err := c.ShouldBind(&ArticleInfo)			//参数绑定

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 2001,
			"message": "无效的参数",
		})
		return
	}

	if username != ArticleInfo.Username {
		c.JSON(http.StatusForbidden, gin.H{
			"status": http.StatusForbidden,
			"message": "你没有权限删除别人的文章",
		})
		return
	}
	if common.DeleteArticle(ArticleInfo) {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"message": "删除成功",
		})
	} else {
		c.JSON(http.StatusForbidden, gin.H{
			"status": http.StatusForbidden,
			"message": "删除失败",
		})
	}

}

//Like
//@title		Like()
//@description	点赞文章请求
//@author		zy
//@param		c *gin.Context
//@return
func Like(c *gin.Context)  {
	var ArticleInfo userinformation.Article
	username, _ := c.Get("username")  //username为当前用户id
	err := c.ShouldBind(&ArticleInfo)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 2001,
			"message": "无效的参数",
		})
		return
	}

	if username == ArticleInfo.Username {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"message": "你不能给自己点赞！",
		})
		return
	}
	n := common.LikeArticle(ArticleInfo)
	if n == 1 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": http.StatusInternalServerError,
			"message": "Some errors in sql",
		})
		return
	} else if n == 2 {
		c.JSON(http.StatusForbidden, gin.H{
			"status": http.StatusForbidden,
			"message": "点赞失败，没有相应的文章",
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"message": "给" + ArticleInfo.Username + "点赞成功！",
		})
	}
}

//MessageToOther
//@title		MessageToOther()
//@description	给其他用户留言请求
//@author		zy
//@param		c *gin.Context
//@return
func MessageToOther(c *gin.Context) {
	var MsgTo userinformation.Msg
	username, _ := c.Get("username")  //username为当前用户id
	err := c.ShouldBind(&MsgTo)
	MsgTo.Username = username.(string)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 2001,
			"message": "无效的参数",
		})
		return
	}

	if MsgTo.OtherUsername == "" || MsgTo.Message == "" {
		c.JSON(http.StatusForbidden, gin.H{
			"status": http.StatusForbidden,
			"message": "留言失败",
			"reason": "目标username或留言内容为空",
		})
		return
	}

	if common.MessageInsert(MsgTo) {
		c.JSON(http.StatusOK, gin.H{
			"status": 200,
			"message": "留言成功",
			"data": gin.H{"article": MsgTo},
		})
		return
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status": http.StatusBadRequest,
			"message": "留言失败！",
		})
	}
}

//DeleteMsg
//@title		DeleteMsg()
//@description	删除留言请求
//@author		zy
//@param		c *gin.Context
//@return
func DeleteMsg(c *gin.Context) {
	var MsgTo userinformation.Msg
	username, _ := c.Get("username")  //username为当前用户id
	err := c.ShouldBind(&MsgTo)
	MsgTo.Username = username.(string)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 2001,
			"message": "无效的参数",
		})
		return
	}


	if MsgTo.OtherUsername == "" || MsgTo.Message == "" {
		c.JSON(http.StatusForbidden, gin.H{
			"status": http.StatusForbidden,
			"message": "删除留言失败",
			"reason": "目标username或留言内容为空",
		})
		return
	}

	if !common.MessageDelete(MsgTo) {
		c.JSON(http.StatusForbidden, gin.H{
			"status": http.StatusForbidden,
			"message": "some errors in sql",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"message": "删除留言成功",
	})
}
