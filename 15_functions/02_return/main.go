package main

import "fmt"

func main() {
	fmt.Println("Hi there", greet("Lizzie", "Mendes"))
}

func greet(fname, lname string) string {
	return fmt.Sprint(fname, " ", lname)
}
