package main

import "sync"

//全局变量
var counter int

func main() {
	//第一版    随机，每次数据不一样
	//var wg sync.WaitGroup
	//var lock sync.Mutex
	//
	//for i := 0; i < 1000; i++ {
	//	wg.Add(1)
	//	go func() {
	//		defer wg.Done()
	//		counter++
	//	}()
	//}
	//wg.Wait()
	//println(counter)

	//第二版
	var wg sync.WaitGroup
	var lock sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			lock.Lock()
			counter++
			lock.Unlock()
		}()
	}
	wg.Wait()
	println(counter)
}
