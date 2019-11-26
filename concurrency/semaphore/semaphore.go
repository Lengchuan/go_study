package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"time"
)

func main() {

	ctx := context.TODO()
	se := semaphore.NewWeighted(10)

	for i := 0; i < 100; i++ {
		if err := se.Acquire(ctx, 1); err != nil {
			break
		}
		go func() {
			defer se.Release(1)

			time.Sleep(time.Second * 1)
			fmt.Println(111)
		}()
	}

	time.Sleep(time.Minute)
}
