package main

import (
	"fmt"
)

func getAReadOnlyChannel() <-chan int {
	fmt.Println("invoke getAReadOnlyChannel")
	c := make(chan int)

	go func() {
		//time.Sleep(3 * time.Second)
		c <- 1
	}()

	return c
}

func getASlice() *[5]int {
	fmt.Println("invoke getASlice")
	var a [5]int
	return &a
}

func getAWriteOnlyChannel() chan<- int {
	fmt.Println("invoke getAWriteOnlyChannel")
	ch := make(chan int)
	/*	go func() {
			println(<-ch)
		}()
	*/return ch
}

func getANumToChannel() int {
	fmt.Println("invoke getANumToChannel")
	return 2
}

func main() {
	select {
	// 从channel接收数据
	case (getASlice())[0] = <-getAReadOnlyChannel():
		fmt.Println("recv something from a readonly channel")
	// 将数据发送到channel
	case getAWriteOnlyChannel() <- getANumToChannel():
		fmt.Println("send something to a writeonly channel")
	default:
		fmt.Println("aaa")
	}
}
