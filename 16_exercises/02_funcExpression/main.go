package main

import "fmt"

func main() {
	mathsOp := func(a int) (int, bool) {
		return a / 2, a%2 == 0
	}

	z := 20
	x, y := mathsOp(z)
	fmt.Println(z, "divided by 2 is", x, z, "is even is", y)
}
