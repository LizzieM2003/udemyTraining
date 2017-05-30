package main

import (
	"fmt"
	"log"
	"math"
)

type NorgateMathError struct {
	lat, long string
	err       error
}

func (n *NorgateMathError) Error() string {
	return fmt.Sprintf("A norgate math error occurred: %v %v %v", n.lat, n.long, n.err)
}

func main() {
	_, err := Sqrt(-10.55)
	if err != nil {
		log.Fatalln(err)
	}
}

func Sqrt(f float64) (float64, error) {
	if f < 0 {
		ErrNorgateMath := fmt.Errorf("norgate math redux: square root of negative number: %v", f)
		return 0, &NorgateMathError{"50.2289N", "99.4656 W", ErrNorgateMath}
	}
	return math.Sqrt(f), nil
}
