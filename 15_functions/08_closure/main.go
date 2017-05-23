package main

import "fmt"

func wrapper() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}

func main() {
	incrementA := wrapper()
	incrementB := wrapper()

	fmt.Println("incrementA", incrementA())
	fmt.Println("incrementA", incrementA())
	fmt.Println("incrementB", incrementB())
	fmt.Println("incrementB", incrementB())
	fmt.Println("incrementB", incrementB())
}
