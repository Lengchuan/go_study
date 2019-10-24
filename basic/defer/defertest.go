package main

import "fmt"

//http://golang.org/ref/spec#defer_statements
//https://tiancaiamao.gitbooks.io/go-internals/content/zh/03.4.html
// return xxx 非原子性
func main() {
	fmt.Println(f())
	fmt.Println(f1())
	fmt.Println(f2())
}

func f() (result int) {
	defer func() {
		result++
	}()
	return 0
}

func f1() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

func f2() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}
