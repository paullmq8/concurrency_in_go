package main

import "fmt"

func main() {
	var ch chan int
	ch = make(chan int, 7)
	fmt.Println(ch)
	var ch1 chan int
	fmt.Println(ch1)
	for i := range ch1 {
		fmt.Println(i)
	}
}
