package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	//1. ClientHello
	//clientHello()

	//2. ServerHello
	//serverHello()

	//3.Certificate
	//certificate()

	//4.Server Exchange Key
	//serverKeyExchange()

	//5. Server Hello Done
	//serverHelloDone()

	//6.clientKeyExchange
	clientKeyExchange()
}

func clientHello() {

	//flows, err := data.LoadData("../go_study/tls/data/ClientHello.txt")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//for _, f := range flows {
	//	var m clientHelloMsg
	//	m.unmarshal(f)
	//	fmt.Println(m)
	//	fmt.Println(m.serverName)
	//}

	f, _ := os.Open("../go_study/tls/data/clienthelo.bin")
	flows, _ := ioutil.ReadAll(f)
	var m clientHelloMsg
	m.unmarshal(flows[5:])
	fmt.Println(m.random)
	fmt.Println(flows)
	fmt.Println(m.serverName)
}

func serverHello() {
	f, _ := os.Open("../go_study/tls/data/serverhello.bin")
	flows, _ := ioutil.ReadAll(f)
	var m serverHelloMsg
	m.unmarshal(flows[5:])
	fmt.Println(m.random)
	fmt.Println(flows)
	fmt.Println(m.vers)
}

func certificate() {
	f, _ := os.Open("../go_study/tls/data/certificate.bin")
	flows, _ := ioutil.ReadAll(f)
	var m certificateMsg
	m.unmarshal(flows[5:])
	fmt.Println(m)
	fmt.Println(len(m.certificates))
}

func serverKeyExchange() {
	f, _ := os.Open("../go_study/tls/data/serverKeyExchange.bin")
	flows, _ := ioutil.ReadAll(f)
	var m serverKeyExchangeMsg
	m.unmarshal(flows[5:])
	fmt.Println(m)
	fmt.Println(m.key)
}

func serverHelloDone() {
	f, _ := os.Open("../go_study/tls/data/serverHelloDone.bin")
	flows, _ := ioutil.ReadAll(f)
	var m serverHelloDoneMsg
	fmt.Println(m.unmarshal(flows[5:]))
}

func clientKeyExchange() {
	f, _ := os.Open("../go_study/tls/data/clientKeyExchange.bin")
	flows, _ := ioutil.ReadAll(f)
	var m clientKeyExchangeMsg
	m.unmarshal(flows[5:])
	fmt.Println(m)
	fmt.Println(m.ciphertext)
}
