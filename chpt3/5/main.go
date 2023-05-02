package main

import (
	"fmt"
	"sync"
)

func main() {
    var queue1 chan int
    var queue2 chan<- int
    var queue3 <-chan int
    queue2 = queue1
    queue3 = queue1
    fmt.Println(queue1, queue2, queue3)
    fmt.Println("========================")
    // var queue4 chan int
    // var queue5 chan<- int
    // var queue6 <-chan int
    // queue4 = queue5
    // queue4 = queue6
    // fmt.Println(queue4, queue5, queue6)
    // fmt.Println("========================")
    queue := make(chan string, 2)
    queue <- "one"
    queue <- "two"
    close(queue)
    for elem := range queue {
        fmt.Println(elem)
    }
    fmt.Println("========================")
    ch := make(chan int)
    var wg sync.WaitGroup
    wg.Add(1)
    go func() {
        defer wg.Done()
        for elem := range ch {
            fmt.Println(elem)
        }
    }()
    arr := []int{1, 2, 3, 4, 5}
    wg.Add(1)
    go func(arr []int) {
        defer wg.Done()
        for _, elem := range arr {
            ch <- elem
        }
        close(ch)
    }(arr)
    wg.Wait()
    fmt.Println("========================")
    // queue7 := make(<-chan int)
    // fmt.Println(queue7)
    // close(queue7)
}