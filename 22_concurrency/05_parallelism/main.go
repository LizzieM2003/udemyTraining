package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

// init is used to do some setup for your program and is run first
func init() {
	//the following code is now the default Go setting so
	// is here to show how parallelism in Go works
	// It uses all of the cores available on your machine
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	wg.Add(2)
	go bar()
	go foo()
	wg.Wait()
}

func foo() {
	for i := 0; i < 100; i++ {
		fmt.Println("Foo:", i)
		time.Sleep(3 * time.Millisecond)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 100; i++ {
		fmt.Println("Bar:", i)
		time.Sleep(20 * time.Millisecond)
	}
	wg.Done()
}
