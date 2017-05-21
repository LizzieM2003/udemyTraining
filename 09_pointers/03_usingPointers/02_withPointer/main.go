package main

import "fmt"

func zero(z *int) {
	fmt.Println("Argument z is", z)
	*z = 0
}

func main() {
	x := 5
	fmt.Println("x starts of as", x)
	fmt.Println("memory address of x is", &x)
	zero(&x)
	fmt.Println("x is now", x)
}
