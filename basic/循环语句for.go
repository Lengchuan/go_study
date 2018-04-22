package main

func main() {

	//1.无限循环
	a := 1
	for {
		a ++
		if a > 3 {
			break
		}
	}
	println(a)

	//2.条件循环
	b := 1
	for b <= 3 {
		b++
	}
	println(b)

	//3.计数器
	c := 1
	for i := 0; i < 3; i++ {
		c++
	}
	println(c)

	d := "ssssssssssss"
	for i, l := 0, len(d); i < l; i++ {
		println(i)
	}
	println(d)
}
