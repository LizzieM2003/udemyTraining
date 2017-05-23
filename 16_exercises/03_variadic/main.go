package main

import "fmt"

func foo(nums ...int) int {
	var total int
	for _, v := range nums {
		total += v
	}

	return total
}

func main() {
	// numTotal := foo(1, 2)
	// numTotal := foo(1, 2, 3)
	aSlice := []int{1, 2, 3, 4}
	numTotal := foo(aSlice...)
	fmt.Println(numTotal)
}
