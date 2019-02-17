package main

import "fmt"

const (
	B float64 = 1 << (iota * 10) //存储单位值
	KB
	MB
	GB
)

func main() {

	fmt.Println(B)
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
}
