package main

import (
	"fmt"
	"github.com/gin-gonic/autotls"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/acme/autocert"
	"log"
	"net/http"
)

func main() {
	route := gin.Default()
	// 提供 unicode 实体
	route.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"html": "Hello, world!",
		})
	})

	// 提供字面字符
	route.GET("/purejson", func(c *gin.Context) {
		c.PureJSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
		//c.AsciiJSON(200, gin.H{
		//	"html": "<b>Hello, world!</b>",
		//})
	})

	route.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("test", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")
		fmt.Printf("id: %s; page: %s; name: %s; message: %s \n", id, page, name, message)
	})

	route.GET("/test", func(c *gin.Context) {
		//http  重定向
		c.Redirect(http.StatusMovedPermanently, "https://www.baidu.com/")
		//路由重定向
		c.Request.URL.Path = "/test2"
		route.HandleContext(c)
	})
	route.GET("/test2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "Success"})
	})

	//从reader读取数据 浏览器下载
	route.GET("/read", func(c *gin.Context) {
		response, err := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")
		if err != nil || response.StatusCode != http.StatusOK {
			c.Status(http.StatusServiceUnavailable)
			return
		}
		reader := response.Body
		contentLength := response.ContentLength
		contentType := response.Header.Get("Content-Type")
		//浏览器下载显示
		//Content-disposition:属性名
		//attachment:表示以附件方式下载，如果要在页面中打开，可以改为inline
		extraHeaders := map[string]string{
			"Content-Disposition": `attachment;filename="gopher.png"`,
		}
		c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
	})

	//静态文件服务
	route.Static("/assets", "./assets")
	//加载单个文件
	route.StaticFile("/favicon.ico", "./resources/favicon.ico")
	//加载资源目录
	route.StaticFS("/more_static", http.Dir("my_file_system"))

	//设置cookie   获取cookie文件
	route.GET("/cookie", func(c *gin.Context) {
		cookie, err := c.Cookie("gin1_cookie") //第一次为空，设置成功以后下次就有值
		if err != nil {
			cookie = "NotSet"
			c.SetCookie("gin1_cookie", "test", 3600, "/", "localhost", false, true)
		}
		//获取到cookie的所有值
		value := c.Request.Header.Get("Cookie")
		fmt.Println("header value :", value)
		fmt.Printf("cookie  value: %s\n", cookie)
		c.JSON(http.StatusOK, cookie)
	})
	// Let's Encrypt  一行式LetsEncrypt证书, 处理https ，定义那些域名需要证书验证
	log.Fatal(autotls.Run(route, "example1.com", "example2.com"))
	//自定义Encrypt
	m := autocert.Manager{
		Prompt:     autocert.AcceptTOS,
		HostPolicy: autocert.HostWhitelist("example/com"),
		Cache:      autocert.DirCache(""),
	}
	autotls.RunWithManager(route, &m)
	//route.Run(":8096")

}
