package main

func main() {
	ch := make(chan bool, 2)
	ch <- true
	ch <- false
	println("done")
}
