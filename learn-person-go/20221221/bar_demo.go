package main

import (
	"github.com/cheggaaa/pb/v3"
	"time"
)

func main() {
	count := 1000
	bar := pb.StartNew(count)

	for i := 0; i < count; i++ {
		bar.Increment()
		time.Sleep(time.Millisecond)
	}
	bar.Finish()

}
