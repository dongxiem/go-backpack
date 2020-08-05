package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// 直接使用Get方法进行请求
	res, err := http.Get("https://www.baidu.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.Status)
}
