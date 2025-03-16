package main

import "fmt"

func main(){

	// if else statement
	age := 18
	if age >= 18 {
		fmt.Println("you are an adult")
	} else {
		fmt.Println("you are a minot")
	}

	// switch statement
	switch day := "Monday";
	day {
	case "Monday", "Tuesday":
			fmt.Println("Start of the week")
		case "Friday":
			fmt.Println("Almost the weekend")
		default:
			fmt.Println("Midweek")
	}

	// for loop
	for i := 0; i < 5; i++ {
		fmt.Println("iteration is", i)
	}

	// Defer statement
	// In Go, the defer keyword is used to delay the execution of a function or statement until the surrounding function returns.
	defer fmt.Println("This will be executed last")
	fmt.Println("This will be executed first")

}

// Additional Notes:
// If-else: Go doesn't require parentheses around conditions, but curly braces {} are mandatory.
// For Loop: Go uses a single loop construct (for), which can be used like while and do-while from other languages.
// Switch: Go’s switch is more powerful, supporting multiple cases on one line and not requiring explicit break statements.
// while loop: Go doesn't have a while loop. You can use a for loop to achieve the same effect.
// Defer: Executes the deferred function after the surrounding function completes.