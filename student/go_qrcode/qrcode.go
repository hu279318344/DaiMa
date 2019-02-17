package main

import (
	"fmt"

	"github.com/skip2/go-qrcode"
)

func main() {
	err := qrcode.WriteFile("http://www.baidu.com", qrcode.Medium, 1024, "qr.png")
	if err != nil {
		fmt.Println("write error")

	}
}
