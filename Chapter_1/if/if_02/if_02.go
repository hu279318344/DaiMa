package main

import (
	"fmt"
)

func for01() {
	a := 0
	for {
		a++
		if a >= 10 {
			break
		}
		fmt.Println(a)
	}
	fmt.Print("For循环结束")

}

func switch01() {
	a := 1
	switch {
	case a >= 0:
		fmt.Println("a=0")
		fallthrough
	case a >= 1:
		fmt.Println("a=1")
	default:
		fmt.Println("None")
	}
}

func break01() {
LABLE1:
	for {
		for i := 0; i < 10; i++ {
			if i > 3 {
				break LABLE1
			}
		}
	}
	fmt.Println("OK")
}

func break02() {

	for {
		for i := 0; i < 10; i++ {
			if i > 3 {
				goto LABLE1
			}
		}
	}

LABLE1:
	fmt.Println("OK-lable1")
}

func continue01() {
LABLE2:
	for i := 0; i < 10; i++ {
		for {
			fmt.Printf("%d\n", i)
			continue LABLE2

		}

	}
	fmt.Println("countinue-ok-LABEL2")

}

func main() {
	//for01()
	//switch01()
	//break01()
	//break02()
	continue01()
}
