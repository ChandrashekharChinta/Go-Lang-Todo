package main

import (
	"fmt"
	"sync"
)

var counter int

func inc(wg *sync.WaitGroup, mu *sync.Mutex) {

	defer wg.Done()

	for i := 0; i < 1000; i++ {
		mu.Lock()
		counter++
		mu.Unlock()
	}

}

func main() {

	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(2)

	go inc(&wg, &mu)
	go inc(&wg, &mu)

	wg.Wait()

	fmt.Println(counter)

}
