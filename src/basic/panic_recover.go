package main

import "fmt"

func main() {
	//panic
	/*	fmt.Println("panic")

		P1()
		P2()
		P3() //不会执行*/

	fmt.Println("recover")
	P1()
	P4()
	P3() //会执行
}

//panic
func P1() {
	fmt.Println("Func P1")
}

func P2() {
	panic("panic in P2")
}

func P3() {
	fmt.Println("Func P3")
}

func P4() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover in  P4")
		}
	}()
	panic("panic in P21")
	panic("panic in P22")
	panic("panic in P22")
	panic("panic in P22")
	panic("panic in P22")
}
