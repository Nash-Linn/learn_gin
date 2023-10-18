# learn_gin

Go世界里最流行的web框架，基于httprouter开发的web框架。

## Gin框架的安装与使用

### 下载并安装 Gin

```
go get -u github.com/gin-gonic/gin
```

如果遇到错误

go: module github.com/gin-gonic/gin: Get “https://proxy.golang.org/github.com/gin-gonic/gin/@v/list”: dial tcp 142.251.42.241:443: connectex: A connection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to respond.
```
# 设置 goproxy.io 代理。
go env -w GOPROXY="https://goproxy.io"
# 设置 GO111MOUDLE。
go env -w GO111MODULE="on"
```

### 第一个例子 hello world

```go
package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//创建一个默认的路由引擎
	r := gin.Default()
	//GET：请求方式；/hello：请求的路径
	//当客户端以GET方法请求/hello路径时，会执行后面的匿名函数
	r.GET("/hello", func(c *gin.Context) {
		//c.JSON：返回JSON格式的数据
		c.JSON(200, gin.H{
			"message": "Hello world!",
		})
	})
	// 启动HTTP服务，默认在127.0.0.1:8080启动服务
	r.Run(":8000") //可将端口号改为 8000
}
```



### RESTful API

 简单来说，REST的含义就是客户端与Web服务器之间进行交互的时候，使用HTTP协议中的4个请求方法代表不同   的动作。

- `GET`用来获取资源
- `POST`用来新建资源
- `PUT`用来更新资源
- `DELETE`用来删除资源。

只要API程序遵循了REST风格，那就可以称其为RESTful API。目前在前后端分离的架构中，前后端基本都是通过RESTful API来进行交互。

例如，我们现在要编写一个管理书籍的系统，我们可以查询对一本书进行查询、创建、更新和删除等操作，我们在编写程序的时候就要设计客户端浏览器与我们Web服务端交互的方式和路径。按照经验我们通常会设计成如下模式：

| 请求方法 |     URL      |     含义     |
| :------: | :----------: | :----------: |
|   GET    |    /book     | 查询书籍信息 |
|   POST   | /create_book | 创建书籍记录 |
|   POST   | /update_book | 更新书籍信息 |
|   POST   | /delete_book | 删除书籍信息 |

同样的需求我们按照RESTful API设计如下：

| 请求方法 |  URL  |     含义     |
| :------: | :---: | :----------: |
|   GET    | /book | 查询书籍信息 |
|   POST   | /book | 创建书籍记录 |
|   PUT    | /book | 更新书籍信息 |
|  DELETE  | /book | 删除书籍信息 |

```go
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
```



### json渲染

```go
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
```



### 获取参数

#### 获取querystring参数

`querystring`指的是URL中`?`后面携带的参数，例如：`/user/search?username=Nash&address=浙江`。 获取请求的querystring参数的方法如下：

```go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/user/search", func(c *gin.Context) {
		// 获取参数
		username := c.Query("username")

		fmt.Println(username)
		// c.DefaultQuery 可以设置默认值
		//username := c.DefaultQuery("username", "Nash")

		address := c.Query("address")
		fmt.Println(address)

		c.JSON(http.StatusOK, gin.H{
			"message":  "okk",
			"username": username,
			"address":  address,
		})
	})

	r.Run(":8000") //可将端口号改为 8000
}
```

#### 获取form参数

当前端请求的数据通过form表单提交时，例如向`/user/search`发送一个POST请求，获取请求数据的方式如下：

```go
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
```

#### 获取JSON参数

```go
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
```

#### 获取path参数

请求的参数通过URL路径传递,例如：`/user/search/Nash/ZJ`。 获取请求URL路径中的参数的方式如下。

```go
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
```

#### 参数绑定

为了能够更方便的获取请求相关参数，提高开发效率，我们可以基于请求的`Content-Type`识别请求数据类型并利用反射机制自动提取请求中`QueryString`、`form表单`、`JSON`、`XML`等参数到结构体中。 下面的示例代码演示了`.ShouldBind()`强大的功能，它能够基于请求自动提取`JSON`、`form表单`和`QueryString`类型的数据，并把值绑定到指定的结构体对象。

`ShouldBind`会按照下面的顺序解析请求中的数据完成绑定：

1. 如果是 `GET` 请求，只使用 `Form` 绑定引擎（`query`）。
2. 如果是 `POST` 请求，首先检查 `content-type` 是否为 `JSON` 或 `XML`，然后再使用 `Form`（`form-data`）。

```go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Remark   string `form:"remark" json:"remark" binding:"-"` // binding:"-" 非必填
}

func main() {
	router := gin.Default()

	// 绑定JSON的示例
	router.POST("/loginForm", func(c *gin.Context) {
		var login Login

		if err := c.ShouldBind(&login); err == nil {
			fmt.Printf("login info:%#v\n", login)
			c.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	// 绑定form表单示例
	router.POST("/loginForm", func(c *gin.Context) {
		var login Login
		// ShouldBind()会根据请求的Content-Type自行选择绑定器
		if err := c.ShouldBind(&login); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	// 绑定QueryString示例
	router.GET("/loginForm", func(c *gin.Context) {
		var login Login
		// ShouldBind()会根据请求的Content-Type自行选择绑定器
		if err := c.ShouldBind(&login); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	router.Run(":8000")
}
```



### 文件上传

#### 单个文件上传

```go
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
```

#### 多个文件上传
