package main

import "fmt"

func main() {
	examsToday := ""

	switch {
	case examsToday == "Adam":
		fmt.Println("Good luck my boy!")
	case examsToday == "Hannah":
		fmt.Println("Good luck my girl!")
	case examsToday == "Speckles":
		fmt.Println("Dogs don't do exams")
	case len(examsToday) == 0:
		fmt.Println("No one is doing exams today")
	}
}
