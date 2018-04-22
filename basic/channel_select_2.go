package main

import "fmt"

func main() {

	c := make(chan int)

	go func() {
		for v := range c {
			fmt.Println(v)

		}
	}()

	//select发送
	for {
		select {
		case c <- 0:
		case c <- 1:
		case c <- 2:

		}

		//空select阻塞，GUI程序不让main退出
		//select {
		//
		//}
	}
}
