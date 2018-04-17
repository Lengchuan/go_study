package main

import "fmt"

func main() {

	a := []int{2, 4, 65, 412, 1, 23, 12}

	fmt.Println(a)

	num := len(a)

	for i := 0; i < num; i++ {
		for j := i + 1; j < num; j++ {
			if (a[i] < a[j]) {
				tmp := a[i]
				tmp = a[j]
				a[j] = a[i]
				a[i] = tmp
				fmt.Println(i)
				fmt.Println(a)
			}
		}
	}
	//fmt.Println(a)


}
