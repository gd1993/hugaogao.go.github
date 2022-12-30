package main

import "fmt"

func main() {

	numbers := []int{100, 200, 300}
	double(numbers)
	fmt.Printf("使用切片变量作为函数参数执行完成后, n = %d\n", numbers) // 切片元素已经发生变化

	numbers1 := [3]int{100, 200, 300}
	double1(numbers1)
	fmt.Printf("使用数组变量作为函数参数执行完成后, n = %d\n", numbers1) // 可以看到，数组元素并未发生变化
}

func double(numbers []int) {
	for i := range numbers {
		numbers[i] *= 2
	}
}

func double1(numbers [3]int) {
	for i := range numbers {
		numbers[i] *= 2
	}
}
