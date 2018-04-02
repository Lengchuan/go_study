package main

import (
	"fmt"
)

func main() {
	fmt.Println(A(1, "A", 3))
	fmt.Println(BB(1, 3, "A"))
	a, b := 33, 22
	C(a, b, 3, 4, 5, 6, 6)
	s1 := []int{1, 2, 3, 4, 5}
	D(s1)
	fmt.Println(s1)

	aa := 33
	F(aa)
	fmt.Println(aa)

	E(&aa)
	fmt.Println(aa)

	g := G
	g()

	//匿名函数
	h := func() {
		fmt.Println("匿名函数")
	}
	h()

	//闭包
	i := closure(10)
	//fmt.Println(i)
	fmt.Println(i(1))
	fmt.Println(i(2))

	//defer
	/*	fmt.Println("defer1-----------")
		fmt.Println("a")
		defer fmt.Println("b")
		defer fmt.Println("c")*/

	/*fmt.Println("defer2-----------")
	for i := 0; i < 3; i++ {
		defer fmt.Println(i)
	}*/

	/*	fmt.Println("defer3-----------")
		for i := 0; i < 3; i++ {
			defer func() {
				fmt.Println(i)
			}()
		}*/


	//panic recover

}

func A(a int, b string, c int) (int string) {
	//fmt.Println(a, b, c)
	return
}

func BB(a int, c int, b string, ) int {
	fmt.Println(a, b, c)
	return a
}

//不定长变参 slice引用传递,不会改变Slice底层数组
func C(s ...int) {
	s[0] = 12
	s[1] = 23
	fmt.Println(s)
}

//会改变Slice底层数组
func D(s []int) {
	s[0] = 12
	s[1] = 23
}

//值拷贝
func F(a int) {
	a = 2
	fmt.Println(a)
}

//地址拷贝,会改变 变量值
func E(a *int) {
	*a = 2
	fmt.Println(*a)
}

//函数也是 一种类型
func G() {
	fmt.Println("Func G")
}

//闭包
func closure(x int) func(int) int {
	fmt.Println(&x)
	return func(y int) int {
		fmt.Println(&x)
		return x + y
	}
}

