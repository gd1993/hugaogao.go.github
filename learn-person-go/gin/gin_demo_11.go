package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

//在中间件中使用 Goroutine
//多个协程之间进行信息传递，包括层层的递进顺序传递
//当顶层context被取消或者超时的时候，所有从这个顶层创建的context也应该结束
//当在中间件或 handler 中启动新的 Goroutine 时，不能使用原始的上下文，必须使用只读副本
func main() {
	r := gin.Default()
	r.GET("/async", func(c *gin.Context) {
		//创建在goroutine使用的副本
		cCp := c.Copy()
		go func() {
			//用time.sleep()模拟一个长任务
			time.Sleep(5 * time.Second)
			log.Println("Done in path:", cCp.Request.URL.Path)
			c.JSON(http.StatusOK, "success")
		}()
	})

	r.GET("/sync", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		log.Println("Done in path:", c.Request.URL.Path)
	})

	r.Run(":8091")
}
