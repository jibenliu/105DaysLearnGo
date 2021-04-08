package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func handle(v int) {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < v; i++ {
		//wg.Add(1)
		//defer wg.Done()
		fmt.Println("脑子进煎鱼了")
		wg.Done()
	}
	wg.Wait()
}

func main() {
	defer func() {
		fmt.Println("goroutines: ", runtime.NumGoroutine())
	}()

	go handle(3)
	time.Sleep(time.Second)
}
