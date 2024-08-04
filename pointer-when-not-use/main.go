package main

import (
	"fmt"
	"sync"
)

func increment(ptr *int, wg *sync.WaitGroup) {
	*ptr++
	wg.Done()
}

func main() {
	var counter int
	var wg sync.WaitGroup
	for i := 0; i < 10000000; i++ {
		wg.Add(1)

		go increment(&counter, &wg)

		fmt.Println("Counter : ", counter)
	}
}
