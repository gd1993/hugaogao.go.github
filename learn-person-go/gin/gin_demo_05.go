package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//关闭高亮
	gin.DisableConsoleColor()
	//强制高亮    高亮显示的是状态码和请求方法类型
	//gin.ForceConsoleColor()
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	router.Run(":8083")

}
