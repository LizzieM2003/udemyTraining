package main

import (
	"fmt"
	"sort"
)

type person struct {
	Name string
	Age  int
}

func (p person) String() string {
	return fmt.Sprintf("YAYAYA %s: %d", p.Name, p.Age)
}

type byAge []person

func (b byAge) Len() int           { return len(b) }
func (b byAge) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b byAge) Less(i, j int) bool { return b[i].Age < b[j].Age }

func main() {
	people := []person{
		{"Bob", 42},
		{"Carol", 48},
		{"Lizzie", 47},
		{"Speckles", 5},
	}
	fmt.Println(people[0])
	sort.Sort(byAge(people))
	fmt.Println(people)
}
