package main

import "fmt"

func main() {
	//list := new([]int)
	//list = append(list,1)
	//上述代码编译不通过，new创建的是一个*[]int类型的指针

	list := make([]int, 0)
	list = append(list, 1)
	fmt.Println(list)

	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	//s1 = append(s1, s2)
	// 上述代码不能通过编译 需使用 … 操作符，将一个切片追加到另一个切片上
	s1 = append(s1, s2...)
	fmt.Println(s1)
}

/** --------------------------------------------------华丽的分割线--------------------------------------------**/
//var(
//	size := 1024
//	max_size = size*2
//)
//func main() {
//	fmt.Println(size,max_size)
//}
//  上述代码不能通过编译 变量声明的简短模式有如下限制
//1.必须使用显示初始化；
//2.不能提供数据类型，编译器会自动推导；
//3.只能在函数内部使用简短模式；