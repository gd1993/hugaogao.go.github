package main

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"os"
)

func main() {
	node, err := snowflake.NewNode(1)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	for i := 0; i < 20; i++ {
		id := node.Generate()
		//fmt.Printf("Int64  ID: %d\n", id)
		fmt.Printf("String ID: %s\n", id)
		//fmt.Printf("Base2  ID: %s\n", id.Base2())
		fmt.Printf("Base64 ID: %s\n", id.Base64())

		//fmt.Printf("ID Time  : %d\n", id.Time())

		//fmt.Printf("ID Node  : %d\n", id.Node())

		//fmt.Printf("ID Step  : %d\n", id.Step())

		//fmt.Println("---------------------------------------------")
	}

}
