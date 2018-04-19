package main

import "fmt"

func main() {
	var a Usb
	a = PhoneConnecter{name: "PhoneConnecter"}
	a.Connect()
	fmt.Println(a.Name())
	DisConnect(a)

	var b Connecter
	b = PhoneConnecter{"PhoneConnecter b"}
	b = Connecter(b)
	b.Connect()

}

//类似于java中的object，所有方法都实现了空接口
type empty interface {
}

type Usb interface {
	Name() string
	Connecter //嵌入接口
}

type Connecter interface {
	Connect()
}

type PhoneConnecter struct {
	name string
}

//只需要实现interface的同名方法
func (pc PhoneConnecter) Name() string {
	return pc.name
}

func (pc PhoneConnecter) Connect() {

	fmt.Println("Connect by ", pc.name)
}

func DisConnect(usb Usb) {
	/*if pc,ok:=usb.(PhoneConnecter);ok {
		fmt.Println("DisConnect：",pc.name)
		return
	}
	fmt.Println("Unkown device")*/

	//使用switch
	switch v := usb.(type) {
	case PhoneConnecter:
		fmt.Println("DisConnect:", v.name)
	default:
		fmt.Println("Unkown device")
	}
}
