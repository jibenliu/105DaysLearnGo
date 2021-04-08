# refer: [https://zhuanlan.zhihu.com/p/358475775](https://zhuanlan.zhihu.com/p/358475775)


```go
    go test -v -bench=. pool_test.go
```
### sync.Pool 本身是线程安全的
### 1.禁止拷贝
### 2.不能存放需要保持长连接的对象
