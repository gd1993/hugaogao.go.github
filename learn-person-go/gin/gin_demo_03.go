package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Person struct {
	Name     string `form:"name"`
	Address  string `form:"address"`
	Birthday string `form:"birthday"`
}

func main() {
	r := gin.Default()
	r.GET("/person", getPerson)
	r.Run(":8081")
}

func getPerson(c *gin.Context) {
	var person Person
	d := true
	if c.ShouldBind(&person) == nil {
		log.Println(person.Name)
		log.Println(person.Address)
		log.Println(person.Birthday)
	} else {
		d = false
	}
	c.JSON(http.StatusOK, d)
	//c.String(http.StatusOK, "success")
}
