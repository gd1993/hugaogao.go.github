package main

import "github.com/gin-gonic/gin"

type Login struct {
	User     string `form:"user" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func main() {
	router := gin.Default()
	router.POST("/login", func(c *gin.Context) {
		// 你可以使用显式绑定声明绑定 multipart form：
		// c.ShouldBindWith(&form, binding.Form)
		// 或者简单地使用 ShouldBind 方法自动绑定：
		var login Login
		//这种情况下自动选择合适的绑定
		if c.ShouldBind(&login) == nil {
			if login.User == "user" && login.Password == "password" {
				c.JSON(200, gin.H{"status": "you are logged in"})
			} else {
				c.JSON(401, gin.H{"status": "unauthorized"})
			}
		} else {
			c.JSON(401, gin.H{"status": "missing params"})
		}
	})
	router.Run(":8093")
}
