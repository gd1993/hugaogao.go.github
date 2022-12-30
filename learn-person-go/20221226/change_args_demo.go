package main

import "fmt"

func main() {

	//传递一个参数
	fmt.Printf("1 == %d\n", sum(1))
	//传递多个参数
	fmt.Printf("1 + 2 + 3 = %d\n", sum(1, 2, 3))
	//传递切片参数
	numbers := []int{1, 2, 3}
	fmt.Printf("1 + 2 + 3 = %d\n", sum(numbers...))
	//初始化map
	var m = make(map[string]string)
	m["k1"] = "k1"
}

func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}
