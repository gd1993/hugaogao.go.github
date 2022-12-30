package main

import (
	"fmt"
	"sync"
)

func main() {
	//var mu sync.Mutex
	var rw sync.RWMutex

	m := make(map[int]bool)
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(key int) {
			defer func() {
				wg.Done()
			}()
			//mu.Lock()
			rw.Lock()
			m[key] = true
			//mu.Unlock()
			rw.Unlock()
		}(i)
	}
	wg.Wait() //为了让程序完成
	fmt.Println("Map size = ", len(m))

}
