package main

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"log"
	"time"
)

func main() {
	//zk 地址
	servers := []string{"127.0.0.1:2181"}

	//1. 连接
	option := zk.WithEventCallback(callback)
	conn, _, err := zk.Connect(servers, time.Second*5, option)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	path := "/lengchuan"
	data := []byte(`test`)
	// 权限
	acls := zk.WorldACL(zk.PermAll)
	var flags int32 = 0

	//2.创建
	p, err := conn.Create(path, data, flags, acls)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(p)
}

//callback
func callback(event zk.Event) {
	fmt.Println("________________________________")
	fmt.Println("path:", event.Path)
	fmt.Println("type:", event.Type.String())
	fmt.Println("state:", event.State.String())
}
