package main

import (
	"fmt"
	"os"
)

var JAVAHOME string

func main() {

	JAVAHOME = os.Getenv("GOROOT")
	fmt.Println(JAVAHOME)
}
