package main

import (
	"crypto/tls"
	"log"
	"net"
	"bufio"
)

func main() {

	//加载服务器证书
	cert, err := tls.LoadX509KeyPair("../go_study/tls/server/cert.pem", "../go_study//tls/server/key.pem");
	if err != nil {
		log.Println(err)
		return
	}

	config := &tls.Config{Certificates: []tls.Certificate{cert}}
	ln, err := tls.Listen("tcp", ":443", config)
	if err != nil {
		log.Println(err)
		return
	}
	defer ln.Close()
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}

}

func handleConn(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	for {
		msg, err := r.ReadString('\n')
		if err != nil {
			log.Println(err)
			return
		}
		println(msg)
		n, err := conn.Write([]byte("world\n"))
		if err != nil {
			log.Println(n, err)
			return
		}
	}
}

//生成秘钥
//go run $GOROOT/src/crypto/tls/generate_cert.go --host localhost
