package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func helloHandler3(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, web 3")
}

func main() {
	server := &http.Server{
		Addr:         ":4000",
		WriteTimeout: 2 * time.Second,
	}

	mux := http.NewServeMux()

	mux.Handle("/", &myHandle3{})
	mux.HandleFunc("/hello", helloHandler3)
	server.Handler = mux

	log.Println("Start server and listen on 4000")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

type myHandle3 struct{}

func (*myHandle3) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome Server 3: "+r.URL.String())
}
