package main

import "fmt"

func main() {
	var i int64
	for i = 1; i <= 100; i++ {
		res := factorialRes(factorialNums(5))
		fmt.Println(<-res)
	}
}

func factorialNums(n int64) chan int64 {
	out := make(chan int64)

	go func() {
		for i := n; i > 0; i-- {
			out <- i
		}
		close(out)
	}()
	return out
}

func factorialRes(in chan int64) chan int64 {
	out := make(chan int64)

	go func() {
		var total int64 = 1
		for n := range in {
			total *= n
		}
		out <- total
		close(out)
	}()
	return out
}
