package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.POST("/user/search", func(c *gin.Context) {
		body, _ := c.GetRawData() // 从 c.Request.Body 中读取请求数据

		// 定义map或结构体
		var m map[string]interface{}
		//反序列化
		_ = json.Unmarshal(body, &m)

		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
			"data":    m,
		})

	})

	r.Run(":8000") //可将端口号改为 8000
}
