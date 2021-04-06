package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func postRequest() {
	resp, err := http.Post("https://httpbin.org/post?name=Regan&money=0",
		"application/x-www-form-urlencoded",
		strings.NewReader("姓名=ReganYue&人民币=0&心愿=发财发财"),
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(string(bytes))
}

func getRequest() {
	resp, err := http.Get("http://www.baidu.com/s?wd=ReganYue")
	if err != nil {
		log.Fatal(err.Error())
	}

	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(string(bytes))
}
