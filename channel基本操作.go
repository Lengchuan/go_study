package main

import (
	"fmt"
)

func main() {
	//go Go()

	//time.Sleep(2 * time.Second)

	c := make(chan bool, 1)
	go func() {
		fmt.Println("Go Go Go!!!")
		//向channel里放入值
		c <- true
	}()

	//从channel里取出值
	<-c

	//for v := range c {
	//	fmt.Println(v)
	//}
}

func Go() {

	fmt.Println("Go Go !!!")
}
