package main

import (
	"fmt"
)

func main() {
	a := 1
	a++
	var p = &a
	fmt.Println(*p)

}
