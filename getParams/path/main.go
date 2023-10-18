package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	r := gin.Default()
	r.POST("/user/search/:username/:address", func(c *gin.Context) {
		// 获取参数
		username := c.Param("username")
		address := c.Param("address")
		c.JSON(http.StatusOK, gin.H{
			"message":  "okk",
			"username": username,
			"address":  address,
		})
	})

	r.Run(":8000") //可将端口号改为 8000
}
