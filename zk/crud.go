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
	conn, _, err := zk.Connect(servers, time.Second*5)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	path := "/lengchuan"
	data := []byte(`test`)
	// 权限
	acls := zk.WorldACL(zk.PermAll)
	var flags int32 = 0

	//2. 创建节点
	p, err := conn.Create(path, data, flags, acls)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(p)

	//3. 判断节点是否存在
	e, stat, err := conn.Exists(path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(e)

	e, _, err = conn.Exists("/test")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(e)

	//4.更新节点数据
	fmt.Println(stat.Version)
	s, err := conn.Set(path, []byte(`hello word`), stat.Version)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s.Version)

	//5.删除节点数据
	err = conn.Delete(path, s.Version)
	if err != nil {
		log.Fatal(err)
	}

	e, _, err = conn.Exists(path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(e)

}
