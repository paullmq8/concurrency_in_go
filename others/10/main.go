package main

func main() {
	/*
		ch := make(chan bool)
		go func() {
			ch <- true
			ch <- false
			close(ch)
		}()
		println(<-ch)
		println(<-ch)
		<-ch
		<-ch
		<-ch
		v, ok := <-ch
		println(v, ok)
	*/
	ch := make(chan bool)
	go func() {
		ch <- true
		ch <- true
		ch <- true
	}()
	for {
		close(ch)
		b, ok := <-ch
		if !ok {
			break
		}
		println(b)
	}
}
