package main

import (
	"blockchain/routeas"
	"github.com/gin-gonic/gin"
	ginsession "github.com/go-session/gin-session"
)

func main() {
	// 创建一个gin对象：r
	r := gin.Default()

	r.Use(ginsession.New()) //使用session插件


	// 设置路由规则： 当浏览器 请求 http://localhost:8080/ping GET 使用func(c *gin.Context)处理
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// 使用group
	v1 := r.Group("/api")
	{
		v1.POST("register", routeas.Register)
		v1.POST("login", routeas.Login)
		v1.POST("send", routeas.Send)
		v1.POST("receive", routeas.Receive)
		v1.POST("receivee", routeas.Receivee)
		//v1.POST("lochs", routeas.LocalHashFile)
		v1.POST("rehash", routeas.GetHashFile)
		v1.GET("logout", routeas.Logout)
	}

	//用户的登录注册
	r.POST("register", routeas.Register)
	r.POST("login", routeas.Login)

	//发送关键词、加密文件
	r.POST("send",routeas.Send)

	//发送陷门
	r.POST("receive",routeas.Receive)


	////验证文本哈希值
	r.POST("rehash", routeas.GetHashFile)

	//退出系统
	r.GET("logout", routeas.Logout)

	r.Run(":8282") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

	/*-----------------------------*/

}

