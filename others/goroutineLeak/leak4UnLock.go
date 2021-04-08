package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	total := 0
	defer func() {
		time.Sleep(time.Second)
		fmt.Println("total: ", total)
		fmt.Println("goroutines: ", runtime.NumGoroutine())
	}()
	var mutex sync.Mutex
	for i := 0; i < 10; i++ {
		go func() {
			mutex.Lock()
			total += 1
		}()
	}
}
