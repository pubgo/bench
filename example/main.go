package main

import (
	"fmt"
	"os/signal"
	"runtime"
	"time"
)

func cost(fn func()) time.Duration {
	var t = time.Now()
	fn()
	return time.Since(t)
}

func main() {
	var d = make([]string, 1E9)
	for {
		_ = d

		fmt.Println(cost(func() {
			runtime.GC()
		}))

		time.Sleep(time.Second)
	}
}
