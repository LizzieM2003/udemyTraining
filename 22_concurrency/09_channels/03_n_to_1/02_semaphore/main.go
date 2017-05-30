package main

import "fmt"

func main() {
	c := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("In func 1", i)
			c <- i
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("In func 2", i)
			c <- i
		}
		done <- true
	}()

	go func() {
		<-done
		<-done
		close(c)
	}()

	for n := range c {
		fmt.Println("In main", n)
	}
}
