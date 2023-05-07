package main

import "fmt"

func main() {

	generator := func(values []any) <-chan any {
		valueStream := make(chan any)
		go func() {
			defer close(valueStream)
			for _, v := range values {
				valueStream <- v
			}
		}()
		return valueStream
	}

	stageFn := func(
		done <-chan any,
		values <-chan any,
		fn func(any) any,
	) <-chan any {
		valueChan := make(chan any)
		go func() {
			defer close(valueChan)
			for v := range values {
				select {
				case <-done:
					return
				case valueChan <- fn(v):
				}
			}
		}()
		return valueChan
	}

	fn1 := func(a any) (r any) {
		s, ok := a.(string)
		if ok {
			return s + "1"
		} else {
			return s
		}
	}

	fn2 := func(a any) (r any) {
		s, ok := a.(string)
		if ok {
			return s + "2"
		} else {
			return s
		}
	}

	fn3 := func(a any) (r any) {
		s, ok := a.(string)
		if ok {
			return s + "3"
		} else {
			return s
		}
	}

	toString := func(
		done <-chan any,
		valueStream <-chan any,
	) <-chan string {
		stringStream := make(chan string)
		go func() {
			defer close(stringStream)
			for v := range valueStream {
				select {
				case <-done:
					return
				case stringStream <- "ss: " + v.(string):
				}
			}
		}()
		return stringStream
	}

	done := make(chan any)
	defer close(done)

	source := generator([]any{"a", "b", "c"})
	for v := range toString(done, stageFn(done, stageFn(done, stageFn(done, source, fn1), fn2), fn3)) {
		fmt.Println(v)
	}
}
