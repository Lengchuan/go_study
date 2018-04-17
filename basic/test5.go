package main

import "strconv"

func main() {
	var a int = 65
	b := string(a)
	println(b)
	println(strconv.Itoa(a))
}
