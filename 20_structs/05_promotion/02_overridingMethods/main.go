package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type DoubleZero struct {
	Person
	LicenseToKill bool
}

func (p Person) Greeting() {
	fmt.Println("I'm just an ordinary person.")
}

func (d DoubleZero) Greeting() {
	fmt.Println("Good morning Miss MoneyPenny")
}

func main() {
	p1 := Person{
		Name: "Ian Flemming",
		Age:  75,
	}

	p2 := DoubleZero{
		Person: Person{
			Name: "James Bond",
			Age:  21,
		},
		LicenseToKill: true,
	}

	p1.Greeting()
	p2.Person.Greeting()
	p2.Greeting()
}
