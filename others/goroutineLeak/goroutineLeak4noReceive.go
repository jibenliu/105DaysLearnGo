package main

import (
	"fmt"
	"runtime"
	"time"
)

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			fmt.Printf("n is %d \n", n)
			out <- n
		}
		close(out)
	}()
	return out
}

func main() {
	defer func() {
		fmt.Println("the number of goroutines: ", runtime.NumGoroutine())
	}()

	out := gen(1, 2, 3, 5, 8, 6, 9, 10)
	for n := range out {
		fmt.Println(n)
		time.Sleep(5 * time.Second) // done thing, 可能异常中断接收
		if true {                   // if err != nil
			break
		}
	}
}
