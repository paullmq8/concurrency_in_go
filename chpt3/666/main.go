package main

import "runtime"

func main() {
	go func() {
		for i := 0; i < 10000; i++ {
			println(i)
		}
	}()
	runtime.GC()
	println("===========", 2)
}
