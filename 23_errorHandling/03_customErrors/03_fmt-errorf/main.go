package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	_, err := Sqrt(-10.23)
	if err != nil {
		log.Fatalln(err)
	}
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		ErrNorgateMath := fmt.Errorf("Norgate math: square root of negative number: %v", f)
		return 0, ErrNorgateMath
	}
	return math.Sqrt(f), nil
}
