package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//创建路由
	r := gin.Default()
	//绑定路由规则，执行函数
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello world")
	})
	//监听端口
	r.Run(":8080")

}
