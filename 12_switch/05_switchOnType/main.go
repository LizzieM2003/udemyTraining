package main

import "fmt"

type contact struct {
	greeting string
	name     string
}

// SwitchOnType(x interface{}) accepts multiple arguments of unknown type
func SwitchOnType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case contact:
		fmt.Println("contact")
	default:
		fmt.Println("unknown")
	}
}

func main() {
	SwitchOnType(42)
	SwitchOnType("Adam")
	t := contact{"Good to see you", "Lizzie"}
	SwitchOnType(t)
	SwitchOnType(t.greeting)
	SwitchOnType(t.name)
}
