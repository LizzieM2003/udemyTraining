package main

import "fmt"

func foo() {
	for i := 0; i < 100; i++ {
		fmt.Println("foo", i)
	}
}

func bar() {
	for i := 0; i < 100; i++ {
		fmt.Println("bar", i)
	}
}

func main() {
	go foo()
	go bar()
}
