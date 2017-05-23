package main

import "fmt"

func changeAge(z *int) {
	fmt.Println(z)
	*z = 24
	fmt.Println(*z)
}

func main() {
	age := 44
	fmt.Println(&age)
	changeAge(&age)
	fmt.Println("Age is now", age)
}
