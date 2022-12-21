package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string
	Hobby string
}

func main() {
	p := Person{"hzh", "男"}
	b, err := json.Marshal(p)
	if err != nil {
		fmt.Println("json err :", err)
	}
	fmt.Println(string(b))

	//格式化输出
	c, err1 := json.MarshalIndent(p, "", "    ")
	if err1 != nil {
		fmt.Println("json err :", err)
	}
	fmt.Println(string(c))
}
