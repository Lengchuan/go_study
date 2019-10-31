package main

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"log"
	"reflect"
)

// tls 握手ClientHello信息解析，抓取的整个数据TLS的数据包
//工298字节，Ethernet层(14)+Internet层(20)+Trnsmission层(20) = 54字节固定
//SSL/TSL 层 = 298-54 = 244字节
var helloData = []byte{
	0x1c, 0xab, 0x34, 0x12, 0xd3, 0xbe, 0x00, 0xe0, 0x4c, 0x36, 0x05, 0x96, 0x08, 0x00, 0x45, 0x00,
	0x01, 0x1c, 0x4f, 0x69, 0x40, 0x00, 0x40, 0x06, 0xd8, 0x40, 0xac, 0x10, 0x45, 0xdf, 0x0b, 0x00,
	0x15, 0x43, 0xee, 0x57, 0x01, 0xbb, 0x78, 0x78, 0x7e, 0x3d, 0x4e, 0x2f, 0xe2, 0xe8, 0x50, 0x18,
	0x02, 0x01, 0xc1, 0x0f, 0x00, 0x00, 0x16, 0x03, 0x01, 0x00, 0xef, 0x01, 0x00, 0x00, 0xeb, 0x03,
	0x03, 0x91, 0xa5, 0x00, 0xda, 0x9a, 0x11, 0xc6, 0x8d, 0x08, 0x33, 0x57, 0x3f, 0x96, 0x0a, 0x9c,
	0xf4, 0x19, 0x99, 0xb7, 0x3d, 0xdb, 0x91, 0xed, 0xc8, 0x88, 0x74, 0xd0, 0xdb, 0xb4, 0x91, 0xe8,
	0x5d, 0x00, 0x00, 0x5c, 0xc0, 0x2f, 0xc0, 0x2b, 0xc0, 0x30, 0xc0, 0x2c, 0x00, 0x9e, 0xc0, 0x27,
	0x00, 0x67, 0xc0, 0x28, 0x00, 0x6b, 0x00, 0xa3, 0x00, 0x9f, 0xcc, 0xa9, 0xcc, 0xa8, 0xcc, 0xaa,
	0xc0, 0xaf, 0xc0, 0xad, 0xc0, 0xa3, 0xc0, 0x9f, 0x00, 0xa2, 0xc0, 0xae, 0xc0, 0xac, 0xc0, 0xa2,
	0xc0, 0x9e, 0xc0, 0x24, 0x00, 0x6a, 0xc0, 0x23, 0x00, 0x40, 0xc0, 0x0a, 0xc0, 0x14, 0x00, 0x39,
	0x00, 0x38, 0xc0, 0x09, 0xc0, 0x13, 0x00, 0x33, 0x00, 0x32, 0x00, 0x9d, 0xc0, 0xa1, 0xc0, 0x9d,
	0x00, 0x9c, 0xc0, 0xa0, 0xc0, 0x9c, 0x00, 0x3d, 0x00, 0x3c, 0x00, 0x35, 0x00, 0x2f, 0x00, 0xff,
	0x01, 0x00, 0x00, 0x66, 0x00, 0x00, 0x00, 0x1c, 0x00, 0x1a, 0x00, 0x00, 0x17, 0x61, 0x7a, 0x75,
	0x72, 0x65, 0x2d, 0x77, 0x77, 0x77, 0x2e, 0x71, 0x69, 0x6e, 0x67, 0x79, 0x69, 0x64, 0x61, 0x69,
	0x2e, 0x63, 0x6f, 0x6d, 0x00, 0x0b, 0x00, 0x04, 0x03, 0x00, 0x01, 0x02, 0x00, 0x0a, 0x00, 0x0a,
	0x00, 0x08, 0x00, 0x1d, 0x00, 0x17, 0x00, 0x19, 0x00, 0x18, 0x00, 0x23, 0x00, 0x00, 0x00, 0x16,
	0x00, 0x00, 0x00, 0x17, 0x00, 0x00, 0x00, 0x0d, 0x00, 0x20, 0x00, 0x1e, 0x06, 0x01, 0x06, 0x02,
	0x06, 0x03, 0x05, 0x01, 0x05, 0x02, 0x05, 0x03, 0x04, 0x01, 0x04, 0x02, 0x04, 0x03, 0x03, 0x01,
	0x03, 0x02, 0x03, 0x03, 0x02, 0x01, 0x02, 0x02, 0x02, 0x03,
}

var testClientHelloDecoded = &layers.TLS{
	BaseLayer: layers.BaseLayer{
		Contents: helloData[54:],
		Payload:  nil,
	},
	ChangeCipherSpec: nil,
	Handshake: []layers.TLSHandshakeRecord{
		{
			layers.TLSRecordHeader{
				ContentType: 22,     //ClientHello 1字节
				Version:     0x0301, //TLS 1.0 2 字节
				Length:      239,    //数据包大小(本身存储占2字节) 244-1-2-2 = 239
			},
		},
	},
	AppData: nil,
	Alert:   nil,
}

var testTLSDecodeOptions = gopacket.DecodeOptions{
	SkipDecodeRecovery:       true,
	DecodeStreamsAsDatagrams: true,
}

func main() {
	p := gopacket.NewPacket(helloData, layers.LinkTypeEthernet, testTLSDecodeOptions)
	if p.ErrorLayer() != nil {
		log.Fatal("Failed to decode packet:", p.ErrorLayer().Error())
	}
	checkLayers(p, []gopacket.LayerType{layers.LayerTypeEthernet, layers.LayerTypeIPv4, layers.LayerTypeTCP, layers.LayerTypeTLS})

	if got, ok := p.Layer(layers.LayerTypeTLS).(*layers.TLS); ok {
		want := testClientHelloDecoded
		if !reflect.DeepEqual(got, want) {
			fmt.Println(fmt.Sprintf("TLS ClientHello packet processing failed:\ngot:\n%#v\n\nwant:\n%#v\n\n", got, want))
		}
	} else {
		fmt.Println("No TLS layer type found in packet")
	}

	//解析出 Random
	//从开始到Random的偏移量为 54+1+2+2+1+3+2+4 = 69       69+32 =101
	randomBytes := helloData[69:101]
	fmt.Println(string(randomBytes[:]))
}

func checkLayers(p gopacket.Packet, want []gopacket.LayerType) {
	layers := p.Layers()
	fmt.Println("Checking packet layers, want", want)
	for _, l := range layers {
		fmt.Println(fmt.Sprintf("  Got layer %v, %d bytes, payload of %d bytes", l.LayerType(),
			len(l.LayerContents()), len(l.LayerPayload())))
	}
	fmt.Println(p)
	if len(layers) < len(want) {
		fmt.Println(fmt.Sprintf("  Number of layers mismatch: got %d want %d", len(layers),
			len(want)))
		return
	}
	for i, l := range want {
		if l == gopacket.LayerTypePayload {
			// done matching layers
			return
		}

		if layers[i].LayerType() != l {
			fmt.Println(fmt.Sprintf("  Layer %d mismatch: got %v want %v", i,
				layers[i].LayerType(), l))
		}
	}

}

//Secure Sockets Layer --- 244字节
//    TLSv1.2 Record Layer: Handshake Protocol: Client Hello
//        Content Type: Handshake (22) --- 1字节
//        Version: TLS 1.0 (0x0301) ---2字节
//        Length: 239 --- 2字节 （244-1-2-2 = 239）
//        Handshake Protocol: Client Hello ---239字节
//            Handshake Type: Client Hello (1) ---1字节
//            Length: 235 ---3字节
//            Version: TLS 1.2 (0x0303) ---2字节
//            Random: 91a500da9a11c68d0833573f960a9cf41999b73ddb91edc8... ---32字节
//                GMT Unix Time: Jun  7, 2047 16:46:18.000000000 中国标准时间 --- 时间 4字节
//                Random Bytes: 9a11c68d0833573f960a9cf41999b73ddb91edc88874d0db... --随机数 28字节
//            Session ID Length: 0
//            Cipher Suites Length: 92
//            Cipher Suites (46 suites)
//            Compression Methods Length: 1
//            Compression Methods (1 method)
//                Compression Method: null (0)
//            Extensions Length: 102
//            Extension: server_name (len=28)
//                Type: server_name (0)
//                Length: 28
//                Server Name Indication extension
//                    Server Name list length: 26
//                    Server Name Type: host_name (0)
//                    Server Name length: 23
//                    Server Name: azure-www.qingyidai.com
//            Extension: ec_point_formats (len=4)
//                Type: ec_point_formats (11)
//                Length: 4
//                EC point formats Length: 3
//                Elliptic curves point formats (3)
//                    EC point format: uncompressed (0)
//                    EC point format: ansiX962_compressed_prime (1)
//                    EC point format: ansiX962_compressed_char2 (2)
//            Extension: supported_groups (len=10)
//                Type: supported_groups (10)
//                Length: 10
//                Supported Groups List Length: 8
//                Supported Groups (4 groups)
//            Extension: SessionTicket TLS (len=0)
//                Type: SessionTicket TLS (35)
//                Length: 0
//                Data (0 bytes)
//            Extension: encrypt_then_mac (len=0)
//                Type: encrypt_then_mac (22)
//                Length: 0
//            Extension: extended_master_secret (len=0)
//                Type: extended_master_secret (23)
//                Length: 0
//            Extension: signature_algorithms (len=32)
//                Type: signature_algorithms (13)
//                Length: 32
//                Signature Hash Algorithms Length: 30
//                Signature Hash Algorithms (15 algorithms)
//                    Signature Algorithm: rsa_pkcs1_sha512 (0x0601)
//                        Signature Hash Algorithm Hash: SHA512 (6)
//                        Signature Hash Algorithm Signature: RSA (1)
//                    Signature Algorithm: SHA512 DSA (0x0602)
//                        Signature Hash Algorithm Hash: SHA512 (6)
//                        Signature Hash Algorithm Signature: DSA (2)
//                    Signature Algorithm: ecdsa_secp521r1_sha512 (0x0603)
//                        Signature Hash Algorithm Hash: SHA512 (6)
//                        Signature Hash Algorithm Signature: ECDSA (3)
//                    Signature Algorithm: rsa_pkcs1_sha384 (0x0501)
//                        Signature Hash Algorithm Hash: SHA384 (5)
//                        Signature Hash Algorithm Signature: RSA (1)
//                    Signature Algorithm: SHA384 DSA (0x0502)
//                        Signature Hash Algorithm Hash: SHA384 (5)
//                        Signature Hash Algorithm Signature: DSA (2)
//                    Signature Algorithm: ecdsa_secp384r1_sha384 (0x0503)
//                        Signature Hash Algorithm Hash: SHA384 (5)
//                        Signature Hash Algorithm Signature: ECDSA (3)
//                    Signature Algorithm: rsa_pkcs1_sha256 (0x0401)
//                        Signature Hash Algorithm Hash: SHA256 (4)
//                        Signature Hash Algorithm Signature: RSA (1)
//                    Signature Algorithm: SHA256 DSA (0x0402)
//                        Signature Hash Algorithm Hash: SHA256 (4)
//                        Signature Hash Algorithm Signature: DSA (2)
//                    Signature Algorithm: ecdsa_secp256r1_sha256 (0x0403)
//                        Signature Hash Algorithm Hash: SHA256 (4)
//                        Signature Hash Algorithm Signature: ECDSA (3)
//                    Signature Algorithm: SHA224 RSA (0x0301)
//                        Signature Hash Algorithm Hash: SHA224 (3)
//                        Signature Hash Algorithm Signature: RSA (1)
//                    Signature Algorithm: SHA224 DSA (0x0302)
//                        Signature Hash Algorithm Hash: SHA224 (3)
//                        Signature Hash Algorithm Signature: DSA (2)
//                    Signature Algorithm: SHA224 ECDSA (0x0303)
//                        Signature Hash Algorithm Hash: SHA224 (3)
//                        Signature Hash Algorithm Signature: ECDSA (3)
//                    Signature Algorithm: rsa_pkcs1_sha1 (0x0201)
//                        Signature Hash Algorithm Hash: SHA1 (2)
//                        Signature Hash Algorithm Signature: RSA (1)
//                    Signature Algorithm: SHA1 DSA (0x0202)
//                        Signature Hash Algorithm Hash: SHA1 (2)
//                        Signature Hash Algorithm Signature: DSA (2)
//                    Signature Algorithm: ecdsa_sha1 (0x0203)
//                        Signature Hash Algorithm Hash: SHA1 (2)
//                        Signature Hash Algorithm Signature: ECDSA (3)
