package main

import "fmt"

func main() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)

	s1 := make([]int, 0)
	s1 = append(s1, 1, 2, 3, 4)
	fmt.Println(s1)
}


// 在函数有多个返回值时，只要有一个返回值有命名，其他的也必须命名。如果有多个返回值必须加上括号()；如果只有一个返回值且命名也必须加上括号()

// new(T) 和 make(T,args) 是 Go 语言内建函数，用来分配内存，但适用的类型不同。都是空切片
//new(T) 会为 T 类型的新值分配已置零的内存空间，并返回地址（指针），即类型为 *T的值。换句话说就是，返回一个指针，该指针指向新分配的、类型为 T 的零值。适用于值类型，如数组、结构体等。
//make(T,args) 返回初始化之后的 T 类型的值，这个值并不是 T 类型的零值，也不是指针 *T，是经过初始化之后的 T 的引用。make() 只适用于 slice、map 和 channel.

// var a []int nil切片
// var a []int{} 空切片


//  最好不要创建空切片，统一使用nil切片

/**
	type something struct{
		Values []int
	}
	var s1 = something{}
	var s2 = something{[]int}
	bs1,_:=json.Marshal(s1)
	// {"values":null}
	bs2,_:=json.Marshal(s2)
	// {"values":[]}
 */