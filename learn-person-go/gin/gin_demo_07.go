package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//自定义 Log 文件
func main() {
	router := gin.New()

	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s]  -  %s  -  %s - %d - %s \n",
			param.ClientIP,
			param.TimeStamp.Format("2006-01-02 04:15:20"),
			param.Method,
			param.Path,
			//param.Request.Proto,
			param.StatusCode,
			//param.Latency,
			//param.Request.UserAgent(),
			param.ErrorMessage,
		)
	}))

	router.Use(gin.Recovery())
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	router.Run(":8085")

}
