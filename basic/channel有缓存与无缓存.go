package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	//设置cpu数
	runtime.GOMAXPROCS(runtime.NumCPU())
	//c := make(chan bool, 10)
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go Caculator(i, &wg)
	}
	//for i := 0; i < 10; i++ {
	//	<-c
	//}

	wg.Wait()

}

func Caculator(index int, wg *sync.WaitGroup) {
	a := 1
	for i := 0; i < 10000000; i++ {
		a += i
	}
	fmt.Println(index, a)
	//c <- true
	wg.Done()
}
