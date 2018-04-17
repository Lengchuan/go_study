package main

import "fmt"

func main() {
	//声明一个切片
	var a []int
	fmt.Println(a)

	b := [10]int{1: 1, 9: 10}
	fmt.Println(b)

	c := b[9]
	fmt.Println(c)

	d := b[5:10]
	fmt.Println(d)

	e := b[:5]
	fmt.Println(e)
	//
	f := b[5:]
	fmt.Println(f)

	//make

	g := make([]int, 3, 10) //3个元素，初始10个长度的内存，每次增加2倍
	fmt.Println(g)
	fmt.Println(len(g), cap(g))

	h := [...]string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	fmt.Println(h)

	sa := h[2:5]
	fmt.Println(sa)

	sb := h[0:]
	fmt.Println(sb)

	sc := sb [2:5]
	fmt.Println(sc)

	sc[1] = "A"
	fmt.Println(sb)
	fmt.Println(h)

	//append

	s1 := make([]int, 3, 6)
	fmt.Println("%v", s1)
	s1 = append(s1, 1, 2, 3, 4, 5)
	fmt.Println("%v", s1)

	s2 := []int{1, 2, 3, 4, 5, 6}
	s3 := []int{7, 8, 9}
	fmt.Println(s3)

	copy(s3, s2)
	fmt.Println(s3)

}
