package main

import "fmt"

func main() {
	examsToday := "Adam"

	switch examsToday {
	case "Adam":
		fmt.Println("Good luck my boy!")
	case "Hannah":
		fmt.Println("Good luck my girl")
	case "Speckles":
		fmt.Println("Dogs don't do exams")
	default:
		fmt.Println("An easy day today")
	}
}
