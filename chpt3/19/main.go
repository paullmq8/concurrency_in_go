package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	begin := make(chan struct{})
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-begin
			fmt.Println("goroutine " + strconv.Itoa(i) + " is done.")
		}(i)
	}
	fmt.Println("main goroutine is running.")
	close(begin)
	wg.Wait()
	fmt.Println("main goroutine is done.")
}
