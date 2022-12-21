package main

import (
	"context"
	"golang.org/x/time/rate"
	"log"
	"time"
)

func main() {
	l := rate.NewLimiter(5, 20)
	log.Println(l.Limit(), l.Burst())

	for i := 0; i < 100; i++ {
		//阻塞等待，直到去到一个token
		log.Println("before wait")
		c, _ := context.WithTimeout(context.Background(), time.Second*2)
		if err := l.Wait(c); err != nil {
			log.Println("limiter wait err:" + err.Error())
		}
		log.Println("after Wait")

		//返回需要等待多久才有新的token，这样可以等待制定时间执行任务

		r := l.Reserve()
		log.Println("reserve Delay:", r.Delay())

		//判断当前知否可以取到token
		a := l.Allow()
		log.Println("Allow:", a)
	}

}
