package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    var wg sync.WaitGroup

    // Let's say we have three chefs (goroutines).
    wg.Add(3)

    // Chef 1
    go func() {
        defer wg.Done()
        fmt.Println("Chef 1: Preparing pizza...")
        time.Sleep(2 * time.Second) // Simulating cooking time
        fmt.Println("Chef 1: Pizza is ready!")
    }()

    // Chef 2
    go func() {
        defer wg.Done()
        fmt.Println("Chef 2: Preparing spaghetti...")
        time.Sleep(3 * time.Second) // Simulating cooking time
        fmt.Println("Chef 2: Spaghetti is ready!")
    }()

    // Chef 3
    go func() {
        defer wg.Done()
        fmt.Println("Chef 3: Preparing salad...")
        time.Sleep(1 * time.Second) // Simulating cooking time
        fmt.Println("Chef 3: Salad is ready!")
    }()

    // Wait for all the chefs to finish cooking before we finish the program
    wg.Wait()
    fmt.Println("All dishes are ready! Enjoy your meal!")
}








package main

import (
	"fmt"
	"time"
)

// Named function that will be executed as a Goroutine
func sayHello() {
	for i := 0; i < 3; i++ {
		fmt.Println("Hello!")
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	go sayHello() // Start a Goroutine that calls the sayHello function
	fmt.Println("Main function continues...")
	time.Sleep(time.Second * 2)
}
