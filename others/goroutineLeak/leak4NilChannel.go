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

	var ch chan int
	go func() {
		<-ch
		// ch<-  //都将会导致阻塞
	}()
}
