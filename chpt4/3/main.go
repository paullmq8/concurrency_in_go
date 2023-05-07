package main

import (
	"bytes"
	"fmt"
	"sync"
)

func main() {
	printData := func(wg *sync.WaitGroup, data []byte) {
		defer wg.Done()

		var buff bytes.Buffer
		for _, b := range data {
			fmt.Fprintf(&buff, "%c", b)
		}
		fmt.Println(buff.String())
	}

	var wg sync.WaitGroup
	wg.Add(2)
	data := []byte("golang")
	s1 := data[:3]
	s2 := data[3:]
	fmt.Printf("%p %p\n", s1, s2)
	go printData(&wg, s1)
	go printData(&wg, s2)
	wg.Wait()
}
