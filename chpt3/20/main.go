package main

import "fmt"

func main() {
	var c1, c2 <-chan any
	var c3 chan<- any
	fmt.Println(c1, c2, c3)
	select {
	case <-c1:
	//Do something
	case <-c2:
	//Do something
	case c3 <- struct{}{}:
		//Do something
	}
}
