package main

import "fmt"

func main() {
	next := incSeq()

	fmt.Printf("初始值 = %d\n", next)

	for i := 0; i < 1000; i++ {
		fmt.Printf("A   第 %d 次迭代后, 值 = %d\n", i, next())
	}

	for i := 0; i < 1000; i++ {
		fmt.Printf("B   第 %d 次迭代后, 值 = %d\n", i, next())
	}

}

func incSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
