# Phase One: Control Structures

## 02-control-structures

### Key Concepts:

1. **If-else Statements**:
   - Conditional logic to execute code based on a condition.
2. **Switch Statements**:
   - Multi-branching without `if-else`, supports multiple cases per branch.
3. **For Loops**:
   - Go’s only looping construct, which can act as a traditional `for` or `while` loop.
4. **Defer Statements**:
   - Used to delay the execution of a function or statement until the surrounding function returns.

---

### Example Code:

```go
package main

import "fmt"

func main() {
    // If-else statement
    age := 18
    if age >= 18 {
        fmt.Println("You are an adult")
    } else {
        fmt.Println("You are a minor")
    }

    // Switch statement
    switch day := "Monday"; day {
    case "Monday", "Tuesday":
        fmt.Println("Start of the week")
    case "Friday":
        fmt.Println("Almost the weekend")
    default:
        fmt.Println("Midweek")
    }

    // For loop
    for i := 0; i < 5; i++ {
        fmt.Println("Iteration is", i)
    }

    // Defer statement
    defer fmt.Println("This will be executed last")
    fmt.Println("This will be executed first")
}
```
