package main

import (
	"bytes"
	"fmt"
)

func main() {
	s := "加拿大"
	data := []byte(s)
	var buff bytes.Buffer
	for _, b := range data {
		fmt.Fprintf(&buff, "%b", b)
	}
	fmt.Println(buff.String())

	for _, v := range s {
		fmt.Printf("%c", v)
	}
}
