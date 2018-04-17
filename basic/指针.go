package main

func main() {
	a := 1
	var p *int = &a
	a++
	println(*p)
}
