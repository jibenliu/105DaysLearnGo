package main

import (
	"fmt"
	"net/http"
	//_ "net/http/pprof"
)

func main() {
	//http.ListenAndServe("0.0.0.0:6060", nil)
	//defer func() {
	//	time.Sleep(time.Second * 2)
	//	fmt.Println("the number of goroutines: ", runtime.NumGoroutine())
	//}()
	for {
		//httpClient := http.Client{
		//	Timeout: time.Second * 15,
		//}
		go func() {
			//_,err:=httpClient.Get("https://www.xxx.com/")
			_, err := http.Get("https://www.xxx.com/")
			if err != nil {
				fmt.Printf("http.get err: %v\n", err)
			}
		}()
	}
}
