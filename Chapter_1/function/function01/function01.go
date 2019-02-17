package main

import (
	"fmt"
)

func A(s ...int) {

	s[0] = 3
	s[1] = 4
	fmt.Println(s)

}

func main() {
	//A(1, 2, 3, 5, 6, 7)

	a, b := 1, 2
	A(a, b)
	fmt.Println(a, b)

}
