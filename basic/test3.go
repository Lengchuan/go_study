package main

import (
	"fmt"
	"math"
)

func main() {
	var b = bool(false)
	var a int
	var c int32
	var d int64
	var e float32
	var f float64
	var aa []  int
	cc:=1
	fmt.Printf("%t", b)
	fmt.Printf("\n")

	fmt.Printf("%d", a)
	fmt.Printf("\n")

	fmt.Printf("%t", b)
	fmt.Printf("\n")

	fmt.Printf("%d", c)
	fmt.Printf("\n")

	fmt.Printf("%d", d)
	fmt.Printf("\n")

	fmt.Printf("%f", e)
	fmt.Printf("\n")

	fmt.Printf("%f", f)
	fmt.Printf("\n")

	fmt.Println(aa)
	fmt.Println(cc)

	fmt.Println(math.MaxInt8)
	fmt.Println(math.MinInt8)
	fmt.Println(cc)
	//fmt.Println(bb)
}
