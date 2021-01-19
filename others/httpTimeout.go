package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var (
	//c = &http.Client{} //此种写法会造成请求被长时间挂起
	c = &http.Client{
		Timeout: time.Second * 30, //加上超时限制，但是所有的请求都会加上此限制
	}
)

func base() {
	req, err := http.NewRequest("GET", "http://google.com", nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close() //必须要有，不然会goroutine泄漏
	b, _ := ioutil.ReadAll(res.Body)
	fmt.Println(b)
}

func upgrade() {
	req, err := http.NewRequest("GET", "http://google.com", nil)
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute) //针对当前请求设置一分钟超时
	defer cancel()
	req = req.WithContext(ctx)
	res, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close() //必须要有，不然会goroutine泄漏
	b, _ := ioutil.ReadAll(res.Body)
	fmt.Println(b)
}
func main() {
	base()    //基础版
	upgrade() //高级版设置独立的超时时间
}
