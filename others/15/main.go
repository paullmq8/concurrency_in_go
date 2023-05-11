package main

import (
	"fmt"
	"strconv"
)

func main() {
	toInt := func(done <-chan any, valueStream <-chan any) <-chan int {
		intStream := make(chan int)
		go func() {
			defer close(intStream)
			for {
				select {
				case <-done:
					return
				case val := <-valueStream:
					if v, ok := val.(int); ok {
						intStream <- v
					}
					if val == nil {
						return
					}
				}
			}
		}()
		return intStream
	}

	done := make(chan any)
	defer close(done)

	values := make(chan any)
	go func() {
		defer close(values)
		for i := 0; i < 10; i++ {
			if i%2 == 0 {
				values <- strconv.Itoa(i)
			} else {
				values <- i
			}
		}
	}()
	for i := range toInt(done, values) {
		fmt.Printf("%T %v\n", i, i)
	}
}
