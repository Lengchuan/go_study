package main

import "fmt"

var (
	a int   = 1
	b int8  = 2
	c int16 = 3
)

func main() {
	var aa, bb, cc int = 11, 22, 33

	fmt.Println(aa)
	fmt.Println(bb)
	fmt.Println(cc)

	var d float32 = 33.33
	var e = int(d)
	var f float64 = 33.33
	println(d)
	println(e)
	println(f)

	g :=float32(e)

	println(g)
}
