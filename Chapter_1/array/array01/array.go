package main

import (
	"fmt"
)

func array01() {
	//var a [2]int
	//var b [1]int
	a := [20]int{1, 2: 6}

	fmt.Println(a)

}

func array02() {

	a := [...]int{99: 1}
	//var p *[100]int = &a
	var p = &a
	fmt.Println(p)

}

func array03() {

	x, y := 1, 2
	a := [...]*int{&x, &y}
	fmt.Println(a)

}

func array04() {
	a := [3][3]int{
		{1, 1, 1},
		{2, 1, 2}}
	fmt.Println(a)
}

//go冒泡排序

func array05() {
	a := [...]int{5, 4, 6, 7, 9, 1, 100, 50, 60, 12, 30}
	fmt.Println(a)
	num := len(a)
	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ {
			if a[i] < a[j] {
				temp := a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}
	fmt.Println(a)
}

func main() {
	//array01()
	//array02()
	//array03()
	//array04()
	array05()
}
