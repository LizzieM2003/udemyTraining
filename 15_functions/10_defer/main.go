package main

import "fmt"

func world() {
	fmt.Println("World!")
}

func hello() {
	fmt.Print("Hello ")
}

func main() {
	defer world()
	hello()
}
