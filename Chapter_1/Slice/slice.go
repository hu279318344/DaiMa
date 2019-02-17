package main

import (
	"fmt"
)

func slice01() {
	var s1 []int
	fmt.Println(s1)
}

func slice02() {
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(a)
	s1 := a[5:10] //a[5 6 7 8 9]
	//s1:=a[5:len(a)]
	//s1:=a[5:]//end 5
	fmt.Println(s1)

}

func slice03() {
	s1 := make([]int, 3, 10) //3个元素。容量为10
	fmt.Println(len(s1), cap(s1))

}

func slice04() {
	a := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k'}
	sa := a[2:5]
	fmt.Println(string(a))
	fmt.Println(string(sa))
}

//Append
func slice05() {
	s1 := make([]int, 3, 6)
	fmt.Printf("%p", s1)
	s1 = append(s1, 1, 2, 3)
	fmt.Printf("%v %p\n", s1, s1)
	s1 = append(s1, 1, 2, 3)
	fmt.Printf("%v %p\n", s1, s1)
	s1 = append(s1, 4, 5, 6)
	fmt.Printf("%v %p\n", s1, s1)
}

//Append02
func slice06() {
	a := []int{1, 2, 3, 4, 5}
	s1 := a[2:5]
	s2 := a[1:3]
	fmt.Println(s1, s2)
	s2 = append(s2, 1, 2, 5, 6, 7, 8, 1)
	s1[0] = 9
	fmt.Println(s1, s2)

}

//copy
func slice07() {
	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := []int{7, 8, 9}
	copy(s1, s2)
	fmt.Println(s1)

}

func main() {

	//slice01()
	//slice02()
	//slice03()
	//slice04()
	//slice05()
	//slice06()
	slice07()
}
