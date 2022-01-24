package main

import "fmt"

func main() {
	var (
		a = 5
		b = 2
	)

	//算术运算符
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	a++ //单独的语句，不能放在=的右边复制 ==> a = a + 1
	b-- //单独的语句，不能放在=的右边复制 ==> b = b - 1
	fmt.Println(a)
	fmt.Println(b)

	//关系运算符
	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a >= b)
	fmt.Println(a <= b)
	fmt.Println(a == b)
	fmt.Println(a < b)

	//逻辑运算符
	//如果年龄大于18岁并且年龄小于60岁
	age := 22
	if age > 18 && age < 60 {
		fmt.Println("苦逼上班的")
	} else {
		fmt.Println("幸福的")
	}

	//如果年龄小于18岁或者年龄大于60岁
	if age < 18 || age < 60 {
		fmt.Println("不用上班的")
	} else {
		fmt.Println("苦逼上班的")
	}

	//not取反，原来为真就为假，原来假就为真

	isMarried := false

	fmt.Println(isMarried)  //false
	fmt.Println(!isMarried) //true
}
