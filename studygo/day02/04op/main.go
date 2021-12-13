package main

import "fmt"

func main() {
	var (
		a =5
		b =2
	)

	//算术运算符
	fmt.Println(a+b)
	fmt.Println(a-b)
	fmt.Println(a*b)
	fmt.Println(a/b)
	fmt.Println(a%b)
	a++ //单独的语句，不能放在=的右边复制 ==> a = a + 1
	b-- //单独的语句，不能放在=的右边复制 ==> b = b - 1
	fmt.Println(a)
	fmt.Println(b)
	//关系运算符
	

		
	}