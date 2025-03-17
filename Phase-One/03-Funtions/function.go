package main

import "fmt"

// Basic function
func sayHello() {
	fmt.Println("Hello, World!")
}

// Function with parameters
func add(a int, b int) int{
	return a + b
}

// We cant do this in Go
// add(10, 30)



// In Go, every executable program requires a special function called main() that serves as the starting point or entry point for the program. This is similar to other compiled languages like C or Java, where there is always a defined entry point (like main()).
func main() {
	sayHello()
	result := add(10, 20)
	fmt.Println(result)



 	// Defining an anonymous function and calling it immediately
    func() {
        fmt.Println("This is an anonymous function")
    }()

	// Assigning an anonymous function to a variable
    greet := func(name string) {
        fmt.Printf("Hello, %s!\n", name)
    }
    
    greet("John") // Calling the anonymous function
}