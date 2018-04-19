package main

import "fmt"

func main() {

	//var a [2] int
	//var b [1]  int
	c := [2]int{1, 2}
	d := [...]int{1, 2, 3, 4, 5}
	//索引赋值
	e := [...] int{0: 1, 1: 2, 2: 3}
	f := [...] int{19: 1}

	//指针数组
	var g *[20] int = &f

	x, y := 1, 2
	h := [...]*int{&x, &y}

	// == 和！=比较

	i := [...] int{1, 2}
	j := [...] int{1, 2}

	//new
	k := new([10]int)

	//多维数组

	l := [2][3] int{
		{1, 2, 3},
		{1, 2, 3}}

	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(len(c))
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)

	fmt.Println(i == j)
	fmt.Println(i != j)
	fmt.Println(k)
	fmt.Println(k[0])
	fmt.Println(&k[0])
	fmt.Println(&k[1])
	fmt.Println(&k[2 ])
	fmt.Println(&k[3])
	fmt.Println(&k[4])
	fmt.Println(l)
	fmt.Println(l[0][1])
	fmt.Println(l[0][2])
	fmt.Println(l)

}
