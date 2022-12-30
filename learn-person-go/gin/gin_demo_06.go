package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//自定义http配置
func main() {
	router := gin.Default()
	http.ListenAndServe(":8084", router)
	//------------------------http自定义配置------------------------------
	s := &http.Server{
		Addr:           ":8084",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
