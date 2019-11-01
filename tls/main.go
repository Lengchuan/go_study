package main

import (
	"crypto/rand"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
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

	//6.Client Key Exchange
	//clientKeyExchange()

	//7. Change Cipher Spec
	//changeCipherSpec()

	//8. Encrypted Handshake Message
	//encryptedHandshakeMessage()

	//9. New Session Ticket
	//newSessionTicket()

	//10. Application Data
	//applicationData()

	Decode()
}

func clientHello() (m clientHelloMsg) {

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
	m.unmarshal(flows[5:])
	fmt.Println(m.random)
	fmt.Println(flows)
	fmt.Println(m.serverName)

	return
}

func serverHello() (m serverHelloMsg) {
	f, _ := os.Open("../go_study/tls/data/serverhello.bin")
	flows, _ := ioutil.ReadAll(f)
	m.unmarshal(flows[5:])
	fmt.Println(m.random)
	fmt.Println(flows)
	fmt.Println(m.vers)

	return
}

func certificate() (m certificateMsg) {
	f, _ := os.Open("../go_study/tls/data/certificate.bin")
	flows, _ := ioutil.ReadAll(f)
	m.unmarshal(flows[5:])
	fmt.Println(m)
	fmt.Println(len(m.certificates))

	return
}

func serverKeyExchange() (m serverKeyExchangeMsg) {
	f, _ := os.Open("../go_study/tls/data/serverKeyExchange.bin")
	flows, _ := ioutil.ReadAll(f)
	m.unmarshal(flows[5:])
	fmt.Println(m)
	fmt.Println(m.key)
	fmt.Println(m.key)

	return
}

func serverHelloDone() (m serverHelloDoneMsg) {
	f, _ := os.Open("../go_study/tls/data/serverHelloDone.bin")
	flows, _ := ioutil.ReadAll(f)
	fmt.Println(m.unmarshal(flows[5:]))

	return
}

func clientKeyExchange() (m clientKeyExchangeMsg) {
	f, _ := os.Open("../go_study/tls/data/clientKeyExchange.bin")
	flows, _ := ioutil.ReadAll(f)
	m.unmarshal(flows[5:])
	fmt.Println(m)
	fmt.Println(m.ciphertext)

	return
}

func changeCipherSpec() {
	//
}

func encryptedHandshakeMessage() {
	//
}

func newSessionTicket() {
	f, _ := os.Open("../go_study/tls/data/newSessionTicket.bin")
	flows, _ := ioutil.ReadAll(f)
	var m newSessionTicketMsg
	m.unmarshal(flows[5:])
	fmt.Println(m)
	fmt.Println(m.ticket)
}

func applicationData() []byte {
	f, _ := os.Open("../go_study/tls/data/applicationData.bin")
	flows, _ := ioutil.ReadAll(f)
	fmt.Println(flows[5:])
	fmt.Println(string(flows[5:]))
	return flows
}

func Decode() {
	clientHello := clientHello()
	serverHello := serverHello()
	//certificate := certificate()
	serverKeyExchange := serverKeyExchange()
	clientKeyExchange := clientKeyExchange()
	//applicationData := applicationData()

	serverCertificate, err := tls.LoadX509KeyPair("../go_study/tls/data/1.pem", "../go_study/tls/data/1.pem")
	if err != nil {
		log.Fatal(err)
	}

	//tls 版本
	version := serverHello.vers
	//随机数
	clientRandom := clientHello.random
	serverRandom := serverHello.random
	//使用的密码套件
	cipherSuiteid := serverHello.cipherSuite
	cipherSuite := cipherSuiteByID(cipherSuiteid)
	//服务器私钥
	//privateKey := serverCertificate.PrivateKey

	//ECDH 算法 参数
	keyAgreement := cipherSuite.ka(version)
	curveType := serverKeyExchange.raw[4:5]
	curveName := serverKeyExchange.raw[5:7]
	pub := serverKeyExchange.key

	fmt.Println("curveType ", curveType)
	fmt.Println("curveName ", curveName)
	fmt.Println("pub ", pub)

	ecdheParameters, err := generateECDHEParameters(rand.Reader, CurveID(curveName[1]))
	ka, ok := keyAgreement.(*ecdheKeyAgreement)
	if ok {
		ka.params = ecdheParameters
		keyAgreement = ka
	}

	//解析pre Master Secret 预主秘钥
	//1. ECDH 算法
	//2. RSA  算法

	config := &Config{Rand: rand.Reader}
	cer := Certificate{
		Certificate:                 serverCertificate.Certificate,
		PrivateKey:                  serverCertificate.PrivateKey,
		OCSPStaple:                  serverCertificate.OCSPStaple,
		SignedCertificateTimestamps: serverCertificate.SignedCertificateTimestamps,
		Leaf:                        serverCertificate.Leaf,
	}
	preMasterSecret, err := keyAgreement.processClientKeyExchange(config, &cer, &clientKeyExchange, version)
	if err != nil {
		log.Fatal(err)
	}

	masterSecret := masterFromPreMasterSecret(version, cipherSuite, preMasterSecret, clientRandom, serverRandom)

	fmt.Println(masterSecret)
	fmt.Println(preMasterSecret)
	fmt.Println(len(masterSecret))

}
