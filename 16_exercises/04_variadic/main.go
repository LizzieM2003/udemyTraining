package main

import "fmt"

func max(nums ...float64) float64 {
	var biggest float64
	for _, v := range nums {
		if v > biggest {
			biggest = v
		}
	}

	return biggest
}

func main() {
	aSlice := []float64{5.7, 3.9, 15.2, 10.7}
	fmt.Println("The largest number is", max(aSlice...))
}
