package main

import "fmt"

func changeAge(a int) {
	a = 24
}

func main() {
	age := 44
	changeAge(age)
	fmt.Println(age) // age is still 44
}
