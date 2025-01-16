package main

import (
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"net/http"
)

func main() {
	// 创建一个go的服务
	r := gin.Default()
	//加载前端页面
	r.LoadHTMLGlob("static/*")
	// 定义一个路由，处理 GET 请求
	r.GET("/helloGin", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Hello Gin": "我是奶龙",
		})
	})
	r.POST("/user", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "发送了POST请求",
		})
	})
	//响应一个页面给前端
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"我是奶龙": "我才是奶龙",
		})
	})
	//?方式来传递请求
	r.GET("/user/info", func(c *gin.Context) {
		//?请求时扫描用Query
		var userid = c.Query("userid")
		var username = c.Query("username")
		c.JSON(200, gin.H{
			"userid":   userid,
			"username": username,
		})
	})
	//RESTful方式来传递请求
	r.GET("/user/info/:userid/:username", func(c *gin.Context) {
		//RESTful请求时用Param扫描
		var userid = c.Param("userid")
		var username = c.Param("username")
		c.JSON(200, gin.H{
			"userid":   userid,
			"username": username,
		})
	})
	//前端发出post请求
	r.POST("/john", func(c *gin.Context) {
		//b用来接收前端传来的数据，返回来的数据分为context和error，error可以不处理
		b, _ := c.GetRawData()
		//创建键值对来接收数据
		var m map[string]interface{}
		//解析b中的JSON数据，并且将其存入m这个键值对中，键值对是引用类型
		_ = json.Unmarshal(b, &m)
		//最终将POST来的结果再返回给前端
		c.JSON(200, m)
	})
	//接收前端的信息
	r.POST("/user/add", func(c *gin.Context) {
		//这里的form是p标签里的name
		username := c.PostForm("username")
		password := c.PostForm("password")
		c.JSON(200, gin.H{
			"msg":      "ok",
			"username": username,
			"password": password,
		})
	})
	//重定向，统一404处理
	r.NoRoute(func(c *gin.Context) {
		c.HTML(404, "404.html", nil)
	})
	//路由组定义
	//相当于SpringBoot里的@RequestMapping("/depts")，后面的路由都不用加上/depts
	userGroup := r.Group("/depts")
	{
		//这个请求路径就是/depts/info
		userGroup.GET("/info", func(c *gin.Context) {})
		userGroup.GET("/show", func(c *gin.Context) {})
		userGroup.POST("/add", func(c *gin.Context) {})
	}
	orderGroup := r.Group("/order")
	{
		orderGroup.GET("/info", func(c *gin.Context) {})
		orderGroup.GET("/show", func(c *gin.Context) {})
		orderGroup.POST("/add", func(c *gin.Context) {})
	}
	// 启动 HTTP 服务器，监听在 8080 端口
	r.Run(":8080")
}
