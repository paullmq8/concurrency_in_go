package main

import (
	"fmt"
	"io/ioutil"
	"time"
)

var query = "test"
var matches int

var workerCount = 0                   // 当前工人的数量
var maxWorkerCount = 32               // 最多工人的数量
var searchRequest = make(chan string) // 由包工头来指派活儿
var workerDone = make(chan bool)      // 工人告诉包工头活儿干完了
var foundMatch = make(chan bool)      // 找到搜索结果的消息

func main() {
	start := time.Now()
	workerCount = 1
	go search("/Users/paulliu/", true)
	waitForWorkers()
	fmt.Println(matches, "matches")
	fmt.Println(time.Since(start))
}

func waitForWorkers() {
	for {
		select {
		case path := <-searchRequest:
			workerCount++
			go search(path, true)
		case <-workerDone:
			workerCount--
			if workerCount == 0 {
				return
			}
		case <-foundMatch:
			matches++
		}
	}
}

func search(path string, master bool) {
	files, err := ioutil.ReadDir(path)
	if err == nil {
		for _, file := range files {
			name := file.Name()
			if name == query {
				foundMatch <- true
			}
			if file.IsDir() {
				if workerCount < maxWorkerCount {
					searchRequest <- path + name + "/"
				} else {
					search(path+name+"/", false)
				}
			}
		}
	}
	// 当前search函数是否是在goroutine里运行的
	if master {
		workerDone <- true
	}
}
