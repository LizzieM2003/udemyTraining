package main

import "fmt"

func main() {
	for i := 0; i < 6; i++ {
		for j := 0; j < 6; j++ {
			fmt.Println(i, "-", j)
		}
	}
}
