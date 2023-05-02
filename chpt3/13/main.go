package main

import (
	"fmt"
	"sync"
)

func main() {
	var onceA, onceB sync.Once
	countA := 0
	countB := 1
	increase := func() {
		countA++
	}
	decrease := func() {
		countB--
	}
	onceA.Do(increase)
	onceB.Do(decrease)
	fmt.Println(countA, countB)
	var onceC sync.Once
	onceC.Do(increase)
	onceC.Do(decrease)
	fmt.Println(countA, countB)
}
