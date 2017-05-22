package main

import "fmt"

func main() {
	whoAteTheBiscuits := "Adam"
	switch whoAteTheBiscuits {
	case "Adam", "Hannah":
		fmt.Println("Was it you Adam, or was it you, Hannah?")
	case "Speckles":
		fmt.Println("You need to go to the vet")
	default:
		fmt.Println("The rats ate them")
	}

}
