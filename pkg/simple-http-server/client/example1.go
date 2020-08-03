package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// try changing the value of this url
	res, err := http.Get("https://www.baidu.com")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.Status)
}
