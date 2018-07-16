package main

import (
	"sync"
	"time"
	"fmt"
	"sync/atomic"
)

var (
	shutdown int64

	w sync.WaitGroup
)

func main()  {
	w.Add(2)

	go doWork("A")
	go doWork("B")

	time.Sleep(1 * time.Second)

	fmt.Println("shutdown now")
	atomic.StoreInt64(&shutdown, 1);

	w.Wait()

}
func doWork(name string) {
	defer w.Done()

	for {
		fmt.Printf("Doing %s Work \n", name)
		time.Sleep(250 * time.Millisecond)

		if atomic.LoadInt64(&shutdown) == 1  {
			fmt.Printf("Shutting %s Down\n", name)
			break
		}
	}
}