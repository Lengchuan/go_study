package main

import "fmt"

func main() {

	c := make(chan bool, 1)
	go func() {
		fmt.Println("Go Go Go!!!")
		//向channel里放入值
		c <- true

		//关闭,不关闭会导致死锁
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}
}
