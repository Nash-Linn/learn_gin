package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	router := gin.Default()
	// 处理multipart forms提交文件时默认的内存限制是32MiB
	// 可以通过下面的方式修改该限制
	// router.MaxMultipartMemory = 8 << 20  //   8 MiB  8*2^20  8*1024*1024

	router.POST("/upload", func(c *gin.Context) {
		// 单文件
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err.Error(),
			})
			return
		}

		log.Printf("上传的文件名为：%s", file.Filename)

		//fmt.Sprintf 格式化输出
		dst := fmt.Sprintf("./%s", file.Filename)

		// 上传文件到指定的目录
		_ = c.SaveUploadedFile(file, dst)
		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("'%s' uploaded!", file.Filename),
		})

	})

	router.Run(":8000") //可将端口号改为 8000

}
