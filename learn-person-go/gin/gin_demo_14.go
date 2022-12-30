package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

type GinPerson struct {
	Name    string `form:"name"`
	Address string `form:"address"`
}

func main() {
	router := gin.Default()
	//Any  支持任何请求
	router.Any("/test", startPage)
	router.Run(":8094")

}

func startPage(c *gin.Context) {
	var person GinPerson
	if c.ShouldBindQuery(&person) == nil {
		log.Println("====== Only Bind By Query String ======")
		log.Println("name:", person.Name)
		log.Println("address:", person.Address)
	}
	c.String(200, "Success")

}
