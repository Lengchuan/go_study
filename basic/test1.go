package main

import (
	"fmt"
	_ "fmt"
	_ "io"
)

//常量
const pi = 3.14

//全局变量声明与赋值
var name = "lengchuan"

//一般类型声明
type inttype int

//结构体
type gostruct struct {
}

//接口声明
type gointerface interface {
}

//main函数作为程序主入口

func main() {
	sayhello()

	n := 363
	fmt.Println(byte(n))
	fmt.Println(363 >> 8)
}

func sayhello() {
	println("hello  world!!!")
	println(name)
}
