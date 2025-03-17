# Phase One: Functions

## 03-functions

### Key Concepts:

1. **Basic Function Declaration**:

   - Functions are declared using the `func` keyword
   - The `main()` function is the entry point of every Go program
   - Example: `func sayHello() { ... }`

2. **Functions with Parameters and Return Values**:

   - Parameters are declared with their type after the parameter name
   - Return type is specified after the parameter list
   - Example: `func add(a int, b int) int { ... }`

3. **Anonymous Functions**:

   - Functions can be defined without a name
   - Can be called immediately after declaration
   - Can be assigned to variables
   - Example:

     ```go
     // Immediately invoked
     func() { ... }()

     // Assigned to variable
     greet := func(name string) { ... }
     ```

4. **Function Calls**:

   - Functions must be called from within other functions
   - Function calls at package level are not allowed
   - Results of function calls should be used (stored in variables or used in expressions)

5. **Function Scope**:
   - Functions can access variables in their outer scope
   - Functions can be defined inside other functions
   - Each function has its own scope for variables

```go
package main

import "fmt"

// Basic function
func sayHello() {
    fmt.Println("Hello, World!")
}

// Function with parameters
func add(a int, b int) int {
    return a + b
}

// In Go, every executable program requires a special function called main()
// which serves as the starting point for the program.
func main() {
    // Calling basic function
    sayHello()

    // Calling function with parameters
    result := add(10, 20)
    fmt.Println(result)

    // Defining and calling an anonymous function immediately
    func() {
        fmt.Println("This is an anonymous function")
    }()

    // Assigning an anonymous function to a variable
    greet := func(name string) {
        fmt.Printf("Hello, %s!\n", name)
    }

    greet("John") // Calling the anonymous function
}
```
