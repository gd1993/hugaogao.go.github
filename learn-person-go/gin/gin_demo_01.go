package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/someJson", func(c *gin.Context) {
		data := map[string]string{
			"lang": "GO语言",
			"tag":  "success",
		}
		//c.JSON(http.StatusOK, data)
		c.AsciiJSON(http.StatusOK, data)
	})
	//监听并在8080上运行
	r.Run(":8080")

}
