package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/login", func(c *gin.Context) {
		c.SetCookie("abc", "123", 60, "/", "localhost", false, true)
		c.String(200, "Login success!")
	})

	r.GET("/home", AuthMiddleWare(), func(c *gin.Context) {
		c.JSON(200, gin.H{"data": "home"})
	})
	r.Run(":8000")

}

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		//回去客户端cookie并校验
		if cookie, err := c.Cookie("abc"); err == nil {
			if cookie == "123" {
				c.Next()
				return
			}
		}
		//返回错误
		c.JSON(http.StatusUnauthorized, gin.H{"error": "err"})
		//若验证不通过，不再调用后续的函数处理
		c.Abort()
		return
	}
}
