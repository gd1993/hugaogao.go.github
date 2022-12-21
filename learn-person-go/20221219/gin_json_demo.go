package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//定义接收数据的结构体
type Login struct {
	User     string `form:"usernmae" json:"user" uri:"user" xml:"user" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" xml:"password" binding:"required"`
}

func main() {
	//创建路由
	r := gin.Default()
	//json 绑定
	r.POST("/loginForm", func(c *gin.Context) {
		//声明变量
		var form Login
		if err := c.Bind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		//判断用户和密码
		if form.User != "root" || form.Password != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"ststus": "304"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})
	r.Run(":8000")
}
