package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string
	Last  string
	Age   int `json:"wisdom score"`
}

func main() {
	var p1 Person
	bs := []byte(`{"First": "James", "Last":"Bond", "wisdom score": 20}`)
	json.Unmarshal(bs, &p1)
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)
}
