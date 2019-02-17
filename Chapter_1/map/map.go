package main

import (
	"fmt"
	"sort"
)

func map01() {
	//var m map[int]string
	//m = make(map[int]string)
	m := make(map[int]string)
	m[1] = "OK"
	//delete(m,1)
	a := m[1]
	fmt.Println(a)
}

//嵌套map
func map02() {
	var m map[int]map[int]string
	m = make(map[int]map[int]string)
	a, ok := m[2][1]
	if !ok {
		m[2] = make(map[int]string)
	}
	m[2][1] = "Good"
	a, ok = m[2][1]
	fmt.Println(a, ok)

}

//不修改slice的值
func map03() {
	sm := make([]map[int]string, 5)
	for _, v := range sm {
		v = make(map[int]string, 1)
		v[1] = "ok"
		fmt.Println(v)
	}
	fmt.Println(sm)
}

//修改slice的值
func map04() {
	sm := make([]map[int]string, 5)
	for i := range sm {
		sm[i] = make(map[int]string, 1)
		sm[i][1] = "OK"
		fmt.Println(sm[i])
	}
	fmt.Println(sm)
}

//map无序性
func map05() {
	m := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
	s := make([]int, len(m))
	i := 0
	for k := range m {
		s[i] = k
		i++
	}
	fmt.Println(s)
}

//map间接排序
func map06() {
	m := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
	s := make([]int, len(m))
	i := 0
	for k := range m {
		s[i] = k
		i++
	}
	sort.Ints(s)
	fmt.Println(s)

}

func map07() {
	m1 := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e", 6: "f"}
	fmt.Println(m1)
	m2 := make(map[string]int)
	for k, v := range m1 {
		m2[v] = k
	}
	fmt.Println(m2)

}

func main() {
	//map01()
	//map02()
	//map03()
	//map04()
	//map05()
	//map06()
	map07()
}
