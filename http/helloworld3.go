package main

import (
	"net/http"
	"time"
	"log"
	"io"
)

func main() {
	server := http.Server{
		Addr:        ":8080",
		Handler:     &myhandler3{},
		ReadTimeout: 6 * time.Second,
	}
	mux = make(map[string]func(w http.ResponseWriter, r *http.Request))
	mux["/hello"] = sayHello3

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}
}

var mux map[string]func(http.ResponseWriter, *http.Request)

type myhandler3 struct {
}

func (*myhandler3) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := mux[r.URL.String()]; ok {
		h(w, r)
		return
	}
	io.WriteString(w, "URL :"+r.URL.String())
}

func sayHello3(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world 3")
}
