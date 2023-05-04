package main

import "fmt"

func main() {
	var c1, c2 <-chan interface{}
	var c3 chan<- interface{}
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
