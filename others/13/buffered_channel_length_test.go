package main

import (
	"fmt"
	"testing"
)

func bufferedChannelIO() {
	chanOwner := func() <-chan int {
		resultStream := make(chan int, 1)
		go func() {
			defer close(resultStream)
			for i := 0; i <= 5; i++ {
				resultStream <- i
			}
		}()
		return resultStream
	}

	resultStream := chanOwner()
	for _ = range resultStream {

	}
	fmt.Println("Done receiving!")
}

func BenchmarkBufferedChannelIO(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bufferedChannelIO()
	}
}
