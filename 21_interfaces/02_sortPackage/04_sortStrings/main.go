package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"Zeno", "Al", "John", "Jenny"}
	fmt.Println(s)
	sort.Strings(s)
	fmt.Println(s)
}
