package main

import (
	"bytes"
	"fmt"
	"github.com/go-redis/redis"
	"runtime"
	"strconv"
	"sync"
	"time"
)

var client = redis.NewClient(&redis.Options{
	Addr:     "192.168.139.140:6379",
	Password: "",
	DB:       0,
})

func lock(myfunc func()) {
	var lockKey = "mylockr"
	gid := getGID()
	lockSuccess, err := client.SetNX(lockKey, gid, time.Second*10).Result()
	if err != nil || !lockSuccess {
		fmt.Println("get lock fail")
		return
	} else {
		fmt.Println("get lock")
	}
	myfunc()
	_, err = client.Del(lockKey).Result()
	if err != nil {
		fmt.Println("unlock fail")
	} else {
		fmt.Println("unlock")
	}

	//用lua实现原子操作
	//var luaScript = redis.NewScript(`
	//    if redis.call("get", KEYS[1]) == ARGV[1]
	//        then
	//            return redis.call("del", KEYS[1])
	//        else
	//            return 0
	//    end
	//`)
	//rs, _ := luaScript.Run(client, []string{lockKey}, gid).Result()
	//if rs == 0 {
	//	fmt.Println("unlock fail")
	//} else {
	//	fmt.Println("unlock")
	//}
}

var counter int64

func incr() {
	counter++
	fmt.Printf("after incr is %d\n", counter)
}

var wg sync.WaitGroup

func main() {
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			lock(incr)
		}()
	}
	wg.Wait()
	fmt.Printf("final counter is %d \n", counter)
}

func getGID() uint64 {
	b := make([]byte, 64)
	b = b[:runtime.Stack(b, false)]
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}