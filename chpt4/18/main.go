package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

/*
How do we know which one to fan out?
Remember our criteria from earlier: order-independence and duration.
Our random integer generator is certainly order-independent,
but it doesn’t take a particularly long time to run.
The primeFinder stage is also order-independent—numbers are either prime
or not—and because of our naive algorithm,
it certainly takes a long time to run.
It looks like a good candidate for fanning out.
*/
func main() {
	repeatFn := func(
		done <-chan any,
		fn func() any,
	) <-chan any {
		valueStream := make(chan any)
		go func() {
			defer close(valueStream)
			for {
				select {
				case <-done:
					return
				case valueStream <- fn():
				}
			}
		}()
		return valueStream
	}
	take := func(
		done <-chan any,
		valueStream <-chan any,
		num int,
	) <-chan any {
		takeStream := make(chan any)
		go func() {
			defer close(takeStream)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case takeStream <- <-valueStream:
				}
			}
		}()
		return takeStream
	}
	toInt := func(
		done <-chan any,
		valueStream <-chan any,
	) <-chan int {
		intStream := make(chan int)
		go func() {
			defer close(intStream)
			for v := range valueStream {
				select {
				case <-done:
					return
				case intStream <- v.(int):
				}
			}
		}()
		return intStream
	}
	primeFinder := func(
		done <-chan any,
		intStream <-chan int,
	) <-chan any {
		primeStream := make(chan any)
		go func() {
			defer close(primeStream)
			for integer := range intStream {
				integer -= 1
				prime := true
				for divisor := integer - 1; divisor > 1; divisor-- {
					if integer%divisor == 0 {
						prime = false
						break
					}
				}

				if prime {
					select {
					case <-done:
						return
					case primeStream <- integer:
					}
				}
			}
		}()
		return primeStream
	}

	fanIn := func(
		done <-chan any,
		channels ...<-chan any,
	) <-chan any {
		var wg sync.WaitGroup
		multiplexedStream := make(chan any)
		multiplex := func(c <-chan any) {
			defer wg.Done()
			for i := range c {
				select {
				case <-done:
					return
				case multiplexedStream <- i:
				}
			}
		}

		// Select from all the channels
		wg.Add(len(channels))
		for _, c := range channels {
			go multiplex(c)
		}

		// Wait for all the reads to complete
		go func() {
			wg.Wait()
			close(multiplexedStream)
		}()

		return multiplexedStream
	}

	rand := func() any { return rand.Intn(50000000) }

	done := make(chan any)
	defer close(done)

	start := time.Now()

	randIntStream := toInt(done, repeatFn(done, rand))

	numFinders := runtime.NumCPU()
	fmt.Printf("Spinning up %d prime finders.\n", numFinders)
	finders := make([]<-chan any, numFinders)
	fmt.Println("Primes:")

	// fan out
	// Here we’re starting up as many copies of this stage as we have CPUs.
	// On my computer, runtime.NumCPU() returns eight,
	// so I’ll continue to use this number in our discussion.
	// In production, we would probably do a little empirical testing
	// to determine the optimal number of CPUs, but here we’ll stay simple
	// and assume that a CPU will be kept busy by only one copy of the findPrimes stage.
	for i := 0; i < numFinders; i++ {
		finders[i] = primeFinder(done, randIntStream)
	}
	for prime := range take(done, fanIn(done, finders...), 10) {
		fmt.Printf("\t%d\n", prime)
	}

	fmt.Printf("Search took: %v", time.Since(start))
}
