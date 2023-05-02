package main

import (
    "fmt"
    "sync"
)

var wg sync.WaitGroup
var pl = fmt.Println

func main() {
    wg.Add(1)
    go sayHello()
    // continue doing other things
    wg.Add(1)
    go func() {
        defer wg.Done()
        pl("1", "hello")
    }()

    sayHello := func() {
        wg.Done()
        pl("2", "hello")
    }
    wg.Add(1)
    go sayHello()

    salutation := "hello"
    wg.Add(1)
    go func() {
        defer wg.Done()
        pl("3", salutation)
        salutation = "welcome"
    }()
    wg.Wait()

    // It turns out that goroutines execute within the same address space they were created in,
    // and so our program prints out the word “welcome.”
    pl("4", salutation)
}

func sayHello() {
    defer wg.Done()
    fmt.Println("5", "hello")
}
