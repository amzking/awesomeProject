package main

import (
	"runtime"
	"fmt"
	"sync"
)

var (
	counter int
	wg sync.WaitGroup
)

func main() {
	wg.Add(2)

	go incCounter(1)
	go incCounter(2)

	wg.Wait()

	fmt.Println("Final", counter)

}

func incCounter(id int) {
	defer wg.Done()

	for count:= 0; count < 2; count++ {
		value := counter

		runtime.Gosched()

		value++

		counter = value
	}

}