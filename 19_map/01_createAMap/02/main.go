package main

import "fmt"

func main() {
	myGreeting := map[string]string{
		"Speckles": "Woof!",
		"Hannah":   "Leave me alone!",
	}

	fmt.Println(myGreeting["Speckles"])
	myGreeting["Marlene"] = "G'day"
	fmt.Println(len(myGreeting))
}
