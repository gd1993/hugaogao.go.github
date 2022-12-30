package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

//自定义中间件
func main() {

	r := gin.New()
	r.Use(Logger())

	r.GET("/test", func(c *gin.Context) {
		example := c.MustGet("example").(string)
		log.Println("example:", example)
		c.JSON(http.StatusOK, example)
	})
	r.Run(":8088")
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		//设置example变量
		c.Set("example", "123456")
		log.Println("------------------请求前----------------------------")
		//请求前
		c.Next()
		log.Println("------------------请求后----------------------------")
		//请求后
		latency := time.Since(t)
		log.Println("latency:", latency)

		//获取发送的状态
		statues := c.Writer.Status()
		log.Println("statues:", statues)
	}
}
