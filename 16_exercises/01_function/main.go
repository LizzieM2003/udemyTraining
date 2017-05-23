package main

import "fmt"

func main() {
	z := 21
	x, y := mathsop(z)
	fmt.Println(z, "divided by 2 is", x, z, "is even is", y)
}

func mathsop(a int) (int, bool) {
	return a / 2, a%2 == 0
}
