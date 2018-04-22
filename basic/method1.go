package main

import "fmt"

func main() {

	var t Ti
	t = 0
	t.Increase(100)
	fmt.Println(t)
}

type Ti int

func (t *Ti) Increase(num int) {
	*t += Ti(num) //需要强制转换类型
}
