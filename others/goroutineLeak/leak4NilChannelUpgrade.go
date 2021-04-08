package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	defer func() {
		time.Sleep(time.Second)
		fmt.Println("the number of goroutines:", runtime.NumGoroutine())
	}()
	done := make(chan struct{})
	var ch chan int

}
