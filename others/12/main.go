package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var stdoutBuff bytes.Buffer
	defer stdoutBuff.WriteTo(os.Stdout)
	fmt.Fprintln(&stdoutBuff, "Producer Done.")
	for i := 0; i < 100000000; i++ {
		fmt.Fprintln(&stdoutBuff, strconv.Itoa(i))
	}
}
