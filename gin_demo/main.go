package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sayhello(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadFile("./hello.txt")

	fmt.Fprintln(w, string(b))

}

func main() {

	http.HandleFunc("/hello", sayhello)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println("http server failed, err:", err)
		return
	}

}
