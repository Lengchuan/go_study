package main

import "fmt"

func main() {
	/*	var m map[int]string
		m = make(map[int]string)
		m1 := make(map[int]string)
		fmt.Println(m)
		fmt.Println(m1)

		m[1] = "ok"
		a := m[1]
		fmt.Println(m)
		fmt.Println(a)

		m[2] = "no ok"
		fmt.Println(m)
		delete(m, 1)
		fmt.Println(m)

		//复杂类型map

		var mm map[int]map[int]string
		mm = make(map[int]map[int]string)
		mm[1] = make(map[int]string)
		mm[1][1] = "PK"
		fmt.Println(mm)
		fmt.Println("--------")

		//不存在键值对时返回 false
		//多返回值判断键值对是否存在

		a, ok := mm[2][1]
		fmt.Println(a, ok)
		fmt.Println(a)
		if !ok {
			mm[2] = make(map[int]string)
		}
		mm[2][1] = "OK -- "
		a = mm[2][1]
		fmt.Println(a, ok)

		a1, ok1 := mm[2][1]
		fmt.Println(a1, ok1)*/

	//迭代操作
	/*

	s1 := []int{1, 2, 3, 4, 5, 6}

	for i, v := range s1 {
		println(i, v)
		s1[i] = v + i
		println("修改后", i, s1[i])
	}
	*/

	/*	m3 := make(map[string]string)
		m3["a"] = "aa"
		m3["b"] = "bb"
		m3["c"] = "cc"
		m3["d"] = "dd"

		for k, v := range m3 {
			println(k, v)
		}*/

	//以map为元素的slice
	/*	var sm = make([]map[string]string, 3, 10)

		for _, v := range sm {
			fmt.Println(v)
			//v 不会对Slice本身产生影响
			v := make(map[string]string, 1)
			v["aa"] = "aa"
			fmt.Println(v)
		}
		fmt.Println(sm)

		var sm1 = make([]map[string]string, 3, 10)

		for i := range sm1{
			fmt.Println(i)
			//v 不会对Slice本身产生影响
			sm1[i] = make(map[string]string, 1)
			sm1[i]["aa"] = "aa"
			fmt.Println(sm1[i])
			}
		fmt.Println(sm1)*/

	//map key value 交换

	m := map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
	s := make(map[int]string, 5)
	mm := make(map[string]int, 5)
	i := 0
	for k := range m {
		i++
		s[i] = m[k]
		//fmt.Println(k, m[k])
		mm[m[k]] = i
	}

	fmt.Println(mm)
}
