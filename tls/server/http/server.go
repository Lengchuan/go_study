package main

import (
	"net/http"
	"log"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("hello word\n"))
}

func main() {
	http.HandleFunc("/hello", HelloServer)
	err := http.ListenAndServeTLS(":443", "../go_study/tls/server/http/www.lengchuan.study_chain.crt", "../go_study/tls/server/http/www.lengchuan.study_key.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
