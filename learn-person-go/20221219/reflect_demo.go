package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	reflect_type(x)

}

func reflect_type(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Println("类型是：", t)
	//kind() 可以获取具体的类型
	K := t.Kind()
	fmt.Println(K)
	switch K {
	case reflect.Float64:
		fmt.Printf("a is float64\n")
	case reflect.String:
		fmt.Println("string")
	}
}
