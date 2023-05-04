package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	var c1, c2 <-chan int
	select {
	case <-c1:
	case <-c2:
	default:
		// You can see that it ran the default statement almost instantaneously.
		fmt.Printf("In default after %v\n\n", time.Since(start))
	}
}
