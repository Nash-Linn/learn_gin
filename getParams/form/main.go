package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.POST("/user/search", func(c *gin.Context) {
		// c.DefaultPostForm 可以设置默认值
		//c.DefaultPostForm("username", "Nash")

		// 获取参数
		username := c.PostForm("username")

		fmt.Println(username)
		// c.DefaultQuery 可以设置默认值
		//username := c.DefaultQuery("username", "Nash")

		address := c.PostForm("address")
		fmt.Println(address)

		c.JSON(http.StatusOK, gin.H{
			"message":  "okk",
			"username": username,
			"address":  address,
		})
	})

	r.Run(":8000") //可将端口号改为 8000
}
