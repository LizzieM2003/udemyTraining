package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) fullname() string {
	return p.first + " " + p.last
}

func main() {
	p1 := person{"James", "Bond", 20}
	p2 := person{"Miss", "Moneypenny", 18}

	fmt.Println(p1.fullname())
	fmt.Println(p2.fullname())
}
