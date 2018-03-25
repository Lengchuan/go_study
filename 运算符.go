package main

const (
	B  int = 1 << (iota * 10)
	KB
	MB
	GB
)

func main() {
	println(B)
	println(KB)
	println(MB)
	println(GB)
}
