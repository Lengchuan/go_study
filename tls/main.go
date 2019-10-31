package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

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

	f, _ := os.Open("../go_study/tls/data/1.bin")
	flows, _ := ioutil.ReadAll(f)
	var m clientHelloMsg
	m.unmarshal(flows)
	fmt.Println(m.random)
	fmt.Println(flows)
	fmt.Println(m.serverName)

}
