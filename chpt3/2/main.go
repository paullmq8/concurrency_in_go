package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var pl = fmt.Println

func main() {
    closure1()
    pl("====================================")
    closure2()
}

func closure2() {
    for _, salutation := range []string{"hello", "greetings", "good day"} {
        pl("closure2", salutation)
        wg.Add(1)
        go func(salutation string) {
            defer wg.Done()
            pl("goroutine", salutation)
        }(salutation)
    }
    wg.Wait()
}

func closure1() {
    for _, salutation := range []string{"hello", "greetings", "good day"} {
        pl("closure1", salutation)
        time.Sleep(1 * time.Second)
        wg.Add(1)
        go func() {
            defer wg.Done()
            pl("goroutine", salutation)
        }()
    }
    wg.Wait()
}