package main
/*
const a int = 1
const b = 'A'
const (
	c = a
	d = a + 1
	e
)*/

//const f, g, k = b, d, a

//枚举
const (
	a = 'A'
	b
	c = iota
	d

)
func main() {

	println(a)
	println(b)
	println(c)
	println(d)
	//println(e)
}
