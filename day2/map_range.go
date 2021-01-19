package main

import "fmt"

func main() {
	sli := []int{0, 1, 2, 3}
	m := make(map[int]*int)
	for key, val := range sli {
		m[key] = &val
	}
	for k, v := range m {
		fmt.Println(k, "->", *v)
	}

	m1 := make(map[int]*int)
	for key, val := range sli {
		value := val
		m1[key] = &value
	}

	for k, v := range m1 {
		fmt.Println(k, "->", *v)
	}
}
