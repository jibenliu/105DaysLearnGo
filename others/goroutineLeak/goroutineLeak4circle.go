package main

import (
	"fmt"
	"runtime"
	"time"
)

func sayHello() {
	for { //死循环，无法关闭该goroutine
		fmt.Println("Hello goroutine")
		time.Sleep(time.Second)
	}
}

func main() {
	defer func() {
		fmt.Println("the number of goroutines :", runtime.NumGoroutine())
		fmt.Println("the number of cpu :", runtime.NumCPU())
		fmt.Println("the number of cgoCall :", runtime.NumCgoCall())
	}()

	go sayHello()
	fmt.Println("Hello main")
}
