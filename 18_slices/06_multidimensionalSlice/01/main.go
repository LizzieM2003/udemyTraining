package main

import "fmt"

func main() {
	records := make([][]string, 0)
	//student 1
	student1 := make([]string, 4)
	student1[0] = "Foster"
	student1[1] = "Nathan"
	student1[2] = "100.00"
	student1[3] = "74.00"

	records = append(records, student1)
	fmt.Println(records)
}
