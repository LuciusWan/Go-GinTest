package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 创建一个go的服务
	r := gin.Default()

	// 定义一个路由，处理 GET 请求
	r.GET("/helloGin", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"Hello Gin": "我是奶龙",
		})
	})
	// 定义一个路由，处理 POST 请求
	r.POST("/hello", func(c *gin.Context) {
		fmt.Println("POST方法被使用了")
		// 解析请求体，假设请求体是 JSON 格式
		var requestBody map[string]string
		if err := c.BindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body111"})
			return
		}
		// 获取用户名和密码
		username := requestBody["username"]
		password := requestBody["password"]

		// 假设正确的用户名和密码
		correctUsername := "admin"
		correctPassword := "123456"

		// 验证用户名和密码
		if username == correctUsername && password == correctPassword {
			c.JSON(200, gin.H{"message": "Login successful"})
		} else {
			c.JSON(401, gin.H{"error": "Invalid username or password222"})
		}
	})

	// 启动 HTTP 服务器，监听在 8080 端口
	r.Run(":8080")
}
