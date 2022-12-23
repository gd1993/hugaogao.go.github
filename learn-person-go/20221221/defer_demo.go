package main

import "fmt"

func f1() (result int) {
	defer func() {
		result++
	}()
	return 0
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return r
}

func main() {
	fmt.Println("f1:", f1())
	fmt.Println("f2:", f2())
	fmt.Println("f3:", f3())

}
