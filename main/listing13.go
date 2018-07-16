package main

import (
	"sync"
	"fmt"
	"sync/atomic"
	"runtime"
)

var (
	counter2 int64
	wg1 sync.WaitGroup
)

func main() {
	wg1.Add(2)

	go incCounter2(1)
	go incCounter2(2)

	wg1.Wait()

	fmt.Println("Final Counter", counter2)
}

func incCounter2(id int) {
	defer wg1.Done()

	for count := 0; count < 2; count++ {
		atomic.AddInt64(&counter2, 1)
		runtime.Gosched()
	}
}