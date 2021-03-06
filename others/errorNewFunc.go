package main

import (
	"errors"
	"fmt"
)

// 自定义的错误类型
type DefineError struct {
	msg string
}

func (d *DefineError) Error() string {
	return d.msg
}

var (
	ErrInvalidUser     = errors.New("invalid user")
	ErrNotFoundUser    = errors.New("not found user")
)

func main()  {
	// wrap error
	err1 := &DefineError{"this is a define error type"}
	err2 := fmt.Errorf("wrap err2: %w\n", err1)
	err3 := fmt.Errorf("wrap err3: %w\n", err2)
	var err4 *DefineError
	if errors.As(err3, &err4) {
		// errors.As() 顺着错误链，从 err3 一直找到被包装最底层的错误值 err1，并且将 err3 与其自定义类型 `var err4 *DefineError` 匹配成功。
		fmt.Println("err1 is a variable of the DefineError type")
		fmt.Println(err4 == err1)
		return
	}
	fmt.Println("err1 is not a variable of the DefineError type")

	err5 := fmt.Errorf("wrap err5: %w\n", ErrInvalidUser)
	err6 := fmt.Errorf("wrap err6: %w\n", err5)
	// golang 1.13 新增 Is() 函数
	if errors.Is(err6, ErrInvalidUser) {
		fmt.Println(ErrInvalidUser)
		return
	}
	fmt.Println("success")
}
