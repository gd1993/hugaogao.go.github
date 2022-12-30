package main

import (
	"github.com/gin-gonic/gin"
)

type Person1 struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

func main() {

	//for i := 0; i < 100; i++ {
	//	//uuid生成    100个以内没重复的
	//	valueUUID, _ := uuid.NewUUID()
	//	fmt.Println("valueUUID:", valueUUID.String())
	//
	//	//循环100 有重复的
	//	uuid1 := uuid.New()
	//	fmt.Println("uuid1:", uuid1.String())
	//}

	route := gin.Default()
	route.GET("/:name/:id", func(c *gin.Context) {
		var person Person1
		if err := c.ShouldBindUri(&person); err != nil {
			c.JSON(400, gin.H{"msg": err})
			return
		}
		c.JSON(200, gin.H{"name": person.Name, "uuid": person.ID})
	})
	route.Run(":8082")

}
