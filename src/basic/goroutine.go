package main

import "fmt"

var c chan string

func main() {

	//实现与主线程的交互，打印信息
	c = make(chan string)
	go Pingpong()
	for i := 0; i < 10; i++ {
		c <- fmt.Sprintf("From Main :Hello,#%d", i)
		fmt.Println(<-c)
	}
}

func Pingpong() {
	i := 0
	for {
		fmt.Println(<-c)
		c <- fmt.Sprintf("From Pingpong :Hi,#%d", i)
		i++
	}
}
