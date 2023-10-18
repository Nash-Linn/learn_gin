package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/getjson1", func(c *gin.Context) {
		// 方式1：自己拼接JSON
		c.JSON(http.StatusOK, gin.H{"message": "Hello world!"})
	})

	r.GET("/getjson2", func(c *gin.Context) {
		//方法2：使用结构体
		type Msg struct {
			Name    string `json:"user"`
			Message string
			Age     int
		}
		var msg Msg
		msg.Name = "Nash"
		msg.Message = "Hello world!"
		msg.Age = 18
		c.JSON(http.StatusOK, msg)
	})

	r.Run(":8000") //可将端口号改为 8000
}
