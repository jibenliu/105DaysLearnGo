package main

import "fmt"

func hello(num ...int) {
	num[0] = 18
}

func hello1(num []int) {
	num[0] = 5
}

func main() {
	i := []int{5, 6, 7}
	hello(i...)
	fmt.Println(i[0]) //18
	hello1(i)
	fmt.Println(i[0]) //5
}
