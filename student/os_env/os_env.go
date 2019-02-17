package main

import "fmt"
import "os"

func main() {
	var JAVAHOME string
	JAVAHOME = os.Getenv("GOROOT")
	fmt.Println(GOROOT)
}
