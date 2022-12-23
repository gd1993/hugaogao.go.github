package main

import "sync"

type Lock struct {
	c chan struct{}
}

func (l Lock) Lock() bool {
	result := false

	select {
	case <-l.c:
		result = true
	default:
	}
	return result
}

func (l Lock) UnLock() {
	l.c <- struct{}{}
}

func NewLock() Lock {
	var l Lock
	l.c = make(chan struct{}, 1)

	l.c <- struct{}{}
	return l
}

var count int

func main() {
	var lock = NewLock()
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if !lock.Lock() {
				println("lock failed")
				return
			}
			count++
			println("current counter:", count)
			lock.UnLock()
		}()
	}
	wg.Wait()
}
