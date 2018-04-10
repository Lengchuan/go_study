package main

import "fmt"

func main() {
	s := S1{}
	s.Print()

	s1 := S2{}
	s1.Print()

	s3 := S3{}
	s3.Print()
	fmt.Println(s3.Name)

	(*S3).Print(&s3);

	//方法访问权限
	fmt.Print(s3.age)
}

type S1 struct {
	Name string
}
type S2 struct {
	Name string
}
type S3 struct {
	Name string
	age  int
}

func (s S1) Print() {
	fmt.Println("A")
}

func (s S2) Print() {
	fmt.Println("B")
}

func (s *S3) Print() {
	s.Name = "CC"
	s.age = 100
	fmt.Println("C")
}
