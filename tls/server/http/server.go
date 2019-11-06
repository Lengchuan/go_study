package main

import (
	"crypto/tls"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("hello word\n"))
}

//GODEBUG=http2server=0
func main() {

	http.HandleFunc("/hello", HelloServer)

	config := &tls.Config{
		CipherSuites: []uint16{tls.TLS_RSA_WITH_AES_256_GCM_SHA384},
	}

	s := &http.Server{
		TLSConfig: config,
		Addr:      ":8443",
	}

	err := s.ListenAndServeTLS("../go_study/tls/server/http/www.lengchuan.study_chain.crt", "../go_study/tls/server/http/www.lengchuan.study_key.key")
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
