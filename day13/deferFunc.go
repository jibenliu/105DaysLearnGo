package main

import "fmt"

func hello(i int) {
	fmt.Println(i)
}

func main() {
	i := 5
	defer hello(i)
	i = i + 10
	defer hello(i)
	i = i + 10
}

// 参考答案及解析：15 5。这个例子中，hello() 函数的参数在执行 defer 语句的时候会保存一份副本，在实际调用 hello() 函数时用，所以是 5.
