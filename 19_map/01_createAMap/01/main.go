package main

import "fmt"

func main() {
	myGreeting := make(map[string]string)
	// Can also use the following to create a map variable
	// var myGreeting = map[string]string{}
	myGreeting["Speckles"] = "Woof"
	myGreeting["Hannah"] = "Bonjour"

	fmt.Println(myGreeting)
}
