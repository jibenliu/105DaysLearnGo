package main

import (
	"fmt"
	"runtime"
	"time"
)

func gen(done chan struct{}, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
	}()
	return out
}

func main() {
	defer func() {
		time.Sleep(time.Second)
		fmt.Println("the number of goroutines:", runtime.NumGoroutine())
	}()

	done := make(chan struct{})
	defer close(done)
	out := gen(done, 2, 3)
	for n := range out {
		fmt.Println(n)
		time.Sleep(5 * time.Second)
		if true {
			break
		}
	}
}
