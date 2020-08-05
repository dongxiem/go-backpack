package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// 使用GET方法进行Request请求
	req, err := http.NewRequest("GET", "https://www.baidu.com", nil)
	if err != nil {
		log.Fatalf("could not create request: %v", err)
	}
	// 使用http创建一个client
	client := http.DefaultClient
	// 执行Request请求，并得到response
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("http request failed: %v", err)
	}
	// 将response的状态进行打印
	fmt.Println(res.Status)
}
