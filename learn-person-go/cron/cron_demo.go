package main

import (
	"fmt"
	"github.com/robfig/cron"
)

type GreetingJob struct {
	Name string
}

func (g GreetingJob) Run() {
	fmt.Println("Hello ", g.Name)
}

func main() {
	//i := 0
	c := cron.New()
	//spec := "0 */1 * * * ?" //一分钟运行一次
	//c.AddFunc(spec, func() {
	//	i++
	//	fmt.Println("cron running:", i)
	//})
	c.AddJob("@every 1s", GreetingJob{"hzh"})
	c.Start()
	//time.Sleep(5 * time.Second)
}
