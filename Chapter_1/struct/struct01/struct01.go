package main

import (
	"fmt"
)

type person struct {
	Name string
	Age  int
}

//地址拷贝
func main() {
	a := &person{
		Name: "join",
		Age:  19,
	}
	//a.Name = "join"
	//a.Age = 19
	fmt.Println(a)

	A(a)
	B(a)
	fmt.Println(a)

}

func A(per *person) {
	per.Name = "boy"
	per.Age = 13
	fmt.Println("A", per)
}

func B(per *person) {
	per.Name = "gale"
	per.Age = 16
	fmt.Println("A", per)
}
