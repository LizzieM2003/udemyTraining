package main

import "fmt"

func greeting(fname, lname string) {
	fmt.Println("Hi there", fname, lname)
}

func main() {
	greeting("Lizzie", "Mendes")
}
