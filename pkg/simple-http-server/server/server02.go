package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, web 2")
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", &myHandle2{})
	mux.HandleFunc("/hello", helloHandler2)

	log.Println("Start server and listen on 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

type myHandle2 struct{}

func (*myHandle2) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome Server 2: "+r.URL.String())
}
