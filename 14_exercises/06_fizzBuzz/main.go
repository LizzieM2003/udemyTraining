package main

import "fmt"

func main() {

	for i := 0; i < 101; i++ {
		message := ""
		if i%3 == 0 {
			message += "Fizz"
		}
		if i%5 == 0 {
			message += "Buzz"
		}
		if message == "" {
			fmt.Println(i)
		} else {
			fmt.Println(message)
		}
	}
}
