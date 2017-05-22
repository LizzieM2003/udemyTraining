package main

import "fmt"

func main() {
	name := ""
	fmt.Print("What's your name? ")
	fmt.Scan(&name)
	fmt.Println("Hello", name)
}
