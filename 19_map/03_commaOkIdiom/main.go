package main

import "fmt"

func main() {
	myGreeting := map[int]string{
		0: "Good morning!",
		1: "Bonjour!",
		2: "Buenos dias!",
	}

	fmt.Println(myGreeting)

	// delete(myGreeting, 2)

	if val, exists := myGreeting[2]; exists {
		fmt.Println("That value exists")
		fmt.Println("Val:", val)
		fmt.Println("Exists:", exists)
	} else {
		fmt.Println("That value doesn't exist")
		fmt.Println("Val:", val)
		fmt.Println("Exists:", exists)
	}
}
