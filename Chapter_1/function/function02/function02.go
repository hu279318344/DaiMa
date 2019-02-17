package main

import (
	"fmt"
)

// A copy
func A(a *int) {
	*a = 2
	fmt.Println(*a)
	fmt.Println(&a)
}

func main() {
	a := 1
	A(&a)
	fmt.Println(a)
	fmt.Println(&a)
}
