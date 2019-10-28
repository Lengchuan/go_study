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

	rootpath := "/lengchuan"
	data := []byte(`test`)
	// 权限
	acls := zk.WorldACL(zk.PermAll)

	p, err := conn.Create(rootpath, data, 0, acls)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(p)

	//监听子节点的事件
	_, _, ch, err := conn.ChildrenW(rootpath)
	if err != nil {
		log.Fatal(err)
	}

	for {

		select {

		case e := <-ch:

			switch e.Type {

			case zk.EventNodeCreated:
				fmt.Println("有新的节点创建: ", e.Path)
				break

			case zk.EventNodeDeleted:
				fmt.Println("有节点被删除: ", e.Path)
				break

			case zk.EventNodeDataChanged:
				fmt.Println("有节点数据发生变化: ", e.Path)
				break

			default:
				fmt.Println(fmt.Sprintf("type :%s err: %s  path:%s ", e.Type, e.Err, e.Path))

			}
		}
	}

}
