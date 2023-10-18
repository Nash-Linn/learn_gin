package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "GET",
		})
	})

	r.POST("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "POST",
		})
	})

	r.PUT("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "PUT",
		})
	})

	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "DELETE",
		})
	})

	// 任意请求方式
	r.Any("/any", func(c *gin.Context) {
		//可以再根据请求方式进行处理
		switch c.Request.Method {
		case "GET":
			c.JSON(200, gin.H{
				"message": "GET any",
			})
		case "POST":
			c.JSON(200, gin.H{
				"message": "POST any",
			})
		}
	})

	r.Run(":8000") //可将端口号改为 8000
}
