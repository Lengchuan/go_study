package main

import "fmt"

func main() {

	//先定义两个channel
	c1, c2 := make(chan int), make(chan string)
	//用于通信的channel
	o := make(chan bool, 2)
	go func() {
		//a, b := false, false
		for {
			select {
			case v, ok := <-c1:
				if !ok {
					fmt.Println("c1")
					//if !a {//导致程序死循环，不能判断两个channel都关闭
					o <- true
					//a = true
					//}

					break
				}
				fmt.Println(v)
			case v, ok := <-c2:
				if !ok {
					fmt.Println("c2")
					//if !b {
					//	b = true
					o <- true
					//}
					break

				}

				fmt.Println(v)

			}
		}
	}()

	c1 <- 1
	c2 <- "ok"
	c1 <- 2
	c2 <- "hello"

	close(c1)
	//close(c1)
	//close(c2)

	for i := 0; i < 2; i++ {
		<-o
	}
	//<-o
}
