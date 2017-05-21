package main

import "fmt"

func main() {
	a := 43
	b := &a

	fmt.Println("a - ", a)
	fmt.Println("&a - ", &a)

	fmt.Println("b - ", b)
	fmt.Println("&b - ", &b)
	fmt.Println("Value stored in memory - ", *b)

	//lets change the value now
	*b = 42
	fmt.Println("a is now", a)
}
