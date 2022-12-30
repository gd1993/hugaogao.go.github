package main

import "github.com/gin-gonic/gin"

func main() {
	route := gin.Default()
	//全局中间件
	//Logger中间件将日志写入gin.DefaultWriter，即使你将 GIN_MODE 设置为 release
	route.Use(gin.Logger())
	//Recover中间件会处理全部的panic
	route.Use(gin.Recovery())
	//可以为每个路添加不限制的中间件
	route.GET("/benchmark", MyBenchLogger(), benchEndpoint)
	// 认证路由组
	// authorized := r.Group("/", AuthRequired())
	// 和使用以下两行代码的效果完全一样:
	authorized := route.Group("/")
	// 路由组中间件! 在此例中，我们在 "authorized" 路由组中使用自定义创建的
	// AuthRequired() 中间件
	authorized.Use(AuthRequired())
	{
		authorized.POST("/login", loginEndpoint)
		authorized.POST("/submit", submitEndpoint)
		authorized.POST("/read", readEndpoint)

		// 嵌套路由组
		testing := authorized.Group("testing")
		testing.GET("/analytics", analyticsEndpoint)
	}

	route.Run(":8089")

}

func analyticsEndpoint(context *gin.Context) {

}

func readEndpoint(context *gin.Context) {

}

func submitEndpoint(context *gin.Context) {

}

func loginEndpoint(context *gin.Context) {

}

func AuthRequired() gin.HandlerFunc {
	return nil
}

func benchEndpoint(context *gin.Context) {

}

func MyBenchLogger() gin.HandlerFunc {
	return nil
}
