package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/subtle"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sync"
)

type halfConn struct {
	sync.Mutex

	err            error       // first permanent error
	version        uint16      // protocol version
	cipher         interface{} // cipher algorithm
	mac            macFunction
	seq            []byte   // 64-bit sequence number
	additionalData [13]byte // to avoid allocs; interface method args escape

	nextCipher interface{} // next encryption state
	nextMac    macFunction // next MAC algorithm

	trafficSecret []byte // current TLS 1.3 traffic secret
}

type cbcMode interface {
	cipher.BlockMode
	SetIV([]byte)
}

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

	//finish()

	//var b cryptobyte.Builder
	//b.AddUint8(typeFinished)
	//ba :=b.BytesOrPanic()
	//fmt.Println(ba)
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

func finish() {
	f, _ := os.Open("../go_study/tls/data/1.bin")
	flows, _ := ioutil.ReadAll(f)
	//var m finishedMsg
	//b :=m.unmarshal(flows[5:])
	//b :=m.unmarshal(flows[:])
	//fmt.Println(b)
	//fmt.Println(m.raw)
	//fmt.Println(m.verifyData)
	fmt.Println(flows)

	return
}

func applicationData(req bool) []byte {
	filename := "../go_study/tls/data/req.bin"
	if !req {
		filename = "../go_study/tls/data/resp.bin"
	}
	f, _ := os.Open(filename)
	flows, _ := ioutil.ReadAll(f)
	return flows
}

func decryp(data []byte, key []byte) {
	// Load your secret key from a safe place and reuse it across multiple
	// NewCipher calls. (Obviously don't use this example key for anything
	// real.) If you want to convert a passphrase to a key, use a suitable
	// package like bcrypt or scrypt.
	//key, _ := hex.DecodeString("6368616e676520746869732070617373")

	//bReader := bytes.NewReader([]byte("some secret text"))
	bReader := bytes.NewReader(data)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// If the key is unique for each ciphertext, then it's ok to use a zero
	// IV.
	var iv [aes.BlockSize]byte
	stream := cipher.NewOFB(block, iv[:])

	var out bytes.Buffer

	writer := &cipher.StreamWriter{S: stream, W: &out}
	// Copy the input to the output buffer, encrypting as we go.
	if _, err := io.Copy(writer, bReader); err != nil {
		panic(err)
	}

	// Note that this example is simplistic in that it omits any
	// authentication of the encrypted data. If you were actually to use
	// StreamReader in this manner, an attacker could flip arbitrary bits in
	// the decrypted result.

	fmt.Printf("%x\n", out.Bytes())
	// Output: cf0495cc6f75dafc23948538e79904a9
}

func Decode() {
	clientHello := clientHello()
	serverHello := serverHello()
	certificate := certificate()
	//serverKeyExchange := serverKeyExchange()
	clientKeyExchange := clientKeyExchange()

	isReq := false
	applicationData := applicationData(isReq)

	serverCertificate, err := tls.LoadX509KeyPair("../go_study/tls/server/http/www.lengchuan.study_chain.crt", "../go_study/tls/server/http/www.lengchuan.study_key.key")
	if err != nil {
		log.Fatal(err)
	}

	//tls 版本
	version := serverHello.vers
	fmt.Println("version ", version)
	//随机数
	clientRandom := clientHello.random
	serverRandom := serverHello.random
	//使用的密码套件
	cipherSuiteid := serverHello.cipherSuite
	//cipherSuiteid := TLS_RSA_WITH_AES_256_GCM_SHA384
	cipherSuite := cipherSuiteByID(cipherSuiteid)
	//服务器私钥
	//privateKey := serverCertificate.PrivateKey

	//ECDH 算法 参数
	keyAgreement := cipherSuite.ka(version)
	//curveType := serverKeyExchange.raw[4:5]
	//curveName := serverKeyExchange.raw[5:7]
	//pub := serverKeyExchange.key

	//fmt.Println("curveType ", curveType)
	//fmt.Println("curveName ", curveName)
	//fmt.Println("pub ", pub)

	//ecdheParameters, err := generateECDHEParameters(rand.Reader, CurveID(curveName[1]))
	//ka, ok := keyAgreement.(*ecdheKeyAgreement)
	//if ok {
	//	ka.params = ecdheParameters
	//	keyAgreement = ka
	//}
	//bs := make([]byte, 8)
	//binary.LittleEndian.PutUint64(bs, 803)
	//fmt.Println("bs :  ", bs)

	//解析pre Master Secret 预主秘钥
	//1. ECDH 算法
	//2. RSA  算法

	config := &Config{}
	cer := Certificate{
		Certificate:                 certificate.certificates,
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

	fmt.Println("version: ", version)
	fmt.Println("client random: ", clientRandom)
	fmt.Println("len client random: ", len(clientRandom))
	fmt.Println("server random: ", serverRandom)
	fmt.Println("preMasterSecret: ", preMasterSecret)
	fmt.Println("len(preMasterSecret): ", len(preMasterSecret))
	fmt.Println("masterSecret: ", masterSecret)
	fmt.Println("len(masterSecret): ", len(masterSecret))
	fmt.Println("len(applicationData): ", len(applicationData))
	fmt.Println("applicationData: ", applicationData)

	clientMAC, serverMAC, clientKey, serverKey, clientIV, serverIV :=
		//clientMAC, _, clientKey, _, clientIV, _ :=
		//_, serverMAC, _, serverKey, _, serverIV :=
		keysFromMasterSecret(version, cipherSuite, masterSecret, clientRandom, serverRandom, cipherSuite.macLen, cipherSuite.keyLen, cipherSuite.ivLen)

	var serverCipher interface{}
	var clientCipher interface{}
	var serverHash macFunction
	var clientHash macFunction

	if cipherSuite.aead == nil {
		clientCipher = cipherSuite.cipher(clientKey, clientIV, isReq /* for reading */)
		clientHash = cipherSuite.mac(version, clientMAC)
		serverCipher = cipherSuite.cipher(serverKey, serverIV, !isReq /* not for reading */)
		serverHash = cipherSuite.mac(version, serverMAC)
	} else {
		clientCipher = cipherSuite.aead(clientKey, clientIV)
		serverCipher = cipherSuite.aead(serverKey, serverIV)
	}

	var hc = &halfConn{
		version: version,
		//nextCipher: clientCipher,
		//nextMac:    clientHash,
		//seq:        bs,
	}
	if isReq {
		hc.cipher = clientCipher
		hc.mac = clientHash
	} else {
		hc.cipher = serverCipher
		hc.mac = serverHash
	}

	fmt.Println("ClientMAC: ", clientMAC)
	fmt.Println("ServerMAC: ", serverMAC)
	fmt.Println("ClientKey: ", clientKey)
	fmt.Println("ServerKey: ", serverKey)
	fmt.Println("ClientIV: ", clientIV)
	fmt.Println("ServerIV: ", serverIV)

	plaintext, _, err := hc.decrypt(applicationData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("plaintext:\r\n", string(plaintext))

}

// decrypt authenticates and decrypts the record if protection is active at
// this stage. The returned plaintext might overlap with the input.
func (hc *halfConn) decrypt(record []byte) ([]byte, recordType, error) {
	var plaintext []byte
	typ := recordType(record[0])
	payload := record[recordHeaderLen:]

	// In TLS 1.3, change_cipher_spec messages are to be ignored without being
	// decrypted. See RFC 8446, Appendix D.4.
	if hc.version == VersionTLS13 && typ == recordTypeChangeCipherSpec {
		return payload, typ, nil
	}

	paddingGood := byte(255)
	paddingLen := 0

	explicitNonceLen := hc.explicitNonceLen()

	if hc.cipher != nil {
		switch c := hc.cipher.(type) {
		case cipher.Stream:
			c.XORKeyStream(payload, payload)
		case aead:
			if len(payload) < explicitNonceLen {
				//return nil, 0, tls.alertBadRecordMAC
				return nil, 0, nil
			}
			nonce := payload[:explicitNonceLen]
			//if len(nonce) == 0 {
			//	nonce = hc.seq[:]
			//}
			payload = payload[explicitNonceLen:]

			additionalData := hc.additionalData[:]
			if hc.version == VersionTLS13 {
				additionalData = record[:recordHeaderLen]
			} else {
				//copy(additionalData, hc.seq[:])
				copy(additionalData, nonce)
				copy(additionalData[8:], record[:3])
				n := len(payload) - c.Overhead()
				additionalData[11] = byte(n >> 8)
				additionalData[12] = byte(n)
			}

			fmt.Println("nonce: ", nonce)
			fmt.Println("additionalData: ", additionalData)
			fmt.Println("payload: ", payload)
			fmt.Println("len(payload): ", len(payload))
			fmt.Println("c.Overhead(): ", c.Overhead())
			fmt.Println(fmt.Sprintf("cipher %v:", c))
			var err error
			plaintext, err = c.Open(payload[:0], nonce, payload, additionalData)
			if err != nil {
				//return nil, 0, alertBadRecordMAC
				return nil, 0, err
			}

			fmt.Println("plaintext ", plaintext)
			fmt.Println("plaintext ", string(plaintext))
		case cbcMode:
			blockSize := c.BlockSize()
			minPayload := explicitNonceLen + roundUp(hc.mac.Size()+1, blockSize)
			if len(payload)%blockSize != 0 || len(payload) < minPayload {
				//return nil, 0, alertBadRecordMAC
				return nil, 0, nil
			}

			if explicitNonceLen > 0 {
				c.SetIV(payload[:explicitNonceLen])
				payload = payload[explicitNonceLen:]
			}
			c.CryptBlocks(payload, payload)

			// In a limited attempt to protect against CBC padding oracles like
			// Lucky13, the data past paddingLen (which is secret) is passed to
			// the MAC function as extra data, to be fed into the HMAC after
			// computing the digest. This makes the MAC roughly constant time as
			// long as the digest computation is constant time and does not
			// affect the subsequent write, modulo cache effects.
			if hc.version == VersionSSL30 {
				paddingLen, paddingGood = extractPaddingSSL30(payload)
			} else {
				paddingLen, paddingGood = extractPadding(payload)
			}
		default:
			panic("unknown cipher type")
		}

		if hc.version == VersionTLS13 {
			if typ != recordTypeApplicationData {
				//return nil, 0, alertUnexpectedMessage
				return nil, 0, nil
			}
			if len(plaintext) > maxPlaintext+1 {
				//return nil, 0, alertRecordOverflow
				return nil, 0, nil
			}
			// Remove padding and find the ContentType scanning from the end.
			for i := len(plaintext) - 1; i >= 0; i-- {
				if plaintext[i] != 0 {
					typ = recordType(plaintext[i])
					plaintext = plaintext[:i]
					break
				}
				if i == 0 {
					//return nil, 0, alertUnexpectedMessage
					return nil, 0, nil
				}
			}
		}
	} else {
		plaintext = payload
	}

	if hc.mac != nil {
		macSize := hc.mac.Size()
		if len(payload) < macSize {
			//return nil, 0, alertBadRecordMAC
			return nil, 0, nil
		}

		n := len(payload) - macSize - paddingLen
		n = subtle.ConstantTimeSelect(int(uint32(n)>>31), 0, n) // if n < 0 { n = 0 }
		//record[3] = byte(n >> 8)
		//record[4] = byte(n)
		//remoteMAC := payload[n : n+macSize]
		//localMAC := hc.mac.MAC(hc.seq[0:], record[:recordHeaderLen], payload[:n], payload[n+macSize:])

		fmt.Println(paddingGood)
		//if subtle.ConstantTimeCompare(localMAC, remoteMAC) != 1 || paddingGood != 255 {
		//	//return nil, 0, alertBadRecordMAC
		//	return nil, 0, nil
		//}

		plaintext = payload[:n]
	}

	//hc.incSeq()
	return plaintext, typ, nil
}

// decrypt authenticates and decrypts the record if protection is active at
// this stage. The returned plaintext might overlap with the input.
func (hc *halfConn) decrypt1(record []byte) ([]byte, recordType, error) {
	var plaintext []byte
	typ := recordType(record[0])
	payload := record[recordHeaderLen:]

	// In TLS 1.3, change_cipher_spec messages are to be ignored without being
	// decrypted. See RFC 8446, Appendix D.4.
	if hc.version == VersionTLS13 && typ == recordTypeChangeCipherSpec {
		return payload, typ, nil
	}

	paddingGood := byte(255)
	paddingLen := 0

	explicitNonceLen := hc.explicitNonceLen()

	if hc.cipher != nil {
		switch c := hc.cipher.(type) {
		case cipher.Stream:
			c.XORKeyStream(payload, payload)
		case aead:
			if len(payload) < explicitNonceLen {
				//return nil, 0, alertBadRecordMAC
				return nil, 0, nil
			}
			nonce := payload[:explicitNonceLen]
			if len(nonce) == 0 {
				nonce = hc.seq[:]
			}
			payload = payload[explicitNonceLen:]

			additionalData := hc.additionalData[:]
			if hc.version == VersionTLS13 {
				additionalData = record[:recordHeaderLen]
			} else {
				copy(additionalData, hc.seq[:])
				copy(additionalData[8:], record[:3])
				n := len(payload) - c.Overhead()
				additionalData[11] = byte(n >> 8)
				additionalData[12] = byte(n)
			}

			var err error
			plaintext, err = c.Open(payload[:0], nonce, payload, additionalData)
			if err != nil {
				//return nil, 0, alertBadRecordMAC
				return nil, 0, nil
			}
		case cbcMode:
			blockSize := c.BlockSize()
			minPayload := explicitNonceLen + roundUp(hc.mac.Size()+1, blockSize)
			if len(payload)%blockSize != 0 || len(payload) < minPayload {
				//return nil, 0, alertBadRecordMAC
				return nil, 0, nil
			}

			if explicitNonceLen > 0 {
				c.SetIV(payload[:explicitNonceLen])
				payload = payload[explicitNonceLen:]
			}
			c.CryptBlocks(payload, payload)

			// In a limited attempt to protect against CBC padding oracles like
			// Lucky13, the data past paddingLen (which is secret) is passed to
			// the MAC function as extra data, to be fed into the HMAC after
			// computing the digest. This makes the MAC roughly constant time as
			// long as the digest computation is constant time and does not
			// affect the subsequent write, modulo cache effects.
			if hc.version == VersionSSL30 {
				paddingLen, paddingGood = extractPaddingSSL30(payload)
			} else {
				paddingLen, paddingGood = extractPadding(payload)
			}
		default:
			panic("unknown cipher type")
		}

		if hc.version == VersionTLS13 {
			if typ != recordTypeApplicationData {
				//return nil, 0, alertUnexpectedMessage
				return nil, 0, nil
			}
			if len(plaintext) > maxPlaintext+1 {
				//return nil, 0, alertRecordOverflow
				return nil, 0, nil
			}
			// Remove padding and find the ContentType scanning from the end.
			for i := len(plaintext) - 1; i >= 0; i-- {
				if plaintext[i] != 0 {
					typ = recordType(plaintext[i])
					plaintext = plaintext[:i]
					break
				}
				if i == 0 {
					//return nil, 0, alertUnexpectedMessage
					return nil, 0, nil
				}
			}
		}
	} else {
		plaintext = payload
	}

	if hc.mac != nil {
		macSize := hc.mac.Size()
		if len(payload) < macSize {
			//return nil, 0, alertBadRecordMAC
			return nil, 0, nil
		}

		n := len(payload) - macSize - paddingLen
		n = subtle.ConstantTimeSelect(int(uint32(n)>>31), 0, n) // if n < 0 { n = 0 }
		record[3] = byte(n >> 8)
		record[4] = byte(n)
		remoteMAC := payload[n : n+macSize]
		localMAC := hc.mac.MAC(hc.seq[0:], record[:recordHeaderLen], payload[:n], payload[n+macSize:])

		if subtle.ConstantTimeCompare(localMAC, remoteMAC) != 1 || paddingGood != 255 {
			//return nil, 0, alertBadRecordMAC
			return nil, 0, nil
		}

		plaintext = payload[:n]
	}

	//hc.incSeq()
	return plaintext, typ, nil
}

func (hc *halfConn) explicitNonceLen() int {
	if hc.cipher == nil {
		return 0
	}

	switch c := hc.cipher.(type) {
	case cipher.Stream:
		return 0
	case aead:
		return c.explicitNonceLen()
	case cbcMode:
		// TLS 1.1 introduced a per-record explicit IV to fix the BEAST attack.
		if hc.version >= VersionTLS11 {
			return c.BlockSize()
		}
		return 0
	default:
		panic("unknown cipher type")
	}
}

func roundUp(a, b int) int {
	return a + (b-a%b)%b
}

func extractPaddingSSL30(payload []byte) (toRemove int, good byte) {
	if len(payload) < 1 {
		return 0, 0
	}

	paddingLen := int(payload[len(payload)-1]) + 1
	if paddingLen > len(payload) {
		return 0, 0
	}

	return paddingLen, 255
}

// extractPadding returns, in constant time, the length of the padding to remove
// from the end of payload. It also returns a byte which is equal to 255 if the
// padding was valid and 0 otherwise. See RFC 2246, Section 6.2.3.2.
func extractPadding(payload []byte) (toRemove int, good byte) {
	if len(payload) < 1 {
		return 0, 0
	}

	paddingLen := payload[len(payload)-1]
	t := uint(len(payload)-1) - uint(paddingLen)
	// if len(payload) >= (paddingLen - 1) then the MSB of t is zero
	good = byte(int32(^t) >> 31)

	// The maximum possible padding length plus the actual length field
	toCheck := 256
	// The length of the padded data is public, so we can use an if here
	if toCheck > len(payload) {
		toCheck = len(payload)
	}

	for i := 0; i < toCheck; i++ {
		t := uint(paddingLen) - uint(i)
		// if i <= paddingLen then the MSB of t is zero
		mask := byte(int32(^t) >> 31)
		b := payload[len(payload)-1-i]
		good &^= mask&paddingLen ^ mask&b
	}

	// We AND together the bits of good and replicate the result across
	// all the bits.
	good &= good << 4
	good &= good << 2
	good &= good << 1
	good = uint8(int8(good) >> 7)

	toRemove = int(paddingLen) + 1
	return
}
