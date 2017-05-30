package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var counter int
var wg sync.WaitGroup

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		x := counter
		x++
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		counter = x
		fmt.Println(s, i, "Counter:", counter)
	}
	wg.Done()
}

func main() {
	wg.Add(2)
	go incrementor("Foo")
	go incrementor("Bar")
	wg.Wait()
	fmt.Println("Final counter:", counter)
}

// go run -race main.go
