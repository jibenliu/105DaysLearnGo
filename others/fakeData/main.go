package main

import (
	"fmt"
	"github.com/jaswdr/faker"
)

type ExampleStruct struct {
	SimpleStringField  string
	SimpleNumber       int
	SimpleBool         bool
	SomeFormatedString string    `fake:"??? ###"`
	SomeStringArray    [5]string `fake:"????"`
}

func main() {
	f := faker.New()
	fmt.Println(f.Person().Name())
	fmt.Println(f.Address().Address())
	fmt.Println(f.Phone().Number())
	fmt.Println(f.Lorem().Text(20))

	example := ExampleStruct{}
	f.Struct().Fill(&example)
	fmt.Printf("%+v", example)
}
