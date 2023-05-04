package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	var stdoutBuff bytes.Buffer
	defer stdoutBuff.WriteTo(os.Stdout)

	intStream := make(chan int, 4)
	go func() {
		defer close(intStream)
		defer fmt.Fprintln(&stdoutBuff, "Producer Done.")
		for i := 0; i < 5; i++ {
			intStream <- i
			fmt.Fprintf(&stdoutBuff, "Sending: %d\n", i)
		}
	}()

	for integer := range intStream {
		fmt.Fprintf(&stdoutBuff, "Received %v.\n", integer)
	}
}
