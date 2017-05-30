package main

import (
	"fmt"
	"sync"
)

func main() {
	in := gen(2, 3)
	// FAN OUT
	c1 := sq("c1", in)
	c2 := sq("c2", in)

	// FAN IN
	for n := range merge(c1, c2) {
		fmt.Println(n) // 4 then 9, or 9 then 4
	}
}

func gen(nums ...int) chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(msg string, in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			fmt.Println("Channel", msg, n)
			out <- n * n
		}
		close(out)
	}()
	return out
}

func merge(cs ...chan int) chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(cs))

	for _, c := range cs {
		go func(ch chan int) {
			for n := range ch {
				out <- n
			}
			wg.Done()
		}(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
