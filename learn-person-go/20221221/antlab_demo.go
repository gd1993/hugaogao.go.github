package main

import (
	"github.com/antlabs/timer"
	"log"
	"sync"
	"time"
)

//一次性定时器
func after(tm timer.Timer) {
	var wg sync.WaitGroup
	wg.Add(1)
	defer wg.Wait()
	go func() {
		defer wg.Done()
		tm.AfterFunc(1*time.Second, func() {
			log.Printf("after 1 second\n")
		})
		defer wg.Add(1)
	}()

	go func() {
		defer wg.Done()
		tm.AfterFunc(10*time.Second, func() {
			log.Printf("after 10 seconds\n")
		})
	}()

	go func() {
		defer wg.Done()
		tm.AfterFunc(30*time.Second, func() {
			log.Printf("after 30 seconds\n")
		})
	}()

	go func() {
		defer wg.Done()
		tm.AfterFunc(time.Minute, func() {
			log.Printf("after 1 minute\n")
		})
	}()

	go func() {
		defer wg.Done()
		tm.AfterFunc(time.Minute+30*time.Second, func() {
			log.Printf("after 1 minute and 30 seconds\n")
		})
	}()

	go func() {
		defer wg.Done()
		tm.AfterFunc(2*time.Minute+45*time.Second, func() {
			log.Printf("after 2 minutes and 45 seconds\n")
		})
	}()
}

// 周期性定时器
func schedule(tm timer.Timer) {
	tm.ScheduleFunc(500*time.Millisecond, func() {
		log.Printf("schedule 500 milliseconds\n")
	})

	tm.ScheduleFunc(time.Second, func() {
		log.Printf("schedule 1 second\n")
	})

	tm.ScheduleFunc(20*time.Second, func() {
		log.Printf("schedule 20 seconds\n")
	})

	tm.ScheduleFunc(1*time.Minute, func() {
		log.Printf("schedule 1 minute\n")
	})
}
func main() {
	tm := timer.NewTimer()
	defer tm.Stop()

	go after(tm)
	//go schedule(tm)

	go func() {
		time.Sleep(2*time.Minute + 50*time.Second)
		tm.Stop()
	}()

	tm.Run()
}
