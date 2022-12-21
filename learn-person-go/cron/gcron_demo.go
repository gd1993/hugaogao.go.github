package main

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
	"time"
)

func task() {
	fmt.Println("I am runnning task.", time.Now())
}

func superWang() {
	fmt.Println("I am runnning superWang.", time.Now())
}
func main() {
	s := gocron.NewScheduler()
	s.Every(1).Seconds().Do(task)
	s.Every(4).Seconds().Do(superWang)
	sc := s.Start()
	go test(s, sc)
	<-sc
}

func test(s *gocron.Scheduler, sc chan bool) {
	time.Sleep(8 * time.Second)
	s.Remove(task) //remove task
	time.Sleep(8 * time.Second)
	s.Clear()
	fmt.Println("All task removed")
	close(sc) // close the channel
}
