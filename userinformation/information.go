//@Title		information.go
//@Description	基本对象信息
//@Author		zy
//@Update		2021.12.5

package userinformation


// UserInfo 用户对象 存储用户名及密码
type UserInfo struct {
	Username string	`form:"username"`	//反射绑定
	Password string `form:"password"`	//反射绑定
}

// Article 文章对象 存储作者、文章标题、文章内容
type Article struct {
	Username string `form:"username"`	//反射绑定
	Title string `form:"title"`			//反射绑定
	Content string `form:"content"`		//反射绑定
}

type Msg struct {
	Username string `form:"username"`					//反射绑定
	OtherUsername string `form:"otherusername"`			//反射绑定
	Message string `form:"message"`						//反射绑定
}