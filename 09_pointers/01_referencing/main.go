package main

import "fmt"

func main() {
	a := 43

	fmt.Println("a - ", a)
	fmt.Println("&a - ", &a)

	b := &a

	fmt.Println("b - ", b)
	fmt.Println("&b - ", &b)
	fmt.Println("Value in memory address - ", *b)
}
