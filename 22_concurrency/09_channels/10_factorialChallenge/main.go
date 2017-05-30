package main

import "fmt"

func main() {
	c := factorial(5)
	for n := range c {
		fmt.Println(n)
	}
}

func factorial(n int) chan int {
	out := make(chan int)

	go func() {
		factorial := 1
		for i := n; i > 0; i-- {
			factorial *= i
		}
		out <- factorial
		close(out)
	}()

	return out
}
