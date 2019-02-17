package main

import (
	"fmt"
)

//匿名方式
func struct02() {
	a := struct {
		Name string
		Age  int
	}{
		Name: "galen",
		Age:  19,
	}
	fmt.Println(a)
}

func main() {
	struct02()
}
