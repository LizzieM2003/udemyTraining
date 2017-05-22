package main

import "fmt"

func main() {
	examsToday := "Adam"

	switch examsToday {
	case "Adam":
		fmt.Println("Good luck my boy!")
		fallthrough
	case "Hannah":
		fmt.Println("Good luck my girl")
	case "Speckles":
		fmt.Println("Wow, clever dog!")
	default:
		fmt.Println("Everyone is going to have a stressful day!")
	}
}
