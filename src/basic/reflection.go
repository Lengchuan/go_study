package main

import (
	"reflect"
	"fmt"
)

func main() {
	u := User{1, "lengchuan", 22}
	Info(u)
	Info(&u) //传入地址

	//获取匿名字段
	m := Manager{User{1, "lengchuan", 22}, "11111"}
	t := reflect.TypeOf(m)

	fmt.Printf("%#v\n", t.Field(0))
	fmt.Printf("%#v\n", t.Field(1))
	fmt.Printf("%#v\n", t.Field)

	//获取值
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 0}))
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 1}))
	fmt.Printf("%#v\n", t.FieldByIndex([]int{0, 2}))
	//fmt.Printf("%#v\n",t.FieldByIndex([]int{1,0})) 报错

	//通过反射修改值
	fmt.Println(m.title)
	v := reflect.ValueOf(&m.title)
	v.Elem().SetString("lengchuan1111")
	fmt.Println(m.title)

	v1 := reflect.ValueOf(&m.Id)
	v1.Elem().SetInt(100)
	fmt.Println(m.Id)

	u2 := User{2, "lengchuan2", 100}
	fmt.Println(u2.Name)
	Set(&u2)
	fmt.Println(u2.Name)

	v2 := reflect.ValueOf(u2)
	mv := v2.MethodByName("Hello1")

	args := []reflect.Value{reflect.ValueOf("lengchuan")}
	mv.Call(args)

}

type User struct {
	Id   int
	Name string
	Age  int
}

type Manager struct {
	//匿名字段
	User
	title string
}

func (u User) Hello() {
	fmt.Println("hello world")
}

func (u User) Hello1(name string) {
	fmt.Println("hello world ", name)
}

func Info(o interface{}) {
	//反射获取类型
	t := reflect.TypeOf(o)
	fmt.Println("Type:", t.Name())

	//
	if k := t.Kind(); k != reflect.Struct {
		fmt.Println("Is not a Struct ")
		return
	}

	//反射获取字段值
	v := reflect.ValueOf(o)
	fmt.Println("Fileds :")
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		val := v.Field(i).Interface()
		fmt.Printf("%6s:%v = %v\n", f.Name, f.Type, val)
	}

	//反射获取方法
	fmt.Println(t.NumMethod())
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		fmt.Printf("%6s: %v\n", m.Name, m.Type)
	}

}

func Set(o interface{}) {
	v := reflect.ValueOf(o)
	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() {
		fmt.Println("Can't set")
		return
	}

	v = v.Elem()
	f := v.FieldByName("Name")
	if !f.IsValid() {
		fmt.Println("Not filed 'Name'")
		return
	}
	if f.Kind() == reflect.String {
		f.SetString("Go Go!!!")
	}
}
