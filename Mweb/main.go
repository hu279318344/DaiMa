package main

//获取http状态码

import (
	"fmt"
	"net/http"
	"net/url"
)

func main() {
	u, _ := url.Parse("https://www.baidu1.com")
	q := u.Query()
	u.RawQuery = q.Encode()
	res, err := http.Get(u.String())
	if err != nil {
		fmt.Println("0")
		return
	}
	resCode := res.StatusCode
	res.Body.Close()
	if err != nil {
		fmt.Println("0")
		return
	}
	fmt.Printf("%d\r\n", resCode)

}
