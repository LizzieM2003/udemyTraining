package main

import "fmt"

func main() {
	var smallNum int
	var largeNum int

	fmt.Print("Enter a smallish number: ")
	fmt.Scan(&smallNum)

	fmt.Print("Enter a larger number: ")
	fmt.Scan(&largeNum)

	fmt.Println("The remainder of", largeNum, "divided by", smallNum, "is", largeNum%smallNum)
}
