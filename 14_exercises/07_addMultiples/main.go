package main

import "fmt"

func main() {
	var sum int

	for i := 0; i < 1000; i++ {
		if i%3 == 0 {
			sum += i
		} else if i%5 == 0 {
			sum += i
		}
	}

	fmt.Println("The sum of multiples of 3 and 5 is", sum)
}
