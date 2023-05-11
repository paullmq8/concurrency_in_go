package main

import (
	"fmt"
	"sync"
)

func main() {
	tryPool()
}

func tryPool() {
	var numCalcsCreated int
	calcPool := &sync.Pool{
		New: func() any {
			numCalcsCreated += 1
			mem := make([]byte, 1024)
			return &mem
		},
	}

	// Seed the pool with 4KB
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())
	calcPool.Put(calcPool.New())

	ch := make(chan bool)

	const numWorkers = 1024 * 1024
	var wg sync.WaitGroup
	wg.Add(numWorkers)
	for i := numWorkers; i > 0; i-- {
		go func() {
			defer wg.Done()
			<-ch
			mem := calcPool.Get().(*[]byte)
			defer calcPool.Put(mem)
			// Assume something interesting, but quick is being done with this memory.
		}()
	}
	ch <- true
	wg.Wait()
	fmt.Printf("%d calculators were created.", numCalcsCreated)
}
