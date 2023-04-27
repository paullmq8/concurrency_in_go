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
        pl("hello")
    }()

    sayHello := func(){
        wg.Done()
        pl("hello")
    }
    wg.Add(1)
    go sayHello()

    salutation := "hello"
    wg.Add(1)
    go func() {
        defer wg.Done()
        pl(salutation)
        salutation = "welcome"
    }()
    wg.Wait()
    
    pl(salutation)
}

func sayHello() {
    defer wg.Done()
    fmt.Println("hello")
}