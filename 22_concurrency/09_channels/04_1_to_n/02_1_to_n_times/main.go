package main

import "fmt"

func main() {
	c := make(chan int)
	done := make(chan bool)
	n := 10

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("Add to channel", i)
			c <- i
		}
		close(c)
	}()

	for i := 0; i < n; i++ {
		go func(i int) {
			for x := range c {
				fmt.Println("Retrieving from channel func", i, x)
			}
			done <- true
		}(i)
	}

	for i := 0; i < n; i++ {
		<-done
	}
}
