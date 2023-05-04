package main

import (
	"fmt"
	"sync"
	"time"
)

// channel + select to stop goroutine
func main() {
	stop := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-stop:
				fmt.Println("stopping goroutine...")
				return
			default:
				fmt.Println("goroutine is still running.")
				time.Sleep(1 * time.Second)
			}
		}
	}()
	time.Sleep(5 * time.Second)
	fmt.Println("stopping goroutine in main")
	stop <- struct{}{}
	wg.Wait()
	fmt.Println("goroutine is stopped.")
}
