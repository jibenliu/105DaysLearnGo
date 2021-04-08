# refer: [https://zhuanlan.zhihu.com/p/358475775](https://zhuanlan.zhihu.com/p/358475775)


```go
    go test -v -bench=. pool_test.go
    go test -v -bench=. -benchtime=5s pool_test.go // 基准测试框架对一个测试用例的默认测试时间是 1 秒
    go test -v -bench=Alloc -benchmem pool_test.go //显示内存分配情况
```
### sync.Pool 本身是线程安全的
### 1.禁止拷贝
### 2.不能存放需要保持长连接的对象
