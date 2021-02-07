package main

import (
	"fmt"
	"io/ioutil"
	"runtime"
	"time"
)

var query = "test"
var matches int

var workerCount = 0
var maxWorkerCount = runtime.NumCPU()
var searchRequest = make(chan string)
var workerDone = make(chan bool)
var foundMatch = make(chan bool)

/**
1. 遍历一个未关闭的channel会造成死循环

2. 即使关闭了一个非空通道，我们仍然可以从通道里面接收到未读取的数据

3. 可以这样理解，close()函数会往channel中压入一条特殊的通知消息，可以用来通知channel接收者不会再收到数据。所以即使channel中有数据也可以close()而不会导致接收者收不到残留的数据

4. channel不需要通过close释放资源，只要没有goroutine持有channel，相关资源会自动释放
 */
func main() {
	start := time.Now()
	workerCount = 1
	go search("/data", true)
	waitForWorkers()
	fmt.Println(matches, "matches")
	fmt.Println(time.Since(start))
}

func waitForWorkers() {
	for {
		select {
		case path := <-searchRequest:
			workerCount++
			go search(path, true)
		case <-workerDone:
			workerCount--
			if workerCount == 0 {
				return
			}
		case <-foundMatch:
			matches++
		default:
		}
	}
}

func search(path string, master bool) {
	files, err := ioutil.ReadDir(path)
	if err == nil {
		for _, file := range files {
			name := file.Name()
			if name == query {
				foundMatch <- true
			}
			if file.IsDir() {
				if workerCount < maxWorkerCount {
					searchRequest <- path + name + "/"
				} else {
					search(path+name+"/", false)
				}
			}
		}
		if master {
			workerDone <- true
		}
	}
}
