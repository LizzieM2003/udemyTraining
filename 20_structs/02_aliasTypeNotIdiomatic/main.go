package main

import "fmt"

// It is not good practice to alias types unless you need to attach
// methods to the type, see godoc.org/time, type Duration int64

type foo int

func main() {
	var myAge foo
	myAge = 47
	fmt.Printf("%T %v\n", myAge, myAge)

	var yourAge int
	yourAge = 75
	fmt.Printf("%T %v\n", yourAge, yourAge)

	//Need to convert myAge to int to add it to yourAge
	fmt.Println("Our ages combined is", int(myAge)+yourAge)
}
