package main

import "fmt"

type test struct {
}

type Person struct {
	Name    string
	Age     int
	Contact Contact
	Hobby   Hobby
}

type Contact struct {
	City, Email string
}

type Hobby struct {
	string
	int
}

func main() {
	a := test{}
	fmt.Println(a)

	p := Person{}
	p.Age = 12
	p.Name = "lengchuan"
	fmt.Println(p)
	fmt.Println(p.Name)

	p1 := Person{
		Name: "lengchuan",
		Age:  12,
	}

	fmt.Println(p1)

	//只会做值拷贝
	fa := AA
	fa(p1)
	fmt.Println(p1)

	//地址拷贝
	fb := AAA
	fb(&p1)
	fmt.Println(p1)
	fmt.Println()

	//指针
	point := &Person{
		"point",
		100,
		Contact{
			Email: "lishuijun1992@gmail.com",
			City:  "北京",
		},
		Hobby{
			"编程",
			100,
		},
	}

	fmt.Println(point)

	point.Name = "aaaa"
	point.Contact.City = "北京海淀"
	fmt.Println(point)

	//匿名结构
	n := struct {
		Name string
		Age  int
	}{
		"匿名结构",
		12,
	}

	fmt.Println(n)
}

func AA(person Person) {
	person.Age = 13
	fmt.Println("A -- Age", person.Age)
}

//地址拷贝
func AAA(person *Person) {
	person.Age = 130
	fmt.Println("AAA -- Age", person.Age)
}
