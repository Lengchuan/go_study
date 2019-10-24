package main

import (
	"fmt"
	"sync"
)

//生成自增id
type AutoInc struct {
	start, step int
	queue       chan int
	running     bool
}

func NewAI(start, step int) (ai *AutoInc) {
	ai = &AutoInc{
		start:   start,
		step:    step,
		running: true,
		queue:   make(chan int, 4),
	}
	go ai.process()
	return
}

func (ai *AutoInc) process() {
	defer func() { recover() }()
	for i := ai.start; ai.running; i = i + ai.step {
		ai.queue <- i
	}
}

func (ai *AutoInc) Id() int {
	return <-ai.queue
}

func (ai *AutoInc) Close() {
	ai.running = false
	close(ai.queue)
}

func main() {
	idmap := make(map[interface{}]interface{})
	lock := sync.Mutex{}

	ai := NewAI(1, 1)

	defer ai.Close()

	wg := sync.WaitGroup{}
	wg.Add(100000)
	for i := 0; i < 100000; i++ {
		go func(i int) {
			id := ai.Id()
			lock.Lock()
			idmap[id] = id
			lock.Unlock()
			wg.Done()
		}(i)
	}

	wg.Wait()

	if len(idmap) != 100000 {
		panic("error")
	}
	fmt.Println("exit 0 ")
}
