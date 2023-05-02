package main

import (
	"context"
	"fmt"
	"sync"
)

func main() {

	type Button struct {
		Clicked *sync.Cond
	}

	button := Button{
		Clicked: sync.NewCond(&sync.Mutex{}),
	}

	subscribe := func(c *sync.Cond, name string, fn func()) {
		var goroutineRunning sync.WaitGroup
		goroutineRunning.Add(1)
		go func() {
			goroutineRunning.Done()
			c.L.Lock()
			defer c.L.Unlock()
			fmt.Println(name + " is waiting...")
			c.Wait()
			fmt.Println(name + "after waiting...")
			fn()
		}()
		goroutineRunning.Wait()
	}

	var clickRegistered sync.WaitGroup
	clickRegistered.Add(1)
	subscribe(button.Clicked, "window", func() {
		fmt.Println("Maximizing window.")
		clickRegistered.Done()

		context.Background()
	})

	/*subscribe(button.Clicked, "mouse", func() {
		fmt.Println("Mouse clicked.")
		clickRegistered.Done()
	})

	subscribe(button.Clicked, "display", func() {
		fmt.Println("Displaying annoying dialogue box!")
		clickRegistered.Done()
	})*/

	button.Clicked.Broadcast()
	fmt.Println("after broadcast")
	clickRegistered.Wait()
	fmt.Println("after wait")
}
