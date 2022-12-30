package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func main() {
	route := gin.Default()
	//单个文件上传
	route.POST("/upload", func(c *gin.Context) {
		//formfile  获取请求的表单名称
		file, _ := c.FormFile("file")
		//存储的路径
		dst := fmt.Sprintf("./file/%s", file.Filename)
		c.SaveUploadedFile(file, dst)
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	//多文件上传
	route.POST("/uploads", func(c *gin.Context) {
		//获取上传的文件表单
		form, _ := c.MultipartForm()
		//文件上传的keyName
		files := form.File["file"]
		//循环存入文件
		for _, file := range files {
			dst := fmt.Sprintf("./file/%s", file.Filename)
			c.SaveUploadedFile(file, dst)
		}
		c.JSON(http.StatusOK, "Success")
	})
	//单个删除文件
	route.POST("/delete", func(c *gin.Context) {
		//filePath := c.PostForm("filePath")
		//删除文件的路径
		file, err1 := c.FormFile("file")
		//所在文件夹的路径
		dst := fmt.Sprintf("./file/%s", file.Filename)
		if err1 == nil {
			//删除方法
			err := os.Remove(dst)
			if err != nil {
				c.JSON(http.StatusFailedDependency, gin.H{"message": "文件删除失败"})
			} else {
				c.JSON(http.StatusOK, gin.H{"message": "文件删除成功"})
			}
		}
	})

	// 路由组使用 gin.BasicAuth() 中间件
	// gin.Accounts 是 map[string]string 的一种快捷方式
	authorized := route.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))

	// /admin/secrets 端点
	// 触发 "localhost:8080/admin/secrets
	authorized.GET("/secrets", func(c *gin.Context) {
		// 获取用户，它是由 BasicAuth 中间件设置的
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})
	route.Run(":8098")
}

//BasicAuth中间件
var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}
