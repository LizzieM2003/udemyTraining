package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"Zeno", "Al", "John", "Jenny"}
	fmt.Println(s)
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	fmt.Println(s)
}
